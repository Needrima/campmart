package middlewares

import (
	"campmart/database"
	"campmart/helpers"
	"campmart/models"
	"context"
	"fmt"
	"log"

	// "strings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetShopProducts() []models.Product {
	productsCollection := database.GetDatabaseCollection("products")

	var products []models.Product

	sampleStage := bson.M{"$sample": bson.M{"size": 12}} // get 12 random products

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

func GetSearchSuggestions(input string) []string {

	productsCollection := database.GetDatabaseCollection("products")

	productsCursor, err := productsCollection.Find(context.TODO(), bson.M{"name": bson.M{"$regex": "(?i)" + input}})
	if err != nil {
		log.Println("Error getting products cursor in search suggestions:", err)
		return []string{}
	}

	var suggestions []string

	for productsCursor.Next(context.TODO()) {
		var p models.Product

		if err := productsCursor.Decode(&p); err != nil {
			fmt.Println("Error getting product:", err)
		}

		if !helpers.FoundString(suggestions, p.Name) {
			suggestions = append(suggestions, p.Name)
		}
	}

	return suggestions
}

func GetProductsFromSearchInput(searchInput string, pageNumber int) []models.Product {
	limit, skip := int64(12), int64(12*pageNumber)
	findOptions := &options.FindOptions{
		Limit: &limit,
		Skip:  &skip,
	}

	productsCursor, err := productsCollection.Find(context.TODO(), bson.M{"name": bson.M{"$regex": "(?i)" + searchInput}}, findOptions)
	if err != nil {
		log.Println("Error getting products cursor:", err)
		return []models.Product{}
	}

	var products []models.Product

	if err := productsCursor.All(context.TODO(), &products); err != nil {
		log.Println("Error getting search prodcucts:", err)
		return []models.Product{}
	}

	return products
}
