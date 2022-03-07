package controllers

import (
	"campmart/database"
	"campmart/helpers"
	"campmart/middlewares"
	"campmart/models"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	uid "github.com/satori/go.uuid"
)

var tpl = helpers.LoadTemplate()

// redirects "/" to "/home"
func RedirectToHome() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		http.Redirect(w, r, "/home", http.StatusSeeOther)
	}
}

//gets items on the "/single-product/add-to-cart" path with AJAX and add to temporary cart database
//check single-product-add-to-cart.js
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

		cartItem := middlewares.GetCartItemFomProduct(product, 1, product.Types[0])

		cartCookie, err := r.Cookie("cart")
		if err == http.ErrNoCookie {
			cookieId := uid.NewV4()
			cookie := &http.Cookie{
				Name:  "cart",
				Value: cookieId.String(),
			}
			http.SetCookie(w, cookie)

			if database.TemporaryCartDB[cookie.Value] == nil {
				database.TemporaryCartDB[cookie.Value] = map[string]models.CartItem{
					cartItem.Id: cartItem,
				}
			}

			successMsg := fmt.Sprintf("successfully added %v to cart", cartItem.Name)
			w.Write([]byte(successMsg))
			return
		}

		if database.TemporaryCartDB[cartCookie.Value] == nil {
			database.TemporaryCartDB[cartCookie.Value] = map[string]models.CartItem{}
		}

		if _, ok := database.TemporaryCartDB[cartCookie.Value][cartItem.Id]; !ok {
			database.TemporaryCartDB[cartCookie.Value][cartItem.Id] = cartItem

			successMsg := fmt.Sprintf("successfully added %v to cart", cartItem.Name)
			w.Write([]byte(successMsg))
			return
		}

		w.Write([]byte("item already in cart"))
	}
}
