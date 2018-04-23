package main

import "crypto/ecdsa"

// publicKey -> privateKey
type Host struct {
	account map[string]*ecdsa.PrivateKey
}

func newTesterAccountPool() *Host {
	return &Host{
		account: make(map[string]*ecdsa.PrivateKey),
	}
}


func (ap *Host) sign(header *types.Header, signer string) {
	// Ensure we have a persistent key for the signer
	if ap.account[signer] == nil {
		ap.account[signer], _ = crypto.GenerateKey()
	}
	// Sign the header and embed the signature in extra data
	sig, _ := crypto.Sign(sigHash(header).Bytes(), ap.account[signer])
	copy(header.Extra[len(header.Extra)-65:], sig)
}

func main() {

}
