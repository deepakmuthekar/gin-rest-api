package repository

import (
	"context"
	"gin-rest-api/api"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	connectionURL = "mongodb://localhost:27017"
	poolSize      = 100
)

//BookRepository Book repository to interact with Mongo
type BookRepository struct{}

var db *mongo.Database

// Establish a connection to database and create mogodb database handle on loading
// for all subsequent DB interactions.
func init() {
	clientOptions := options.Client().ApplyURI(connectionURL).SetMaxPoolSize(poolSize)
	client, err := mongo.NewClient(clientOptions)

	if err != nil {
		log.Fatal(err)
	}
	err = client.Connect(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	db = client.Database("library")
}

//List list all books
func (br *BookRepository) List() []api.Book {
	log.Println("[BookRepository] List() Getting books from Mongo.")
	collection := db.Collection("books")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{})

	if err != nil {
		log.Fatal("Error while retriving Books..")
	}
	defer cursor.Close(ctx)

	var books []api.Book
	for cursor.Next(ctx) {
		var book api.Book
		cursor.Decode(&book)
		books = append(books, book)
	}
	if err := cursor.Err(); err != nil {
		log.Println("Error while traversing Curor")
		//return error
	}
	return books
}

//Get Book by ID
func (br *BookRepository) Get(id string) api.Book {
	col := db.Collection("books")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	objID, _ := primitive.ObjectIDFromHex(id)

	var b api.Book
	col.FindOne(ctx, bson.M{"_id": objID}).Decode(&b)
	return b
}

//Create creates new Book Document in Mongo
func (br *BookRepository) Create(book api.Book) interface{} {
	log.Println("[BookRepository] Create() Creating new book document in Mongo.")
	collection := db.Collection("books")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := collection.InsertOne(ctx, book)

	if err != nil {
		log.Fatal("Error while Inserting new Book..")
	}
	return result.InsertedID
}
