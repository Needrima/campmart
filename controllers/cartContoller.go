package controllers

import (
	"campmart/database"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
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

		cartItems := database.TemporaryCartDB[cartCookie.Value]

		if len(cartItems) == 0 {
			if err := tpl.ExecuteTemplate(w, "empty-cart.html", nil); err != nil {
				log.Fatal("ExecuteTemplate error:", err)
			}
			return
		}

		if err := tpl.ExecuteTemplate(w, "cart.html", nil); err != nil {
			log.Fatal("ExexcuteTemplate error:", cartItems)
		}
	}
}
