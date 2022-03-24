package controllers

import (
	"campmart/database"
	"campmart/helpers"
	"campmart/middlewares"
	"context"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

var tpl = helpers.LoadTemplate()

// RedirectToHome redirects "/" to "/home"
func RedirectToHome() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		http.Redirect(w, r, "/home", http.StatusSeeOther)
	}
}

// SubscribeToNewsLetter registers a new subscriber and sends an email to subscriber on successful registration
func SubscribeToNewsLetter() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		subscriber, err := middlewares.CreateNewSubscriber(r)
		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}

		subscribersCollection := database.GetDatabaseCollection("subscribers")

		if err := middlewares.SendMail(subscriber.Email, "welcomeEmail.html", "Welcome to Campmart", nil); err != nil {
			fmt.Println(err)
		}

		insertOneResult, err := subscribersCollection.InsertOne(context.TODO(), subscriber)
		if err != nil {
			fmt.Println("Error inserting subscriber to subscribers collection:", err)
			w.Write([]byte("something went wrong, try again later"))
			return
		}

		fmt.Println("successfully added new subscriber with id:", insertOneResult.InsertedID)
		w.Write([]byte("you have succesfully subscribed to our newsletter"))
	}
}
