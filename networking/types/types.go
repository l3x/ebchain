package types

import (
	"fmt"
	"github.com/l3x/hlp"
)

const (
	AddressLength = 20
	HashLength    = 32
)

// Message takes incoming JSON payload for transaction
type Message struct {
	Transaction Transaction `json:"transaction"`
}

// Blockchain is a series of validated Blocks
type BlockChain []Block
var Blockchain BlockChain

func (bc BlockChain) Print()  {
	for _, block := range bc {
		fmt.Printf("%s\n%s\n", block, hlp.Dash(78))
	}
}

// Block represents each 'item' in the blockchain
type Block struct {
	Index     int `json:"index"`
	Timestamp string `json:"timestamp"`
	Transaction Transaction `json:"transaction,string"`
	Hash      string `json:"hash"`
	PrevHash  string `json:"prevhash"`
}

func (b Block) String() string {
	return fmt.Sprintf(
		"Index: %7v\nTimestamp:   %1s\nTransaction: %v\nHash: %71s\nPrevHash: %67s",
		b.Index, b.Timestamp, b.Transaction, b.Hash, b.PrevHash)
}

// Message takes incoming JSON payload
type Transaction struct {
	From   Address `json:"from"`
	To     Address `json:"to"`
	Amount int     `json:"amount"`
}

func (tx Transaction) String() string {
	return fmt.Sprintf("%d: %s -> %s", tx.Amount, tx.From, tx.To)
}

// AddressBytes represents the 20 byte address of an Ethereum account.
type Address string

func (a Address) String() string {
	return string(a)
}