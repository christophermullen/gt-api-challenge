package model

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*
Conncection timeout length
*/
const timeoutSec = 3

/*
Save mongo client after InitDB
*/
var client *mongo.Client

/*
Connect to MongoDB
*/
func InitDB(mongoURI string) {

	// Add timeout for connection
	cxt, cancel := context.WithTimeout(context.Background(), time.Duration(timeoutSec)*time.Second)
	defer cancel()

	// Set client options, URI
	clientOptions := options.Client().ApplyURI(mongoURI)

	// Connect to MongoDB
	var err error
	client, err = mongo.Connect(cxt, clientOptions)
	if err != nil {
		log.Fatalf("Error connecting to MongoDB. Is database running with correct URI?\n %v\n", err)
	}

	// Ensure connected
	err = client.Ping(cxt, nil)
	if err != nil {
		log.Fatalf("Error connecting to MongoDB. Is database running with correct URI?\n %v\n", err)
	}
	fmt.Println("Connected to MongoDB!")
}

/*
Disconnect from MongoDB
*/
func Close() {

	// Add timeout for graceful disconnect
	cxt, cancel := context.WithTimeout(context.Background(), time.Duration(timeoutSec)*time.Second)
	defer cancel()

	// Disconnect
	err := client.Disconnect(cxt)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connection to MongoDB closed.")
}
