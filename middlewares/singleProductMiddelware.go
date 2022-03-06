package middlewares

import (
	"campmart/database"
	"campmart/models"
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var productsCollection = database.GetDatabaseCollection("products")

func GetSingeProduct(id string) models.Product {
	var product models.Product

	singleResult := productsCollection.FindOne(context.TODO(), bson.M{"id": id})

	if err := singleResult.Err(); err == mongo.ErrNoDocuments {
		log.Printf("No document with id{%v} found in database: Error Message{%v}\n", id, err)
		return models.Product{}
	}

	if err := singleResult.Decode(&product); err != nil {
		log.Printf("Error decoding document into product struct: %v\n", err)
		return models.Product{}
	}

	fmt.Println("Single product:", product)

	return product
}

func GetSuggestionsProducts() []models.Product {
	var products []models.Product

	sampleStage := bson.M{"$sample": bson.M{"size": 4}} // get 4 random products for suggestions

	productsCursor, err := productsCollection.Aggregate(context.TODO(), []bson.M{sampleStage})
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
