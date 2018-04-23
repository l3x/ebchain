package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

// create handlers
func makeMuxRouter(h *hub) http.Handler {
	muxRouter := mux.NewRouter()
	muxRouter.HandleFunc("/favbook", handleGetMyFav).Methods("GET")
	muxRouter.HandleFunc("/peers", handleGetPeers).Methods("GET")
	muxRouter.HandleFunc("/peersFavBook", handleGetPeersFavBook).Methods("GET")
	muxRouter.HandleFunc("/newPeers", handleNewPeers).Methods("POST")
	muxRouter.HandleFunc("/newFavBook", handleFavBook).Methods("POST")
	muxRouter.Handle("/ws", wsHandler{h: h})
	return muxRouter
}


