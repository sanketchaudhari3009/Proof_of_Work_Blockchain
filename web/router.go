package web

import "github.com/gorilla/mux"

func makeMuxRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/", handleGetBlockchain).Methods("GET")
	router.HandleFunc("/", handleWriteBlock).Methods("POST")

	// Serve static files or HTML templates if needed
	// router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	// router.HandleFunc("/", handleHTMLTemplate)

	return router
}
