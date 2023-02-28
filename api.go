package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// simulated database, array of notes
var Notes []Note

func main() {
	Notes = []Note{
		Note{Title: "Hello", Description: "Article Description"},
		Note{Title: "Hello 2", Description: "Article Description"},
	}
	requestHandler()
}

func requestHandler() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/articles", returnAllArticles)
	log.Fatal(http.ListenAndServe(":12345", nil))
}

func homePage(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("Endpoint Hit: homePage")
	fmt.Fprintf(writer, "Welcome to the HomePage!")
}

func returnAllArticles(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("Endpoint Hit: returnAllArticles")
	json.NewEncoder(writer).Encode(Notes)
}
