package main

import (
	"strconv"
	"time"
	"net/http"
	"log"
)

// web server
func run(port int) error {
	wsHub := newHub()
	mux := makeMuxRouter(wsHub)
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

