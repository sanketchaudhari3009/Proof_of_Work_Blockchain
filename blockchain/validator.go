package blockchain

// isBlockValid checks if a new block is valid
func isBlockValid(newBlock *Block, prevBlock *Block) bool {
	if newBlock.Index != prevBlock.Index+1 {
		return false
	}

	if newBlock.PrevHash != prevBlock.Hash {
		return false
	}

	if calculateHash(newBlock.Hash) != newBlock.Hash {
		return false
	}

	if !isHashValid(newBlock.Hash, newBlock.Difficulty) {
		return false
	}

	return true
}
