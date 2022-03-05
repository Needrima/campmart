package controllers

import (
	"campmart/middlewares"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func HomeGet() httprouter.Handle {
	products := middlewares.GetHomeProducts()
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		if err := tpl.ExecuteTemplate(w, "index.html", nil); err != nil {
			log.Fatal("Exexcute Template error:", products)
		}
	}
}
