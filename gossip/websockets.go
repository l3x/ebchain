package main

import (
	"log"
	"net/http"
	"sync"
	"github.com/gorilla/websocket"
)

type connection struct {
	// Buffered channel of outbound messages.
	send chan []byte
	// The hub.
	h *hub
}
var hubConn *connection

func (c *connection) reader(wg *sync.WaitGroup, wsConn *websocket.Conn) {
	defer wg.Done()
	for {
		_, message, err := wsConn.ReadMessage()
		if err != nil {
			break
		}
		c.h.broadcast <- message
	}
}

func (c *connection) writer(wg *sync.WaitGroup, wsConn *websocket.Conn) {
	defer wg.Done()
	for message := range c.send {
		err := wsConn.WriteMessage(websocket.TextMessage, message)
		if err != nil {
			break
		}
	}
}

var upgrader = &websocket.Upgrader{ReadBufferSize: 1024, WriteBufferSize: 1024, CheckOrigin: func(r *http.Request) bool { return true }}

type wsHandler struct {
	h *hub
}

func (wsh wsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	wsConn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("error upgrading %s", err)
		return
	}
	hubConn = &connection{send: make(chan []byte, 256), h: wsh.h}
	hubConn.h.addConnection(hubConn)
	defer hubConn.h.removeConnection(hubConn)
	var wg sync.WaitGroup
	wg.Add(2)
	go hubConn.writer(&wg, wsConn)
	go hubConn.reader(&wg, wsConn)
	wg.Wait()
	wsConn.Close()
}
