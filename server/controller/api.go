package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"server/model"

	"github.com/gorilla/mux"
)

/*
Using Gorilla mux router: supports method-based routing
*/
var router *mux.Router

/*
Start listening on port
*/
func Start(port string) {
	router = mux.NewRouter()
	initHandlers()
	fmt.Println("Router initialized and listening on " + port)
	log.Fatal(http.ListenAndServe(":12345", router))
}

/*
Initialize handlers for http requests
*/
func initHandlers() {
	router.HandleFunc("/notes", GetAllNotes).Methods("GET")
	router.HandleFunc("/notes", CreateNote).Methods("POST") // same resource, different verbs
}

/*
GET: Serve all notes to client as JSON
*/
func GetAllNotes(w http.ResponseWriter, r *http.Request) {

	// Set Header
	w.Header().Set("Content-Type", "application/json")

	// Turn slice of posts into json
	notes, err := model.GetAllNotes()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Send json to client
	json.NewEncoder(w).Encode(notes)
}

/*
POST: Add new note to collection. Prohibits duplicates
*/
func CreateNote(w http.ResponseWriter, r *http.Request) {

	// Set Header
	w.Header().Set("Content-Type", "application/json")

	// Turn request json body into note
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	var newNote model.Note
	err := decoder.Decode(&newNote)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Notes must only contain a 'title' and a 'description'"))
		return
	}

	// Prohibit notes without a title
	if newNote.Title == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Notes must have a nonempty 'title'"))
		return
	}

	// Add note to database
	err = model.CreateNote(newNote)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("A note with that title already exists"))
		return
	}

	w.WriteHeader(http.StatusOK)
}
