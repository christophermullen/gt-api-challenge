# Notes API Service


## Prerequisites

Running this server requires Docker and Docker Compose be installed and available.


## Starting the Server Using Docker Compose

Run the following command in the top-level directory of this project:
```shell
$ docker-compose up
```


## Calling the API Endpoints

You may use any tool of your choosing to call the API endpoints (eg cURL, Postman, etc), or you may use the provided scripts. The endpoints are as follows:

### Create a note
```
POST http://localhost:12345/notes
```
- Accepts a JSON object containing a nonempty 'title' and an optional 'description'
- An existing note with the same 'title' cannot already exist in the database

### Read all notes
```
GET http://localhost:12345/notes
```
- Returns a JSON array of all notes in the database


## Using the Provided Scripts

The provided scripts provide a simple way to call the API endpoints. Here are a couple examples of them in action.

### Creating a note

```shell
$ ./scripts/add_note.sh "Cool Title" "Cool Description"
Adding one note:
{"title":"Cool Title","description":"Cool Description"}
```

### Reading all notes

```shell
$ ./scripts/list_notes.sh
Listing all notes:
[{"title":"Cool Title","description":"Cool Description"},{"title":"Meh Title","description":"So-So Description"}]
```


## Necessary Improvements for Production

There are a number of necessary changes if something like this were to be used in production, namely:
- Adding authentication to MongoDB
- Main.go has hard-coded values that should use environment variables
- Unit tests


