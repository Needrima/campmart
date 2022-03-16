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
	"strconv"
	"strings"

	"github.com/julienschmidt/httprouter"
	uid "github.com/satori/go.uuid"
)

// CartGet gets cart items from temporary database and serves to cart page.
// See var TemporaryCartDB in tempDB.go
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

		if err := tpl.ExecuteTemplate(w, "cart.html", cartItems); err != nil {
			log.Fatal("Exexcute Template cart.html:", err)
		}
	}
}

// AddItemToCart adds a new item to temporary cart database on two paths "/add-to-cart"
// and "/single-to-cart". See var TemporaryCartDB in tempDB.go
// "/add-to-cart" is called when users clicks the cart icon without specifying the quantity and type.
// "/single-to-cart" is called when users adds an item to cart from single-product.html(specifying the type and quantity).
// Both paths work with AJAX calls. Check add-to-cart.js and single-product-add-to-cart.js
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
		case "/single-to-cart":
			bs, err := ioutil.ReadAll(r.Body)
			if err != nil {
				log.Println("Error reading response body:", err)
				w.Write([]byte("Could not add item to cart, something went wrong, try again later"))
				return
			}

			data := string(bs) // Ex. "675030nvjdkshg84ndj 3 small"
			// fmt.Println("Data from frontend:", data)

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
			} else if qty < 1 { // if user tries to send a request without going through the form
				log.Println("Invalid form input for quantity")
				w.Write([]byte("Invalid form input for quantity"))
				return
			}

			product := middlewares.GetSingeProduct(id)
			if !helpers.FoundString(product.Types, selectedType) { // if user tries to send a request without going through the form
				log.Println("Invalid form input for selected type")
				w.Write([]byte("Invalid form input for selected type"))
				return
			}

			cartItem = middlewares.GetCartItemFomProduct(product, qty, selectedType)
		}

		// if cookie does not exist, create new cookie and use the value
		// to store cart items for the particular user in cart database
		// see tempDB.go for temporary DB and orderModel.go for type CartItem
		cartCookie, err := r.Cookie("cart")
		if err == http.ErrNoCookie {
			cookieID := uid.NewV4()
			cookie := &http.Cookie{
				Name:  "cart",
				Value: cookieID.String(),
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

// RemoveItemFromCart removes an item from cart by deleting it from the temporary cart database
// See var TemporaryCartDB in tempDB.go
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

		fmt.Println(len(tempDB))

		// if user has no item in cart anymore
		// delete cart from temporarydatabase to free up space
		if len(tempDB) == 0 {
			delete(database.TemporaryCartDB, cookieValue)
		}

		http.Redirect(w, r, "/cart", http.StatusSeeOther)
	}
}

// UpdateCartItems update quantity and type of items in the cart and before a user checksout the product
func UpdateCartItems() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		cartCookie, err := r.Cookie("cart")
		if err == http.ErrNoCookie {
			http.Redirect(w, r, "/home", http.StatusSeeOther)
			return
		}

		tempCartDB := database.TemporaryCartDB
		usersDB := tempCartDB[cartCookie.Value]
		if usersDB == nil || len(usersDB) == 0 { // user has no item in cart database
			http.Redirect(w, r, "/home", http.StatusSeeOther)
			return
		}

		// fmt.Println("Database before update:", usersDB)

		for id := range database.TemporaryCartDB[cartCookie.Value] {
			formerQty, formerType := usersDB[id].Quantity, usersDB[id].SelectedType
			updatedQty, updatedType := r.FormValue(id+"-qty"), r.FormValue(id+"-type")

			// if user tries to send a request without going through the form
			updatedQtyToInt, err := strconv.Atoi(updatedQty)
			if err != nil {
				log.Println("Invalid form input for qty")
				http.Error(w, "Invalid qunatity", http.StatusBadRequest)
				return
			} else if updatedQtyToInt < 1 {
				log.Println("Invalid form input for qty")
				http.Error(w, "Invalid qunatity", http.StatusBadRequest)
				return
			}

			if !helpers.FoundString(usersDB[id].Types, updatedType) {
				log.Println("Invalid form input for type")
				http.Error(w, "Invalid qunatity", http.StatusBadRequest)
				return
			}

			if formerQty != updatedQtyToInt {
				DbToUpdate := database.TemporaryCartDB[cartCookie.Value][id]
				DbToUpdate.Quantity = updatedQtyToInt
				database.TemporaryCartDB[cartCookie.Value][id] = DbToUpdate
			}

			if formerType != updatedType {
				DbToUpdate := database.TemporaryCartDB[cartCookie.Value][id]
				DbToUpdate.SelectedType = updatedType
				database.TemporaryCartDB[cartCookie.Value][id] = DbToUpdate
			}
		}

		// fmt.Println("Database after update:", usersDB)

		http.Redirect(w, r, "/checkout", http.StatusSeeOther)
	}
}
