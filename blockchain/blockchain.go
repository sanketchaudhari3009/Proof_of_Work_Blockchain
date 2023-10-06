package blockchain

import (
	"sync"
)

type Blockchain struct {
	blocks []*Block
	mu     sync.Mutex
}

func (bc *Blockchain) SetBlockchain(newBlockchain []*Block) {
	bc.mu.Lock()
	defer bc.mu.Unlock()
	bc.blocks = newBlockchain
}

func NewBlockchain() *Blockchain {
	genesisBlock := NewBlock(0, 0, "Genesis Block", 0)
	return &Blockchain{blocks: []*Block{genesisBlock}}
}

func (bc *Blockchain) AddBlock(data int) {
	bc.mu.Lock()
	defer bc.mu.Unlock()

	prevBlock := bc.blocks[len(bc.blocks)-1]
	newBlock := NewBlock(prevBlock.Index+1, data, prevBlock.Hash, 0)
	bc.blocks = append(bc.blocks, newBlock)
}

func (bc *Blockchain) GetAllBlocks() []*Block {
	bc.mu.Lock()
	defer bc.mu.Unlock()
	return bc.blocks
}

func (bc *Blockchain) GetLatestBlock() *Block {
	bc.mu.Lock()
	defer bc.mu.Unlock()

	if len(bc.blocks) == 0 {
		return nil
	}
	return bc.blocks[len(bc.blocks)-1]
}
