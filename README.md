go get go.mongodb.org/mongo-driver/mongo

db.notesCollection.createIndex({"title":1},{unique:true})