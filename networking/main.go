package main

import (
	"bufio"
	"encoding/json"
	"io"
	"log"
	"net"
	"os"
	"sync"
	"time"
	"fmt"

	. "github.com/l3x/ebchain/networking/types"
	"github.com/l3x/hlp"
	"github.com/l3x/ebchain/networking/web"
)


var mutex = &sync.Mutex{}

func main() {
	hlp.LoadEnv()

	go func() {
		// create genesis block
		t := time.Now()
		genesisBlock := Block{}
		genesisBlock = Block{0, t.String(), Transaction{}, CalculateHash(genesisBlock), ""}
		hlp.Debug("genesisBlock", genesisBlock)

		mutex.Lock()
		Blockchain = append(Blockchain, genesisBlock)
		mutex.Unlock()
	}()
	go func() {
		// Listen for create block requests
		log.Fatal(web.RunServer())
	}()

	//BlockChainChannel = make(chan []Block)

	// start TCP and serve TCP server
	tcpPort := os.Getenv("TCP_PORT")
	server, err := net.Listen("tcp", ":"+tcpPort)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("TCP Server Listening on port :", tcpPort)
	defer server.Close()

	for {
		conn, err := server.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handleConn(conn)
	}

}

func handleConn(conn net.Conn) {

	defer conn.Close()

	io.WriteString(conn, "Wait for broadcasts...\n")

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		input := scanner.Text()
		io.WriteString(conn, input)
	}


	// simulate receiving broadcast
	go func() {
		for {
			time.Sleep(30 * time.Second)
			fmt.Println("simulate receiving broadcast")
			mutex.Lock()
			output, err := json.Marshal(Blockchain)
			if err != nil {
				log.Fatal(err)
			}
			mutex.Unlock()
			io.WriteString(conn, string(output))
		}
	}()

	for _ = range web.BlockChainChannel {
		hlp.Debug("Blockchain", Blockchain)
	}

}
