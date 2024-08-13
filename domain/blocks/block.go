package blocks

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"
)

type Block struct {
	Timestamp    int64
	Data         string
	PreviousHash string
	Hash         string
}

func CreateBlock(timestamp int64, data string, previousHash string) Block {
	calculateHash := CalculateHash(timestamp, data, previousHash)
	return Block{
		Timestamp:    timestamp,
		Data:         data,
		PreviousHash: previousHash,
		Hash:         calculateHash,
	}

}

func CalculateHash(timestamp int64, data string, previousHash string) string {
	value := fmt.Sprintf("%d%s%s", timestamp, data, previousHash)
	hash := sha256.Sum256([]byte(value))
	hashHex := hex.EncodeToString(hash[:])
	return hashHex
}

func CreateBlockData(from string, to string, amount float64) string {
	return fmt.Sprintf("from: %s, to: %s, amount: %.6f", from, to, amount)
}

func CreateGenesisBlock() Block {
	timestamp := time.Date(2009, 1, 3, 0, 0, 0, 0, time.UTC).Unix()
	data := "The Times 03/Jan/2009 Chancellor on brink of second bailout for banks"
	previousHash := "0"
	return CreateBlock(timestamp, data, previousHash)
}
