package models

import (
	"mime/multipart"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Subscriber struct {
	DatabaseID primitive.ObjectID `bson:"_id"`
	Id         string             `bson:"id"`
	Email      string             `bson:"email"`
}

type Image struct {
	File      multipart.File
	Name      string
	Extension string
}
