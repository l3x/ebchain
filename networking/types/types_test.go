package types

import (
	"strings"
	"testing"
	"fmt"
)

func TestIsHexAddress(t *testing.T) {
	tests := []struct {
		str string
		exp bool
	}{
		{"0x5aaeb6053f3e94c9b9a09f33669435e7ef1beaed", true},
		{"5aaeb6053f3e94c9b9a09f33669435e7ef1beaed", true},
		{"0X5aaeb6053f3e94c9b9a09f33669435e7ef1beaed", true},
		{"0XAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA", true},
		{"0xAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA", true},
		{"0x5aaeb6053f3e94c9b9a09f33669435e7ef1beaed1", false},
		{"0x5aaeb6053f3e94c9b9a09f33669435e7ef1beae", false},
		{"5aaeb6053f3e94c9b9a09f33669435e7ef1beaed11", false},
		{"0xxaaeb6053f3e94c9b9a09f33669435e7ef1beaed", false},
	}

	for _, test := range tests {
		if result := IsHexAddress(test.str); result != test.exp {
			t.Errorf("IsHexAddress(%s) == %v; expected %v",
				test.str, result, test.exp)
		}
	}
}


func TestCalculateHash(t *testing.T) {
	tests := []struct {
		str string
		exp bool
	}{
		{"0x095e7baea6a6c7c4c2dfeb977efac326af552d87", false},
		{"095e7baea6a6c7c4c2dfeb977efac326af552d87", false},
		{strings.ToUpper("0x095e7baea6a6c7c4c2dfeb977efac326af552d87"), false},
		{strings.ToUpper("095e7baea6a6c7c4c2dfeb977efac326af552d87"), false},
	}

	for _, test := range tests {
		if result := (test.str == HexToAddress(test.str).Hex()) ; result != test.exp {
			t.Errorf("TestHexToAddress(%s) == %v; expected %v\ntest.str: %s, HexToAddress(test.str).Hex(): %s",
				test.str, result, test.exp, test.str, HexToAddress(test.str).Hex())
		}
	}
}


func newBlock() Block {
	var from, to Address
	from = "0x095e7baea6a6c7c4c2dfeb977efac326af552d87"
	to = "0xb94f5374fce5edbc8e2a8697c15331677e6ebf0b"
	tx := &Transaction{from,to,1000000000000000000}
	fmt.Printf("tx: %s\n", tx)
	newBlock := generateBlock(genesisBlock, *tx)
	return newBlock
}