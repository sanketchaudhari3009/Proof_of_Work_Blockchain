package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/sanketchaudhari3009/Proof_of_Work_Blockchain/blockchain"
	"github.com/sanketchaudhari3009/Proof_of_Work_Blockchain/web"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	genesisBlock := blockchain.NewBlock(0, 0, "Genesis Block", 0)
	initialBlockchain := []*blockchain.Block{genesisBlock}
	blockchainInstance := blockchain.NewBlockchain()
	blockchainInstance.SetBlockchain(initialBlockchain)

	if err := web.StartServer(); err != nil {
		log.Fatal(err)
	}
}
