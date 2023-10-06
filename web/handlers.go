package web

import (
	"encoding/json"
	"net/http"

	"github.com/sanketchaudhari3009/Proof_of_Work_Blockchain/blockchain"
)

func handleGetBlockchain(w http.ResponseWriter, r *http.Request) {

	blockchainInstance := blockchain.NewBlockchain()
	blocks := blockchainInstance.GetAllBlocks()

	respondWithJSON(w, http.StatusOK, blocks)
}

func handleWriteBlock(w http.ResponseWriter, r *http.Request) {
	var message struct {
		Data int `json:"data"`
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&message); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	blockchainInstance := blockchain.NewBlockchain()

	blockchainInstance.AddBlock(message.Data)

	respondWithJSON(w, http.StatusCreated, "Block added successfully")
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}
