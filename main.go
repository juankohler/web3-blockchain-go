package main

import (
	"fmt"
	"time"

	"github.com/juankohler/web3-blockchain-go/domain/blockchains"
	"github.com/juankohler/web3-blockchain-go/domain/blocks"
)

func main() {
	/** Create blockchain */
	blockchain := blockchains.CreateBlockchain()

	/** Add Block */
	block1 := blocks.CreateBlock(
		time.Date(2009, 1, 3, 0, 0, 0, 0, time.UTC).Unix(),
		blocks.CreateBlockData("Juancho", "Messi", 10),
		blockchain.GetLastestBlock().Hash,
	)
	blockchain.AddBlock(block1)

	/** Add Block */
	block2 := blocks.CreateBlock(
		time.Date(2009, 1, 3, 0, 0, 0, 0, time.UTC).Unix(),
		blocks.CreateBlockData("Messi", "Juancho", 10),
		blockchain.GetLastestBlock().Hash,
	)
	blockchain.AddBlock(block2)

	/** Print blockchain */
	blockchain.Print()

	/** Validate chains */
	fmt.Printf("Is blockchain valid? %t \n\n", blockchain.ValidateChain())

	/** Invalidate chain */
	aux := blockchain.Chains[1]
	aux.PreviousHash = "fake_hash"
	blockchain.Chains[1] = aux

	/** Validate chains */
	fmt.Printf("Is blockchain valid? %t \n\n", blockchain.ValidateChain())
}
