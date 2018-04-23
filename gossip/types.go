package main

import (
	"fmt"
)

type Host struct {
	Port          int					// 7001
	BookInfo
	Peers         []string				// http://localhost:7001
	PeersBookInfo map[string]BookInfo	//
}

type BookInfo struct {
	FavBook	string `json:"fav_book"`	// A Handful of Dust by Evelyn Waugh
	FavCntr	int	`json:"fav_cntr"` 		// Sequential count of favorite book choices
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


func newHost(port int) (Host, error) {
	bookInfo := &BookInfo{FavBook:randomBook(), FavCntr:0}
	host := &Host{Port: port, BookInfo: *bookInfo, PeersBookInfo: make(map[string]BookInfo)}

	fmt.Println("Host", Self)
	return *host, nil
}

type PeersMsg struct {
	Peers []string `json:"peers"`  // [ "http://localhost:7001", "http://localhost:7002" ]
}

type FavBookMsg struct {
	PeerAddress string  `json:"peer_address"`
	BookInfo BookInfo `json:"book_info"`  // [ "http://localhost:7001", "http://localhost:7002" ]
}

