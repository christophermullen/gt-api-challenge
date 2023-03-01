// Add index to "title" field of notes, which prohibits duplicates
db = db.getSiblingDB("api_challenge_db")
db.createCollection('notesCollection');
db.notesCollection.createIndex({"title":1},{unique:true})