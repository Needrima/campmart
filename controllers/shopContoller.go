package controllers

import (
	// "campmart/helpers"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// ShopGet gets 16 random products from database and serve to shop.html
func ShopGet() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		if err := tpl.ExecuteTemplate(w, "shop.html", nil); err != nil {
			log.Fatal("ExexcuteTemplate error:", err)
		}
	}
}
