package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*
CONSTS
*/
const MONGO_URI = "mongodb://localhost:27017"
const DB_NAME = "api_challenge_db"
const COLLECTION_NAME = "notesCollection"

/*
Document struct for notes
*/
type Note struct {
	Title       string `bson:"title,omitempty"`
	Description string `bson:"description,omitempty"`
}

/*
Main
*/
/*
func main() {

	// Connect to MongoDB
	client := connect(MONGO_URI)

	// Get sole collection we'll be using
	collection := client.Database(DB_NAME).Collection(COLLECTION_NAME)

	// Insert notes
	insertNote(collection, "title2", "desc2")

	// Print notes
	printNotes(collection)

	// Disconnect from MongoDB
	err := client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connection to MongoDB closed.")

}
*/

/*
Connect to mongoDB and returns client
*/
func connect(mongoURI string) (client *mongo.Client) {

	// Set client options, URI
	clientOptions := options.Client().ApplyURI(mongoURI)

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Ensure connected
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")

	return client
}

/*
Add note to collection, returns false if a note with the same title already exists
*/
func insertNote(collection *mongo.Collection, title string, desc string) (success bool) {

	// Try to insert note
	note := Note{title, desc}
	insertResult, err := collection.InsertOne(context.TODO(), note)
	// If schema prevents insertion, we have a duplicate, return false
	if err != nil {
		fmt.Println("Duplicate titile! Note not inserted")
		return false
	}
	fmt.Println("Inserted a single document: ", insertResult.InsertedID)
	return true
}

/*
Print notes in collection
*/
func printNotes(collection *mongo.Collection) {

	// Print all notes
	var results []bson.D

	// Prepare find options
	findOptions := options.Find()
	findOptions.SetProjection(bson.M{"_id": 0}) //dont return _id field

	// Passing bson.D{} (empty document) as the filter matches all documents in the collection
	cur, err := collection.Find(context.TODO(), bson.D{}, findOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Finding multiple documents returns a cursor
	// Iterating through the cursor allows us to decode documents one at a time
	for cur.Next(context.TODO()) {
		// create a value into which the single document can be decoded
		var elem bson.D
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}
		results = append(results, elem)
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	// Close the cursor once finished
	cur.Close(context.TODO())
	fmt.Printf("Found multiple documents (array of pointers): %+v\n", results)
}
