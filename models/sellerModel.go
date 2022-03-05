package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Seller struct {
	DatabaseID   primitive.ObjectID `bson:"_id"`
	Id           string             `bson:"id"`
	Seller_name  string             `bson:"seller_name"`
	Seller_phone string             `bson:"seller_phone"`
	Seller_email string             `bson:"seller_email"`
}
