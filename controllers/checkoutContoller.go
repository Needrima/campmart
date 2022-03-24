package controllers

import (
	"campmart/database"
	"campmart/middlewares"
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// CheckoutGet get data from temporary cart database and place an order based on the items in the cart
func CheckoutGet() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		cartCookie, err := r.Cookie("cart")
		if err == http.ErrNoCookie { // user has not added any item to cart
			http.Redirect(w, r, "/shop", http.StatusSeeOther)
			return
		}

		tempCartDB := database.TemporaryCartDB
		usersCartDB := tempCartDB[cartCookie.Value]
		if len(usersCartDB) == 0 { // user has no item in cart database
			http.Redirect(w, r, "/shop", http.StatusSeeOther)
			return
		}

		if err := tpl.ExecuteTemplate(w, "checkout.html", nil); err != nil {
			log.Fatal("ExexcuteTemplate error:", err)
		}
	}
}

// PlaceNewOrder adds a new order to the database and sends email to buyer on successful placement
func PlaceNewOrder() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		if r.Method != http.MethodPost {
			http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
			return
		}

		cartCookie, err := r.Cookie("cart")
		if err == http.ErrNoCookie { // user has not added any item to cart
			http.Redirect(w, r, "/shop", http.StatusSeeOther)
			return
		}

		tempCartDB := database.TemporaryCartDB
		usersCartDB := tempCartDB[cartCookie.Value]
		fmt.Println("items in cart:", usersCartDB)

		newOrder, err := middlewares.CreateNewOrder(r, usersCartDB)
		if err != nil {
			if err := tpl.ExecuteTemplate(w, "checkout.html", err.Error()); err != nil {
				log.Fatal("ExexcuteTemplate error:", err)
			}
			return
		}

		newOrder.OrderTotal = middlewares.CalculateOrderTotal(newOrder)
		fmt.Println("New order:", newOrder)

		cartCookie.MaxAge = -1
		http.SetCookie(w, cartCookie)

		collection := database.GetDatabaseCollection("orders")

		insertOneResult, err := collection.InsertOne(context.TODO(), newOrder)
		if err != nil {
			fmt.Println("error inserting new order into database:", err.Error())
			http.Error(w, "something went wrong, try again later", http.StatusInternalServerError)
			return
		}

		if err := middlewares.SendMail(newOrder.BuyersEmail, "sucessfulOrderEmail.html", "Your order has been placed", newOrder); err != nil {
			log.Println("Error sending successful order mail:", err)
		}

		fmt.Printf("new order placed with id %v\n", insertOneResult.InsertedID)

		http.Redirect(w, r, "/shop", http.StatusSeeOther)
	}
}
