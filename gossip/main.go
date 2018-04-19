package main

import (
	"fmt"
	"log"
	"flag"
	"net/http"
	"github.com/gorilla/mux"
	"encoding/json"
	"io"
	"time"
	"strconv"
	"io/ioutil"
	"strings"
	"math/rand"
	"sync"
	"bytes"
	"os"
	"github.com/joho/godotenv"
	"github.com/l3x/hlp"
)


var (
	Self               Host
	books              []string
	numberOfBooks      int
	mutex               = &sync.Mutex{}
	myPort				int
	myAddress			string
	myBoostrapAddress	string
	myConfigPath		string
	myRootPath			string
	myTimerFilePath		string
	gotNewPeers			bool
)

type Host struct {
	Port  	int			// 7001
	BookInfo
	Peers 	[]string	// http://localhost:7001
}


type BookInfo struct {
	FavBook	string		// A Handful of Dust by Evelyn Waugh
	FavCntr	int			// Sequential count of favorite book choices
}
func (bi BookInfo) String() string {
	return fmt.Sprintf("FavBook: %s, FavCntr: %d", bi.FavBook, bi.FavCntr)
}

func (h Host) String() string {
	return fmt.Sprintf(
		"Port:  %d\nFavBook:  %s\nFavCntr:  %d\nPeers: %+v",
		h.Port, h.FavBook, h.FavCntr, h.Peers)
}

func hostAddressFromPort(port int) string {
	return fmt.Sprintf("http://localhost:%d", port)
}

// Message takes incoming JSON payload for transaction
type Message struct {
	Peers []string `json:"peers"`  // [ "http://localhost:7001", "http://localhost:7002" ]
}

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	myRootPath = os.Getenv("ROOT_DIR")

	fmt.Printf("%s/books.txt\n", myRootPath)
	books = readLines(fmt.Sprintf("%s/books.txt", myRootPath))
	numberOfBooks = len(books)
	log.Println("Read ", numberOfBooks, " books!")

}

func main() {

	port := flag.Int("p", 0, "HTTP API port")
	bootstrapPort := flag.Int("b", 0, "Bootstrap server port")
	flag.Parse()
	if *port == 0 {
		log.Fatal("Please provide the API port for this host using the -p parameter")
	}
	if *bootstrapPort == 0 {
		log.Fatal("Please provide the port for the boostrap server using the -b parameter")
	}
	myPort = *port
	myConfigPath = fmt.Sprintf("%s/configs/%d", myRootPath, myPort)
	fmt.Println("myConfigPath", myConfigPath)
	myTimerFilePath = fmt.Sprintf("%s/%s", myConfigPath, "disable-timer.txt")

	var err error
	Self, err = newHost(myPort)
	if err != nil {
		log.Fatal(err)
	}
	addPeer(hostAddressFromPort(*bootstrapPort))
	myAddress = hostAddressFromPort(Self.Port)
	addPeer(myAddress)
	setFavBook()
	fmt.Printf("host =>\n%+v\n", Self)
	hlp.Debug("Self", Self)

	go func() {
		for {
			time.Sleep(time.Duration(5) * time.Second)
			if !disableTimer() {
				fmt.Println("Selecting new book...")
				setFavBook()
				fmt.Printf("Host =>\n%+v\n\n", Self)
				if gotNewPeers {
					broadcastPeers()
				}
				fmt.Println("\n")
			}
		}
		}()
	log.Fatal(run(Self.Port))
}

func broadcastPeers() {
	for _, peer := range Self.Peers {
		if peer != myAddress {
			// Don't broadcast to self!
			fmt.Printf("Broadcasting my peers to %s/peers\n", peer)
			postMyPeersTo(peer)
		}
	}
}

// web server
func run(port int) error {
	mux := makeMuxRouter()
	log.Println("HTTP Server Listening on port :", port)
	s := &http.Server{
		Addr:           ":" + strconv.Itoa(port),
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	if err := s.ListenAndServe(); err != nil {
		return err
	}
	return nil
}

// create handlers
func makeMuxRouter() http.Handler {
	muxRouter := mux.NewRouter()
	muxRouter.HandleFunc("/favbook", handleGetMyFav).Methods("GET")
	muxRouter.HandleFunc("/peers", handleGetPeers).Methods("GET")
	muxRouter.HandleFunc("/newPeers", handleNewPeers).Methods("POST")
	return muxRouter
}

func handleGetMyFav(w http.ResponseWriter, r *http.Request) {
	fmt.Println(">> GET  handleGetMyFav")
	bytes, err := json.MarshalIndent(Self.BookInfo, "", "  ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	io.WriteString(w, string(bytes))
}

func handleGetPeers(w http.ResponseWriter, r *http.Request) {
	fmt.Println(">> GET  handleGetPeers")
	bytes, err := json.MarshalIndent(Self.Peers, "", "  ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	io.WriteString(w, string(bytes))
}

// takes JSON payload as an input
func handleNewPeers(w http.ResponseWriter, r *http.Request) {
	fmt.Println(">> POST handleNewPeers ")
	w.Header().Set("Content-Type", "application/json")
	var m Message

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&m); err != nil {
		respondWithJSON(w, r, http.StatusBadRequest, r.Body)
		return
	}
	defer r.Body.Close()

	for _, peer := range m.Peers {
		if !stringInSlice(peer, Self.Peers) {
			addPeer(peer)
		}
	}

	respondWithJSON(w, r, http.StatusCreated, Self.Peers)
}

func respondWithJSON(w http.ResponseWriter, r *http.Request, code int, payload interface{}) {
	response, err := json.MarshalIndent(payload, "", "  ")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("HTTP 500: Internal Server Error"))
		return
	}
	w.WriteHeader(code)
	w.Write(response)
}


func newHost(port int) (Host, error) {
	bookInfo := &BookInfo{FavBook:randomBook(), FavCntr:0}
	host := &Host{Port: port, BookInfo: *bookInfo}

	log.Println("Host", Self)
	return *host, nil
}


func readLines(path string) ([]string) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(fmt.Sprintf("Unable to read %", path))
	}
	return strings.Split(string(content), "\n")
}


func disableTimer() bool {
	disableTimerLines := readLines(myTimerFilePath)
	return stringInSlice("true", disableTimerLines)
}

func random(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max - min) + min
}

func randomBook() string {
	 return books[random(0, numberOfBooks)]
}

func setFavBook() {
	Self.FavBook = randomBook()
	Self.FavCntr += 1
}


func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}


func postMyPeersTo(peer string) {

	fmt.Printf(">> postMyPeersTo (%v) < myPeers: %v\n", fmt.Sprintf("%s/newPeers", peer),  Self.Peers)

	msg := &Message{Self.Peers}
	bytesRepresentation, err := json.Marshal(msg)
	if err != nil {
		log.Println(err)
		return
	}

	resp, err := http.Post(
		fmt.Sprintf("%s/newPeers", peer),
		"application/json",
		bytes.NewBuffer(bytesRepresentation))

	if err != nil {
		//log.Fatalln(err)
		log.Println(err)
		return
	}

	var result map[string]interface{}

	json.NewDecoder(resp.Body).Decode(&result)

	log.Println(result)
	log.Println(result["data"])
	gotNewPeers = false
}

// Ex: addPeer("http://localhost:7001")
func addPeer(peer string) {
	// If not already in peers list (do include my address into peers list)
	if !stringInSlice(peer, Self.Peers) {
		mutex.Lock()
		Self.Peers = append(Self.Peers, peer)
		gotNewPeers = true
		mutex.Unlock()
	}
}
