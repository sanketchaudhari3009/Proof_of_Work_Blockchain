package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/big"
	"strings"
	"time"
)

type ProofOfWork struct {
	block      *Block
	targetBits int
}

const maxNonce = 1<<32 - 1

func NewProofOfWork(block *Block, targetBits int) *ProofOfWork {
	return &ProofOfWork{
		block:      block,
		targetBits: targetBits,
	}
}

func (pow *ProofOfWork) Run() (int, string) {
	var (
		nonce     int
		hash      string
		hashInt   big.Int
		target    big.Int
		maxTarget big.Int
	)

	target.SetString(strings.Repeat("0", pow.targetBits), 16)
	maxTarget.SetUint64(maxNonce)

	fmt.Printf("Mining a block with difficulty %d\n", pow.targetBits)

	for nonce <= maxNonce {
		select {
		case <-time.After(1 * time.Millisecond):
			data := pow.prepareData(nonce)
			hash = calculateHash(data)
			fmt.Printf("\rNonce: %d Hash: %s", nonce, hash)

			hashInt.SetString(hash, 16)
			if hashInt.Cmp(&target) == -1 {
				fmt.Println("\n")
				return nonce, hash
			}

			nonce++
		}
	}

	fmt.Println("\nMining process exhausted")
	return nonce, ""
}

func (pow *ProofOfWork) prepareData(nonce int) string {
	data := fmt.Sprintf("%s%d%s%d", pow.block.PrevHash, pow.block.Data, pow.block.Timestamp, nonce)
	return data
}

func (pow *ProofOfWork) Validate() bool {
	data := pow.prepareData(pow.block.Nonce)
	hash := calculateHash(data)
	hashInt := new(big.Int)
	hashInt.SetString(hash, 16)

	target := new(big.Int)
	target.SetString(strings.Repeat("0", pow.targetBits), 16)

	return hashInt.Cmp(target) == -1
}

func calculateHash(data string) string {
	hash := sha256.Sum256([]byte(data))
	return hex.EncodeToString(hash[:])
}
