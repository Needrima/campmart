package middlewares

import (
	"campmart/database"
	"campmart/models"
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
)

func GetHomeProducts() []models.Product { // gets 8 random prodicts for home page
	collection := database.GetDatabaseCollection("products")

	var products []models.Product

	sampleStage := bson.M{"$sample": bson.M{"size": 8}} // get 8 random products

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
