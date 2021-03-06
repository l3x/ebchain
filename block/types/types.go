package types

import (
	"fmt"
	"encoding/hex"
	"github.com/ethereum/go-ethereum/crypto/sha3"
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
type Transaction struct {
	From    Address
	To 		Address
	Amount 	int
}

func (tx Transaction) String() string {
	return fmt.Sprintf("%d: %s -> %s", tx.Amount, tx.From.Hex(), tx.To.Hex())
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
