package main

import (
	"fmt"

	"github.com/nomad-learn/nomad-coin/blockchain"
)

func main() {
	blockchain := blockchain.GetBlockchain()
	blockchain.AddBlock("Second BLock")
	blockchain.AddBlock("Third BLock")

	for _, block := range blockchain.AllBlocks() {
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %s\n", block.Hash)
		fmt.Printf("Prev Hash: %s\n", block.PrevHash)
	}
}
