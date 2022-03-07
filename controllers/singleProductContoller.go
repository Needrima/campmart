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
	"time"

	"github.com/julienschmidt/httprouter"
	uid "github.com/satori/go.uuid"
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

		cartCookie, err := r.Cookie("cart")
		if err == http.ErrNoCookie {
			cookieId := uid.NewV4()
			cookie := &http.Cookie{
				Name:    "cart",
				Value:   cookieId.String(),
				Expires: time.Now().Add(time.Hour * 24),
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
