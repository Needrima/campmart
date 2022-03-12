package controllers

import (
	"campmart/database"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func CheckoutGet() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		cartCookie, err := r.Cookie("cart")
		if err == http.ErrNoCookie { // user has not added any item to cart
			http.Redirect(w, r, "/shop", http.StatusSeeOther)
			return
		}

		tempCartDB := database.TemporaryCartDB
		usersDB := tempCartDB[cartCookie.Value]
		if usersDB == nil || len(usersDB) == 0 { // user has no item in cart database
			http.Redirect(w, r, "/shop", http.StatusSeeOther)
			return
		}

		if err := tpl.ExecuteTemplate(w, "checkout.html", nil); err != nil {
			log.Fatal("ExexcuteTemplate error:", err)
		}
	}
}
