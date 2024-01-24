package main

import (
	"fmt"
	"strconv"

	blocks "github.com/Alphonnse/blockChain/blockchain"
)

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

// type Block struct {
// 	Hash     []byte
// 	Data     []byte
// 	PrevHash []byte
// 	Nonce    int
// }
//
// type BlockChain struct {
// 	blocks []*Block
// }

// ------------------------------
// func (b *Block) DeriveHash() {
// 	info := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})
//
// 	hash := sha256.Sum256(info)
//
// 	b.Hash = hash[:]
// }

// func CreateBlock(data string, prevHash []byte) *Block {
// 	block := &Block{[]byte{}, []byte(data), prevHash}
// 	block.DeriveHash()
// 	return block
// }

// -----------------------------

// func CreateBlock(data string, prevHash []byte) *Block {
// 	block := &Block{[]byte{}, []byte(data), prevHash, 0}
//
// 	pow := NewProofOfWork(block)
// 	nonce, hash := pow.Run()
//
// 	block.Hash = hash[:]
// 	block.Nonce = nonce
//
// 	return block
// }
//
// func (chain *BlockChain) AddBlock(data string) {
// 	prevBlock := chain.blocks[len(chain.blocks)-1]
// 	new := CreateBlock(data, prevBlock.Hash)
// 	chain.blocks = append(chain.blocks, new)
// }
//
// func Genesis() *Block {
// 	return CreateBlock("Genesis", []byte{})
// }
//
// func InitBlockChain() *BlockChain {
// 	return &BlockChain{[]*Block{Genesis()}}
// }
