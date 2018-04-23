package main

import (
	"fmt"
	"encoding/json"
	"io"
	"net/http"
)

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


func handleGetPeersFavBook(w http.ResponseWriter, r *http.Request) {
	fmt.Println(">> GET  handleGetPeersFavBook")
	bytes, err := json.MarshalIndent(Self.PeersBookInfo, "", "  ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	io.WriteString(w, string(bytes))
}


func handleNewPeers(w http.ResponseWriter, r *http.Request) {
	fmt.Println(">> POST handleNewPeers ")
	w.Header().Set("Content-Type", "application/json")
	var m PeersMsg

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


func handleFavBook(w http.ResponseWriter, r *http.Request) {
	fmt.Println(">> POST handleFavBook ")
	w.Header().Set("Content-Type", "application/json")
	var m FavBookMsg

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&m); err != nil {
		respondWithJSON(w, r, http.StatusBadRequest, r.Body)
		return
	}
	defer r.Body.Close()

	for _, peer := range Self.Peers {
		if peer != myAddress {
			updatePeerFavBook(m.PeerAddress, m.BookInfo)
		}
	}

	respondWithJSON(w, r, http.StatusCreated, Self.Peers)
}



func handleSocketIo(w http.ResponseWriter, r *http.Request) {
	fmt.Println(">> POST handleFavBook ")
	w.Header().Set("Content-Type", "application/json")
	var m FavBookMsg

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&m); err != nil {
		respondWithJSON(w, r, http.StatusBadRequest, r.Body)
		return
	}
	defer r.Body.Close()

	for _, peer := range Self.Peers {
		if peer != myAddress {
			updatePeerFavBook(m.PeerAddress, m.BookInfo)
		}
	}

	respondWithJSON(w, r, http.StatusCreated, Self.Peers)
}
