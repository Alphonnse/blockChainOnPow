package blocks

import (
	"time"

	"github.com/cbergoon/merkletree"
)

// Its the base block and all stuff works good with it
// type Block struct {
// 	Hash     []byte
// 	Data     []byte
// 	PrevHash []byte
// 	Nonce    int
// }

type Tx struct {
	InputVal  float64
	OutputVal float64
	TxTime    time.Time
}

// type MarkelTree struct {
// 	Hash  []byte
// 	Left  *MarkelTree
// 	Right *MarkelTree
// }

type BlockHeader struct {
	PrevHash   []byte
	TimeStamp  time.Time
	Nonce      int
	Difficulty int
	MarkelRoot merkletree.MerkleTree
}

type BlockBody struct {
	Txs [2048]*Tx
}

type Block struct {
	Header BlockHeader
	Body   BlockBody
}

type BlockChain struct {
	Blocks []*Block
}

func CreateBlock(data string, prevHash []byte) *Block {
	block := &Block{[]byte{}, []byte(data), prevHash, 0}

	pow := NewProofOfWork(block)
	nonce, hash := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce

	return block
}

func (chain *BlockChain) AddBlock(data string) {
	prevBlock := chain.Blocks[len(chain.Blocks)-1]
	new := CreateBlock(data, prevBlock.Hash)
	chain.Blocks = append(chain.Blocks, new)
}

func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
}

func InitBlockChain() *BlockChain {
	return &BlockChain{[]*Block{Genesis()}}
}
