package controllers

import (
	"campmart/database"
	"campmart/middlewares"
	"campmart/models"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/julienschmidt/httprouter"
	uid "github.com/satori/go.uuid"
)

// get cart items from temmporary database database and serves to cart page
func CartGet() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		cartCookie, err := r.Cookie("cart")
		if err == http.ErrNoCookie {
			if err := tpl.ExecuteTemplate(w, "empty-cart.html", nil); err != nil {
				log.Fatal("ExecuteTemplate error:", err)
			}
			return
		}

		if database.TemporaryCartDB[cartCookie.Value] == nil {
			if err := tpl.ExecuteTemplate(w, "empty-cart.html", nil); err != nil {
				log.Fatal("ExecuteTemplate error:", err)
			}
			return
		}

		cartItems := database.TemporaryCartDB[cartCookie.Value]
		if len(cartItems) == 0 {
			if err := tpl.ExecuteTemplate(w, "empty-cart.html", nil); err != nil {
				log.Fatal("ExecuteTemplate error:", err)
			}
			return
		}

		fmt.Println("CartItems:", cartItems)

		if err := tpl.ExecuteTemplate(w, "cart.html", cartItems); err != nil {
			log.Fatal("Exexcute Template cart.html:", err)
		}
	}
}

//gets items on the "/single-product/add-to-cart" path with AJAX and add to temporary cart database
//check add-to-cart.js single-product-add-to-cart.js
func AddItemToCart() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		var cartItem models.CartItem

		switch r.URL.Path {
		// adding to cart through cart icon
		// qty = 1, selected type will be first type in product types
		case "/add-to-cart":
			bs, err := ioutil.ReadAll(r.Body)
			if err != nil {
				log.Println("Error reading response body:", err)
				w.Write([]byte("Could not add item to cart, something went wrong, try again later"))
				return
			}

			id := string(bs)

			product := middlewares.GetSingeProduct(id)

			cartItem = middlewares.GetCartItemFomProduct(product, 1, product.Types[0])

		// adding to cart through single-product.html page
		// qty and selected type will be specified by user
		case "/single/add-to-cart":
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

			cartItem = middlewares.GetCartItemFomProduct(product, qty, selectedType)
		}

		// if cookie does not exist, create new cookie and use the value
		// to store cart items for the particular user in cart database
		// see tempDB.go for temporary DB and orderModel.go for type CartItem
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

		// if cookie exist, get the value and use it to access the cart items for user
		// and store new item to cart item
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

func RemoveItemFromCart() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		id := ps.ByName("id")

		cartCookie, err := r.Cookie("cart")
		if err == http.ErrNoCookie {
			http.Error(w, "something went wrong, try again later", http.StatusBadRequest)
			return
		}

		cookieValue := cartCookie.Value

		tempDB := database.TemporaryCartDB[cookieValue]

		delete(tempDB, id)

		http.Redirect(w, r, "/cart", http.StatusSeeOther)
	}
}
