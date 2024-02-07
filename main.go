package main

import (
	"fmt"
	"strconv"

	blocks "github.com/Alphonnse/blockChain/blockchain"
)

// Last TODO 
// add some structures do this thing more really (TSserver, Merkle Tree, payloads to blocks)
func main() {
	chain := blocks.InitBlockChain()

	chain.AddBlock("First block")
	chain.AddBlock("Second block")
	chain.AddBlock("third block")

	for _, block := range chain.Blocks {
		fmt.Printf("Previous hash: %x\n", block.PrevHash)
		fmt.Printf("data: %s\n", block.Data)
		fmt.Printf("hash: %x\n", block.Hash)

		pow := blocks.NewProofOfWork(block)
		fmt.Printf("Pow: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()
	}
}
