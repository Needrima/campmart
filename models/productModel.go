package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Product struct {
	DatabaseID  primitive.ObjectID `bson:"_id"`
	Id          string             `bson:"id"`
	Name        string             `bson:"name"`
	Price       int                `bson:"price"`
	Types       []string           `bson:"type"` //comma seperated values
	Image_names []string           `bson:"image_names"`
	Description string             `bson:"description"`
	Properties  []string           `bson:"properties, omitempty"` //comma seperated values
	Category    string             `bson:"category"`
	Rating      int                `bson:"rating"`
	Brand       string             `bson:"brand, omitempty"`
	Date_added  string             `bson:"date_added"`
	Seller      `bson:"seller"`
}

type ProductAndSuggestions struct {
	Product
	Suggestions []Product
}
