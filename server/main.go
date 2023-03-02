package main

import (
	"log"

	"server/controller"
	"server/model"
)

const port = ":12345"
const mongoURI = "mongodb://mongo:27017"
const timeoutSec = 5
const dbName = "api_challenge_db"
const collectionName = "notesCollection"

func main() {
	log.Printf("Connecting to MongoDB with URI: %v\n", mongoURI)
	model.InitDB(mongoURI, timeoutSec)
	defer model.CloseDB(timeoutSec)

	log.Println("Initializing notes collection...")
	model.InitNotes(dbName, collectionName)

	log.Println("Initializing server...")
	controller.Start(port)
}
