package types

import (
	"fmt"
	"encoding/hex"
	"github.com/ethereum/go-ethereum/crypto/sha3"
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
	//return fmt.Sprintf("%d: %s -> %s", tx.Amount, tx.From.Hex(), tx.To.Hex())
	return fmt.Sprintf("%d: %s -> %s", tx.Amount, tx.From, tx.To)
}


// AddressBytes represents the 20 byte address of an Ethereum account.
type AddressBytes [AddressLength]byte
type Address string

// Sets the address to the value of b. If b is larger than len(a) it will panic
func (a *AddressBytes) SetBytes(b []byte) {
	if len(b) > len(a) {
		b = b[len(b)-AddressLength:]
	}
	copy(a[AddressLength-len(b):], b)
}

// Hex returns an EIP55-compliant hex string representation of the address.
func (a AddressBytes) Hex() string {
	unchecksummed := hex.EncodeToString(a[:])
	sha := sha3.NewKeccak256()
	sha.Write([]byte(unchecksummed))
	hash := sha.Sum(nil)

	result := []byte(unchecksummed)
	for i := 0; i < len(result); i++ {
		hashByte := hash[i/2]
		if i%2 == 0 {
			hashByte = hashByte >> 4
		} else {
			hashByte &= 0xf
		}
		if result[i] > '9' && hashByte > 7 {
			result[i] -= 32
		}
	}
	return "0x" + string(result)
}


func BytesToAddress(b []byte) AddressBytes {
	var a AddressBytes
	a.SetBytes(b)
	return a
}
func HexToAddress(s string) AddressBytes { return BytesToAddress(FromHex(s)) }
