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

