package middlewares

import (
	"campmart/database"
	"campmart/models"
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
)

// GetHomeProducts reurn 8 random products to be served to index.html
func GetHomeProducts() []models.Product {
	collection := database.GetDatabaseCollection("products")

	var products []models.Product

	sampleStage := bson.M{"$sample": bson.M{"size": 4}} // get 8 random products

	productsCursor, err := collection.Aggregate(context.TODO(), []bson.M{sampleStage})
	if err != nil {
		log.Printf("error getting sample cursor: %v", err)
		return []models.Product{}
	}

	if err := productsCursor.All(context.TODO(), &products); err != nil {
		log.Println("Error writing cursor content to product:", err)
		return []models.Product{}
	}

	return products
}
