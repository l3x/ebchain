package main

import (
	"fmt"
	"encoding/hex"
	"github.com/ethereum/go-ethereum/crypto/sha3"
	"math/big"
	"sync/atomic"
)

const (
	AddressLength = 20
	HashLength    = 32
)



// Block represents each 'item' in the blockchain
type Block struct {
	Index     int
	Timestamp string
	Transaction Transaction
	Hash      string
	PrevHash  string
}

// Message takes incoming JSON payload for writing heart rate
type txdata struct {
	From    Address
	To 		Address
	Amount 	int

	Nonce uint64          `json:"nonce"`

	// Signature values
	V *big.Int `json:"v""`
	R *big.Int `json:"r"`
	S *big.Int `json:"s"`

	// This is only used when marshaling to JSON.
	Hash *Hash `json:"hash"`
}

type Header struct {
	hash string
	mrkl_root string
}

type Transaction struct {
	data txdata
	// caches
	Header
}



func (tx Transaction) String() string {
	return fmt.Sprintf("%d: %s -> %s", tx.Amount, tx.From.Hex(), tx.To.Hex())
}

// WithSignature returns a new transaction with the given signature.
// This signature needs to be formatted as described in the yellow paper (v+27).
func (tx *Transaction) WithSignature(signer Signer, sig []byte) (*Transaction, error) {
	r, s, v, err := signer.SignatureValues(tx, sig)
	if err != nil {
		return nil, err
	}
	cpy := &Transaction{data: tx.data}
	cpy.data.R, cpy.data.S, cpy.data.V = r, s, v
	return cpy, nil
}

// Signer encapsulates transaction signature handling. Note that this interface is not a
// stable API and may change at any time to accommodate new protocol rules.
type Signer interface {
	// Sender returns the sender address of the transaction.
	Sender(tx *Transaction) (Address, error)
	// SignatureValues returns the raw R, S, V values corresponding to the
	// given signature.
	SignatureValues(tx *Transaction, sig []byte) (r, s, v *big.Int, err error)
	// Hash returns the hash to be signed.
	Hash(tx *Transaction) common.Hash
	// Equal returns true if the given signer is the same as the receiver.
	Equal(Signer) bool
}

// Address represents the 20 byte address of an Ethereum account.
type Address [AddressLength]byte

// Sets the address to the value of b. If b is larger than len(a) it will panic
func (a *Address) SetBytes(b []byte) {
	if len(b) > len(a) {
		b = b[len(b)-AddressLength:]
	}
	copy(a[AddressLength-len(b):], b)
}

// Hex returns an EIP55-compliant hex string representation of the address.
func (a Address) Hex() string {
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


func BytesToAddress(b []byte) Address {
	var a Address
	a.SetBytes(b)
	return a
}
func HexToAddress(s string) Address   { return BytesToAddress(FromHex(s)) }
