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
	Title       string `bson:"title" json:"title"`
	Description string `bson:"description" json:"description"`
}

/*
Notes collection
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
If it fails to decode a note, it returns an empty slice and non-nil error
*/
func GetAllNotes(cxt context.Context) (notes []Note, err error) {

	// Prepare find options
	findOptions := options.Find()
	findOptions.SetProjection(bson.M{"_id": 0}) //dont return _id field

	// Passing bson.D{} (empty document) as the filter matches all documents in the collection
	cur, err := notesCollection.Find(cxt, bson.D{}, findOptions)
	if err != nil {
		return notes, err
	}

	// Find() returns a cursor
	// Iterating through the cursor decodes documents one at a time
	for cur.Next(cxt) {
		// create value onto which the single document can be decoded
		var elem Note
		err := cur.Decode(&elem)
		if err != nil {
			notes = make([]Note, 0)
			return notes, err
		}
		notes = append(notes, elem)
	}
	err = cur.Err()
	if err != nil {
		return notes, err
	}

	// Close the cursor once finished
	cur.Close(cxt)

	return notes, nil
}

/*
Create note and add to collection
Returns non-nil error if note of identital title is already in collection
*/
func CreateNote(cxt context.Context, newNote Note) error {

	// Try to insert note
	_, err := notesCollection.InsertOne(cxt, newNote)

	// If schema prevents insertion, we have a duplicate, return error
	if err != nil {
		return err
	}

	return nil
}
