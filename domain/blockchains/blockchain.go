package blockchains

import (
	"encoding/json"
	"fmt"

	"github.com/juankohler/web3-blockchain-go/domain/blocks"
)

type Blockchain struct {
	Chains map[int]blocks.Block
}

func CreateBlockchain() *Blockchain {
	return &Blockchain{
		Chains: map[int]blocks.Block{
			0: blocks.CreateGenesisBlock(),
		},
	}
}

func (s Blockchain) GetLastestBlock() blocks.Block {
	return s.Chains[len(s.Chains)-1]
}

func (s Blockchain) AddBlock(newBlock blocks.Block) {
	s.Chains[len(s.Chains)] = newBlock
}

func (s Blockchain) ValidateChain() bool {
	for i := 1; i < len(s.Chains); i++ {
		currentBlock := s.Chains[i]
		previousBlock := s.Chains[i-1]

		if currentBlock.Hash != blocks.CalculateHash(currentBlock.Timestamp, currentBlock.Data, currentBlock.PreviousHash) {
			return false
		}

		if currentBlock.PreviousHash != previousBlock.Hash {
			return false
		}
	}

	return true
}

func (s Blockchain) Print() {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		fmt.Println("Marshal json error:", err)
		return
	}
	fmt.Println()
	fmt.Println(string(jsonData))
	fmt.Println()

}
