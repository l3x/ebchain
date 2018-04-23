package main

import (
	"fmt"
	"crypto/ecdsa"
)

type Host struct {
	Port          int 					// 7001
	BookInfo
	Peers         []string            	// http://localhost:7001
	PeersBookInfo map[string]BookInfo 	// "fav_book": "Book Title.. by Philip Pullman", "fav_cntr": 12
	Key           *ecdsa.PrivateKey
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

func (h Host) SignTx() string {
	return fmt.Sprintf(
		"Port:  %d\nFavBook:  %s\nFavCntr:  %d\nPeers: %+v",
		h.Port, h.FavBook, h.FavCntr, h.Peers)
}

// SignTx signs the transaction using the given signer and private key
func SignTx(tx *Transaction, s Signer, prv *ecdsa.PrivateKey) (*Transaction, error) {
	h := s.Hash(tx)
	sig, err := Sign(h[:], prv)
	if err != nil {
		return nil, err
	}
	return tx.WithSignature(s, sig)
}



func hostAddressFromPort(port int) string {
	return fmt.Sprintf("http://localhost:%d", port)
}


func newHost(port int) (Host, error) {
	bookInfo := &BookInfo{FavBook:randomBook(), FavCntr:0}
	privateKey, _ := GenerateKey()
	host := &Host{Port: port, BookInfo: *bookInfo, PeersBookInfo: make(map[string]BookInfo), Key: privateKey}

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

