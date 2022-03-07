package controllers

import (
	"campmart/helpers"
	"campmart/middlewares"
	"campmart/models"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/julienschmidt/httprouter"
)

//gets product by id and 4 randoms products as suggestions for serve to single-product.html
func SingleProductGet() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		id := ps.ByName("id")

		productAndSuggestions := models.ProductAndSuggestions{
			Product:     middlewares.GetSingeProduct(id),
			Suggestions: middlewares.GetSuggestionsProducts(),
		}

		if err := tpl.ExecuteTemplate(w, "single-product.html", productAndSuggestions); err != nil {
			log.Fatal("ExexcuteTemplate error:", err)
		}
	}
}

func SingeProductItemToCart() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		bs, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println("Error reading response body:", err)
			w.Write([]byte("Could not add item to cart, something went wrong, try again later"))
			return
		}

		data := string(bs) // Ex. "675030nvjdkshg84ndj 3 small"
		fmt.Println("Data from frontend:", data)

		dataParts := strings.Split(data, " ")
		if len(dataParts) != 3 {
			log.Println("Data parts more than three")
			w.Write([]byte("Could not add item to cart, something went wrong, try again later"))
			return
		}

		id, selectedType := dataParts[0], dataParts[2]
		qty, err := strconv.Atoi(dataParts[1])
		if err != nil {
			log.Println("Invalid form input for qty")
			w.Write([]byte("Could not add item to cart, something went wrong, try again later"))
			return
		}

		product := middlewares.GetSingeProduct(id)

		cartItem := middlewares.GetCartItemFomProduct(product, qty, selectedType)
		fmt.Println(cartItem)

		if err := helpers.AddToTemporaryCartDatabase(cartItem); err != nil {
			w.Write([]byte(err.Error()))
			return
		}

		successMsg := fmt.Sprintf("Successfully added %v to cart", cartItem.Name)
		w.Write([]byte(successMsg))

	}
}
