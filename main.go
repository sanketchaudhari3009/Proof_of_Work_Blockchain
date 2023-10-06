package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/sanketchaudhari3009/Proof_of_Work_Blockchain/blockchain"
	"github.com/sanketchaudhari3009/Proof_of_Work_Blockchain/web"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	// Initialize the blockchain
	genesisBlock := blockchain.NewBlock(0, "Genesis Block")
	blockchain.Blockchain = append(blockchain.Blockchain, genesisBlock)

	// Start the HTTP server
	if err := web.RunServer(); err != nil {
		log.Fatal(err)
	}
}
