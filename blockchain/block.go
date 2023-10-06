package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"strconv"
	"strings"
	"sync"
	"time"
)

// Block represents each 'item' in the blockchain
type Block struct {
	Index      int
	Timestamp  string
	Data       int
	Hash       string
	PrevHash   string
	Difficulty int
	Nonce      int
}

var mu sync.Mutex

func NewBlock(index int, data int, prevHash string, nonce int) *Block {
	timestamp := time.Now().String()
	block := &Block{
		Index:      index,
		Timestamp:  timestamp,
		Data:       data,
		PrevHash:   prevHash,
		Difficulty: 1,
		Nonce:      nonce,
	}
	block.Hash = block.calculateHash()
	return block
}

func (b *Block) calculateHash() string {
	record := strconv.Itoa(b.Index) + b.Timestamp + strconv.Itoa(b.Data) + b.PrevHash + strconv.Itoa(b.Nonce)
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

func (b *Block) isValid(prevBlock *Block) bool {
	mu.Lock()
	defer mu.Unlock()

	if b.Index != prevBlock.Index+1 {
		return false
	}

	if b.PrevHash != prevBlock.Hash {
		return false
	}

	if b.calculateHash() != b.Hash {
		return false
	}

	if !isHashValid(b.Hash, b.Difficulty) {
		return false
	}

	return true
}

func isHashValid(hash string, difficulty int) bool {
	prefix := strings.Repeat("0", difficulty)
	return strings.HasPrefix(hash, prefix)
}
