package middlewares

import (
	"campmart/database"
	"campmart/helpers"
	"campmart/models"
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CreateNewSubscriber return a new subscriber from email from form input.
// See type Subscriber
func CreateNewSubscriber(r *http.Request) (models.Subscriber, error) {

	bs, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("Error reading email from response body:", err)
		return models.Subscriber{}, errors.New("something went wrong try again later")
	}

	email, exp := string(bs), `^[a-zA-Z0-9.!#$%&'*+/=?^_{|}~-]+@[a-zA-Z0-9-]+(?:\.[a-zA-Z0-9-]+)*$`
	if !helpers.ValidFormInput(email, exp) {
		fmt.Println("Invaid email")
		return models.Subscriber{}, errors.New("invalid email, check and try again")
	}

	subscribersCollection := database.GetDatabaseCollection("subscribers")

	subscribersCursor, err := subscribersCollection.Find(context.TODO(), bson.M{})
	if err != nil {
		fmt.Println("subscriber cursor error:", err)
		return models.Subscriber{}, errors.New("something went wrong, try again later")
	}

	for subscribersCursor.Next(context.TODO()) {
		var s models.Subscriber

		if err := subscribersCursor.Decode(&s); err != nil {
			fmt.Println("Error getting subscriber:", err)
			continue
		}

		if s.Email == email {
			fmt.Println("User already a subscriber")
			return models.Subscriber{}, errors.New("you are already a subscriber")
		}
	}

	databaseID := primitive.NewObjectID()
	id := databaseID.Hex()
	newSubcriber := models.Subscriber{
		DatabaseID: databaseID,
		Id:         id,
		Email:      email,
	}

	return newSubcriber, nil
}
