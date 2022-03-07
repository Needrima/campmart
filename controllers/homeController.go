package controllers

import (
	"campmart/middlewares"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

//get 8 random products from databse and serve to index.html
func HomeGet() httprouter.Handle {

	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		products := middlewares.GetHomeProducts()

		if err := tpl.ExecuteTemplate(w, "index.html", products); err != nil {
			log.Fatal("Exexcute Template error:", err)
		}
	}
}
