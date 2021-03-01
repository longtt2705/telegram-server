package utils

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// Model returns model for user collection
type Model struct {
	Collection *mongo.Collection
}

var database *mongo.Database

// ConnectDatabase connects to a database with given name
func ConnectDatabase(dbName string) func() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb+srv://dbTelegram:waOhUa55qYqg38nH@cluster0.al5iq.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"))
	if err != nil {
		log.Fatal(err)
	}
	disconnectCallback := func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}
	database = client.Database(dbName)
	return disconnectCallback
}

// GetCollection returns collection from database
func GetCollection(collectionName string) *mongo.Collection {
	if database == nil {
		return nil
	}

	return database.Collection(collectionName)
}
