package middlewares

import (
	"campmart/database"
	"campmart/models"
	"context"
	"fmt"
	"log"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
)

func GetShopProducts() []models.Product {
	productsCollection := database.GetDatabaseCollection("products")

	var products []models.Product

	sampleStage := bson.M{"$sample": bson.M{"size": 16}} // get 16 random products

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

	productsCursor, err := productsCollection.Find(context.TODO(), bson.M{})
	if err != nil {
		log.Println("Error getting products cursor:", err)
		return []string{}
	}

	var products []models.Product

	if err := productsCursor.All(context.TODO(), &products); err != nil {
		return []string{}
	}

	var suggestions []string

	for _, p := range products {
		if strings.Contains(strings.ToLower(p.Name), strings.ToLower(input)) {
			fmt.Println(p.Name)
			suggestions = append(suggestions, p.Name)
		}
	}

	fmt.Println("--------------------------")

	return suggestions
}
