package controllers

import (
	"campmart/middlewares"
	"campmart/models"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// var tpl = helpers.LoadTemplate()

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
