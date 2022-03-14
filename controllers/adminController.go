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

// NewProductGet serves the new-product.html page to browser to add new product for sale
func NewProductGet() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		if err := tpl.ExecuteTemplate(w, "new-product.html", nil); err != nil {
			log.Fatal("Exexcute Template error:", err)
		}
	}
}

// AddNewProduct creates a new product from form submitted by admin and stores in database
func AddNewProduct() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		product, err := middlewares.CreateNewProduct(r)

		//check if string non numeric values is submited for numeric value form field
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		collection := database.GetDatabaseCollection("products")

		insertOneResult, err := collection.InsertOne(context.TODO(), product)
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		fmt.Printf("Inserted product id: %v", insertOneResult.InsertedID)

		successMsg := fmt.Sprintf("Successfully added product with id %v", insertOneResult.InsertedID)

		if err := tpl.ExecuteTemplate(w, "new-product.html", successMsg); err != nil {
			log.Fatal("Exexcute Template error:", err)
		}
	}
}
