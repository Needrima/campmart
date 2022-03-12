package database

import "go.mongodb.org/mongo-driver/mongo"

// GetDatabaseCollection returns a database collection by name provided
func GetDatabaseCollection(name string) *mongo.Collection {
	db := initializeDB()

	return db.Collection(name)
}
