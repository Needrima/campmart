package controllers

import (
	"campmart/helpers"
	"campmart/middlewares"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

var tpl = helpers.LoadTemplate()

func RedirectToHome() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		http.Redirect(w, r, "/home", http.StatusSeeOther)
	}
}

func AddItemToCart() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		bs, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println("Error reading response body:", err)
			w.Write([]byte("Could not add item to cart, something went wrong, try again later"))
			return
		}

		id := string(bs)

		product := middlewares.GetSingeProduct(id)

		cartItem := middlewares.GetCartItemFomProduct(product, 1, "")

		if err := helpers.AddToTemporaryCartDatabase(cartItem); err != nil {
			w.Write([]byte(err.Error()))
			return
		}

		successMsg := fmt.Sprintf("Successfully added %v to cart", cartItem.Name)
		w.Write([]byte(successMsg))
	}
}
