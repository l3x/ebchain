package main

import (
	"fmt"
	"encoding/json"
	"bytes"
	"net/http"
	"log"
)

func broadcastPeers() {
	for _, peer := range Self.Peers {
		if peer != myAddress {
			// Don't broadcast to self!
			fmt.Printf("Broadcasting my peers to %s/peers\n", peer)
			postMyPeersTo(peer)
		}
	}
}


func broadcastMyFavBook() {
	for _, peer := range Self.Peers {
		if peer != myAddress {
			// Don't broadcast to self!
			fmt.Printf("Broadcasting my peers to %s/peers\n", peer)
			postMyFavBookTo(peer)
		}
	}
}


func postMyFavBookTo(peer string) {

	fmt.Printf(">> postMyFavBookTo (%v) < myPeers: %v\n", fmt.Sprintf("%s/newFavBook", peer),  Self.Peers)

	msg := &FavBookMsg{myAddress, Self.BookInfo}
	bytesRepresentation, err := json.Marshal(msg)
	if err != nil {
		log.Println(err)
		return
	}

	resp, err := http.Post(
		fmt.Sprintf("%s/newFavBook", peer),
		"application/json",
		bytes.NewBuffer(bytesRepresentation))

	if err != nil {
		log.Println(err)
		return
	}
	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)
	log.Println("result", result)
}

func postMyPeersTo(peer string) {

	fmt.Printf(">> postMyPeersTo (%v) < myPeers: %v\n", fmt.Sprintf("%s/newPeers", peer),  Self.Peers)

	msg := &PeersMsg{Self.Peers}
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
		log.Println(err)
		return
	}
	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)
	log.Println("result", result)
	gotNewPeers = false
}

