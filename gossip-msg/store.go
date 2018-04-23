package main

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

func randomBook() string {
	return books[random(0, numberOfBooks)]
}

func setFavBook() {
	Self.FavBook = randomBook()
	Self.FavCntr += 1
	broadcastMyFavBook()
}

func updatePeerFavBook(peerAddress string, bi BookInfo) {
	mutex.Lock()
	Self.PeersBookInfo[peerAddress] = bi
	mutex.Unlock()
}