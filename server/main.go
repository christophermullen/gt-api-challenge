package main

import (
	"log"

	"server/controller"
	"server/model"
)

const port = ":12345"
const mongoURI = "mongodb://localhost:27017"
const dbName = "api_challenge_db"
const collectionName = "notesCollection"

func main() {
	log.Printf("Connecting to MongoDB with URI: %v\n", mongoURI)
	model.InitDB(mongoURI)
	defer model.CloseDB()

	log.Println("Initializing notes collection...")
	model.InitNotes(dbName, collectionName)

	log.Println("Initializing server...")
	controller.Start(port)
}
