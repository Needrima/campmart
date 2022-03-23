package controllers

import (
	// "campmart/helpers"
	"campmart/middlewares"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// ShopGet gets 16 random products from database and serve to shop.html
func ShopGet() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		products := middlewares.GetShopProducts()
		if err := tpl.ExecuteTemplate(w, "shop.html", products); err != nil {
			log.Fatal("ExexcuteTemplate error:", err)
		}
	}
}

// Updates search suggestions as a user tries to search for a product
func SearchSuggestions() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		bs, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println("Error reading request body:", err)
		}

		suggestions := middlewares.GetSearchSuggestions(string(bs))
		if len(suggestions) == 0 {
			json.NewEncoder(w).Encode([]string{"no results found"})
			return
		}

		json.NewEncoder(w).Encode(suggestions)
	}
}

func ChangePage() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		http.Redirect(w, r, "/shop", http.StatusSeeOther)
	}
}
