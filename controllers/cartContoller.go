package controllers

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// get cart items from database and serves to cart page
func CartGet() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		if err := tpl.ExecuteTemplate(w, "cart.html", nil); err != nil {
			log.Fatal("ExexcuteTemplate error:", err)
		}
	}
}
