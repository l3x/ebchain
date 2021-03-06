package types

import (
	"time"
	"strconv"
	"crypto/sha256"
	"encoding/hex"
)

func GenesisBlock() Block {
	t := time.Now()
	genesisBlock := Block{}
	genesisBlock = Block{0, t.String(), Transaction{}, CalculateHash(genesisBlock), ""}
	return 	genesisBlock
}

// SHA256 hashing
func CalculateHash(block Block) string {
	record := strconv.Itoa(block.Index) + block.Timestamp + block.Transaction.String() + block.PrevHash
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

// make sure block is valid by checking index, and comparing the hash of the previous block
func IsBlockValid(newBlock, oldBlock Block) bool {
	if oldBlock.Index+1 != newBlock.Index {
		return false
	}

	if oldBlock.Hash != newBlock.PrevHash {
		return false
	}

	if CalculateHash(newBlock) != newBlock.Hash {
		return false
	}
	return true
}

// create a new block using previous block's hash
func GenerateBlock(oldBlock Block, tx Transaction) Block {

	var newBlock Block

	t := time.Now()

	newBlock.Index = oldBlock.Index + 1
	newBlock.Timestamp = t.String()
	newBlock.Transaction = tx
	newBlock.PrevHash = oldBlock.Hash
	newBlock.Hash = CalculateHash(newBlock)

	return newBlock
}
