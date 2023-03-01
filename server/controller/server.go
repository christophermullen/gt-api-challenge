package controller

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/gorilla/mux"
)

/*
Using Gorilla mux router: supports easy method-based routing
*/
var router *mux.Router

/*
Start listening on port
*/
func Start(port string) {

	// Set up router and handlers
	router = mux.NewRouter()
	initHandlers()

	// Start server
	server := &http.Server{
		Addr:    port,
		Handler: router,
	}
	go func() {
		err := server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatalf("Error with ListenAndServe(). Port already in use?: %v\n", err)
		}
		log.Println("Shut down server successfully.")
	}()
	log.Println("Server initialized and listening on " + port)

	// Capture SIGINT from server console
	sigChannel := make(chan os.Signal, 1)
	signal.Notify(sigChannel, os.Interrupt)

	// Wait for SIGINT
	<-sigChannel

	// Shut down server
	err := server.Shutdown(context.Background())
	if err != nil {
		log.Fatalf("Failed to shutdown server gracefully: %v\n", err)
	}
}

/*
Initialize handlers for http requests
*/
func initHandlers() {
	router.HandleFunc("/notes", GetAllNotes).Methods("GET")
	router.HandleFunc("/notes", CreateNote).Methods("POST") // same resource, different verbs
}
