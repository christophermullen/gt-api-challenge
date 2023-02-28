package main

import (
	"server/controller"
	"server/model"
)

const port = ":12345"
const mongoURI = "mongodb://localhost:27017"
const dbName = "api_challenge_db"
const collectionName = "notesCollection"

func main() {
	model.InitDB(mongoURI)
	defer model.Close()
	model.InitNotes(dbName, collectionName)
	controller.Start(port)
}
