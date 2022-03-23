package middlewares

import (
	"campmart/helpers"
	"campmart/models"
	"errors"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CreateNewOrder return a new order with information from form from frontend
// See type Order struct{} in models folder
func CreateNewOrder(r *http.Request, usersCartDB map[string]models.CartItem) (models.Order, error) {
	buyersName, exp := r.FormValue("buyers_name"), `^[a-zA-Z]{3,}$`
	if !helpers.ValidFormInput(buyersName, exp) {
		log.Printf("invalid buyers name %v\n", buyersName)
		return models.Order{}, errors.New("invalid name, only english alphabets allowed with atleast three characters")
	}

	buyersEmail, exp := r.FormValue("buyers_email"), `^[a-zA-Z0-9.!#$%&'*+/=?^_{|}~-]+@[a-zA-Z0-9-]+(?:\.[a-zA-Z0-9-]+)*$`
	if !helpers.ValidFormInput(buyersEmail, exp) {
		log.Printf("invalid email address %v\n", buyersEmail)
		return models.Order{}, errors.New("invalid email address, check that email address is correct")
	}

	buyersNumber, exp := r.FormValue("buyers_number"), `^0\d{10}$` // only nigerian numbers allowed
	if !helpers.ValidFormInput(buyersNumber, exp) {
		log.Printf("invalid buyers phone number %v\n", buyersNumber)
		return models.Order{}, errors.New("invalid phone number, only nigerian numbers are allowed")
	}

	buyersNumber = "+234" + buyersNumber[1:]

	buyersAddress, exp := r.FormValue("buyers_address"), `^[a-zA-Z0-9\s,]{3,}$`
	if !helpers.ValidFormInput(buyersAddress, exp) {
		log.Printf("invalid delivery address %v\n", buyersNumber)
		return models.Order{}, errors.New("we suspect that you did not input your full delivery address")
	}

	optionalComment := r.FormValue("checkout_comment")

	databaseId := primitive.NewObjectID()
	id := databaseId.Hex()

	newOrder := models.Order{
		DatabaseID:      databaseId,
		Id:              id,
		CartItems:       usersCartDB,
		BuyersName:      buyersName,
		BuyersEmail:     buyersEmail,
		BuyersNumber:    buyersNumber,
		BuyersAddress: 	 buyersAddress,
		OptionalComment: optionalComment,
		ShippingFee:     0,
	}

	return newOrder, nil
}

// CacluateOrderTotal return the total amount for a new order
func CalculateOrderTotal(order models.Order) int {
	var cartTotal int
	for _, item := range order.CartItems {
		cartTotal += item.Price * item.Quantity
	}

	orderTotal := cartTotal + order.ShippingFee

	return orderTotal
}
