package model

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*
Document struct for notes
*/
type Note struct {
	Title       string `bson:"title,omitempty"`
	Description string `bson:"description,omitempty"`
}

/*
Save the note collection after InitNotes
*/
var notesCollection *mongo.Collection

/*
At server startup, prepare and save our "notes" collection variable
*/
func InitNotes(dbName string, collectionName string) {
	notesCollection = client.Database(dbName).Collection(collectionName)
}

/*
Returns all notes in the collection as a slice
*/
func GetAllNotes() (notes []Note, err error) {

	// Prepare find options
	findOptions := options.Find()
	findOptions.SetProjection(bson.M{"_id": 0}) //dont return _id field

	// Passing bson.D{} (empty document) as the filter matches all documents in the collection
	cur, err := notesCollection.Find(context.TODO(), bson.D{}, findOptions)
	if err != nil {
		return notes, err
	}

	// Find() returns a cursor
	// Iterating through the cursor decodes documents one at a time
	for cur.Next(context.TODO()) {
		// create value onto which the single document can be decoded
		var elem Note
		err := cur.Decode(&elem)
		if err != nil {
			return notes, err
		}
		notes = append(notes, elem)
	}
	err = cur.Err()
	if err != nil {
		return notes, err
	}

	// Close the cursor once finished
	cur.Close(context.TODO())

	return notes, nil
}

/*
Create note and add to collection
Returns non-nil error if note of identital title is already in collection
*/
func CreateNote(newNote Note) error {

	// Try to insert note
	_, err := notesCollection.InsertOne(context.TODO(), newNote)

	// If schema prevents insertion, we have a duplicate, return error
	if err != nil {
		return err
	}

	return nil
}
