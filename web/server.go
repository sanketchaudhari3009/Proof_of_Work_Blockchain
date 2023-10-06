package web

import (
	"log"
	"net/http"
	"os"
)

func StartServer() error {
	router := makeMuxRouter()

	httpPort := os.Getenv("PORT")
	if httpPort == "" {
		httpPort = "8080"
	}

	serverAddr := ":" + httpPort
	log.Printf("HTTP Server Listening on port %s\n", httpPort)

	server := &http.Server{
		Addr:    serverAddr,
		Handler: router,
	}

	return server.ListenAndServe()
}
