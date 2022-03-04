package database

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func initializeDB() *mongo.Database {
	shellURI := "mongodb://localhost:27017"
	clientOptions := options.Client().ApplyURI(shellURI)

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatalf("Error creating mongodb client %v\n:", err)
	}
	defer client.Disconnect(context.Background())

	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		log.Fatalf("Error connecting to mongodb %v\n:", err)
	}

	db := client.Database("campmart")

	return db
}
