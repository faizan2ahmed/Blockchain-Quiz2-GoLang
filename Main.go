package main

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

type Blockchain struct {
	blocks []*Block
}

func NewBlock(data string, previousHash string) *Block {
	return &Block{
		Timestamp:    time.Now().Unix(),
		Data:         data,
		PreviousHash: previousHash,
	}
}

func (b *Block) CalculateHash() string {
	hash := sha256.New()
	hash.Write([]byte(fmt.Sprintf("%d%s%s", b.Timestamp, b.Data, b.PreviousHash)))
	return hex.EncodeToString(hash.Sum(nil))
}

func (bc *Blockchain) AddBlock(data string) {
	prevBlock := bc.blocks[len(bc.blocks)-1]
	newBlock := NewBlock(data, prevBlock.Hash)
	newBlock.Hash = newBlock.CalculateHash()
	bc.blocks = append(bc.blocks, newBlock)
}

func (bc *Blockchain) DisplayAllBlocks() {
	for _, block := range bc.blocks {
		fmt.Printf("Timestamp: %d, Data: %s, PreviousHash: %s, Hash: %s\n",
			block.Timestamp, block.Data, block.PreviousHash, block.Hash)
	}
}

func (bc *Blockchain) NewBlock(data string) {
	prevBlock := bc.blocks[len(bc.blocks)-1]
	newBlock := NewBlock(data, prevBlock.Hash)
	newBlock.Hash = newBlock.CalculateHash()
	bc.blocks = append(bc.blocks, newBlock)
}

func (bc *Blockchain) ModifyBlock(index int, newData string) error {
	if index < 0 || index >= len(bc.blocks) {
		return fmt.Errorf("invalid block index")
	}
	bc.blocks[index].Data = newData
	bc.blocks[index].Hash = bc.blocks[index].CalculateHash()
	return nil
}

func NewBlockchain() *Blockchain {
	genesisBlock := NewBlock("Genesis Block", "")
	genesisBlock.Hash = genesisBlock.CalculateHash()
	return &Blockchain{
		blocks: []*Block{genesisBlock},
	}
}

func main() {

	bc := NewBlockchain()

	bc.NewBlock("First Block after Genesis")
	bc.NewBlock("Second Block after Genesis")
	bc.NewBlock("Third Block after Genesis")

	bc.DisplayAllBlocks()

	err := bc.ModifyBlock(1, "Modified Second Block")
	if err != nil {
		fmt.Println("Failed to modify block:", err)
		return
	}

	fmt.Println("\nAfter modification:")
	bc.DisplayAllBlocks()
}
