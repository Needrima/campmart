package database

import "go.mongodb.org/mongo-driver/mongo"

func GetDatabaseCollection(name string) *mongo.Collection {
	db := initializeDB()

	return db.Collection(name)
}
