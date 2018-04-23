package main

import (
	"fmt"
	"log"
	"flag"
	"time"
	"sync"
	"os"
	"github.com/joho/godotenv"
)

var (
	Self				Host
	books				[]string
	numberOfBooks		int
	mutex               = &sync.Mutex{}
	myPort				int
	myAddress			string
	myBoostrapAddress	string
	myConfigPath		string
	myRootPath			string
	myTimerFilePath		string
	myIntervalFilePath	string
	gotNewPeers			bool
)

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
	// Get flags
	port := flag.Int("p", 0, "HTTP API port")
	bootstrapPort := flag.Int("b", 0, "Bootstrap server port")
	flag.Parse()
	if *port == 0 {
		log.Fatal("Please provide the API port for this host using the -p parameter")
	}
	if *bootstrapPort == 0 {
		log.Fatal("Please provide the port for the boostrap server using the -b parameter")
	}
	// Set configs
	myPort = *port
	myConfigPath = fmt.Sprintf("%s/configs/%d", myRootPath, myPort)
	fmt.Println("myConfigPath", myConfigPath)
	myTimerFilePath = fmt.Sprintf("%s/%s", myConfigPath, "disable-timer.txt")
	myIntervalFilePath = fmt.Sprintf("%s/%s", myConfigPath, "interval-secs.txt")

	// Set host
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

	// Select favbook and broadcast it to peers (and push it to corresponding web app) every getIntevalSecs
	go func() {
		for {
			time.Sleep(time.Duration(getIntevalSecs()) * time.Second)
			if !maybeDisableTimer() {
				fmt.Println("Selecting new book...")
				setFavBook()
				fmt.Printf("Host =>\n%+v\n\n", Self)
				if gotNewPeers {
					broadcastPeers()
				}
				// websocket call to client
				message := []byte(Self.FavBook)
				if hubConn == nil || hubConn.h == nil {
					log.Println("websocket not ready to broadcast message")
				} else {
					hubConn.h.broadcast <- message
				}
			}
		}
		}()
	log.Fatal(run(Self.Port))
}

