package main

import (
	"crypto/sha256"
	"encoding/hex"
	"strconv"
	"time"

	"github.com/l3x/hlp"
	. "github.com/l3x/ebchain/block/types"
	"fmt"
)


func main() {

	t := time.Now()
	genesisBlock := Block{}
	genesisBlock = Block{0, t.String(), Transaction{}, calculateHash(genesisBlock), ""}
	hlp.Debug("genesisBlock", genesisBlock)

	from := HexToAddress("095e7baea6a6c7c4c2dfeb977efac326af552d87")
	to := HexToAddress("b94f5374fce5edbc8e2a8697c15331677e6ebf0b")
	tx := &Transaction{from,to,1000000000000000000}
	fmt.Printf("tx: %s\n", tx)
	newBlock := generateBlock(genesisBlock, *tx)

	if isBlockValid(newBlock, genesisBlock) {
		hlp.Debug("newBlock", genesisBlock)
	}
}

// make sure block is valid by checking index, and comparing the hash of the previous block
func isBlockValid(newBlock, oldBlock Block) bool {
	if oldBlock.Index+1 != newBlock.Index {
		return false
	}

	if oldBlock.Hash != newBlock.PrevHash {
		return false
	}

	if calculateHash(newBlock) != newBlock.Hash {
		return false
	}

	return true
}

// SHA256 hashing
func calculateHash(block Block) string {
	record := strconv.Itoa(block.Index) + block.Timestamp + block.Transaction.String() + block.PrevHash
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

// create a new block using previous block's hash
func generateBlock(oldBlock Block, tx Transaction) Block {

	var newBlock Block

	t := time.Now()

	newBlock.Index = oldBlock.Index + 1
	newBlock.Timestamp = t.String()
	newBlock.Transaction = tx
	newBlock.PrevHash = oldBlock.Hash
	newBlock.Hash = calculateHash(newBlock)

	return newBlock
}
