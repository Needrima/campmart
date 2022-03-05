package controllers

import (
	"campmart/database"
	"campmart/middlewares"
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func NewProductGet() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		if err := tpl.ExecuteTemplate(w, "new-product.html", nil); err != nil {
			log.Fatal("Exexcute Template error:", err)
		}
	}
}

func AddNewProduct() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		product, err := middlewares.CreateNewProduct(r)
		if err == middlewares.ErrStringToInt {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if err == middlewares.ErrFormFile {
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
