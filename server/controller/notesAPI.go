package controller

import (
	"encoding/json"
	"net/http"
	"server/model"
)

/*
GET: Serve all notes to client as JSON
*/
func GetAllNotes(w http.ResponseWriter, r *http.Request) {

	// Set Header
	w.Header().Set("Content-Type", "application/json")

	// Turn slice of posts into json
	notes, err := model.GetAllNotes(r.Context())
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
		w.Write([]byte("Notes must only contain a 'title' and a 'description'\n"))
		return
	}

	// Prohibit notes without a title
	if newNote.Title == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Notes must have a nonempty 'title'\n"))
		return
	}

	// Add note to database
	err = model.CreateNote(r.Context(), newNote)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("A note with that title already exists\n"))
		return
	}

	w.WriteHeader(http.StatusOK)
}
