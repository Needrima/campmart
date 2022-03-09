package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type CartItem struct {
	Id           string
	Image_name   string
	Name         string
	Price        int
	Quantity     int
	Types        []string
	SelectedType string
}

type Order struct {
	DatabaseID   primitive.ObjectID `bson:"_id"`
	Id           string             `bson:"id"`
	Cart_items   []CartItem         `bson:"cart_items"`
	Shipping_fee int                `bson:"shipping_fee, omitempty"`
}
