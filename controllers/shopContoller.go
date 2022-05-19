package controllers

import (
	"campmart/middlewares"
	"campmart/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/julienschmidt/httprouter"
)

// ShopGet gets 12 random products from database and serve to shop.html
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

		// if len(string(bs)) < 3 {
		// 	json.NewEncoder(w).Encode([]string{"no results found"})
		// 	return
		// }

		suggestions := middlewares.GetSearchSuggestions(string(bs))
		if len(suggestions) == 0 {
			json.NewEncoder(w).Encode([]string{"no results found"})
			return
		}

		json.NewEncoder(w).Encode(suggestions)
	}
}

func NextOrPreviousPage() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		http.Redirect(w, r, "/shop", http.StatusSeeOther)
	}
}

// Search searches for products matching the users search input
// a search input is stored in a cookie and updated on every new entry
func Search() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		searchInput := strings.TrimSpace(r.FormValue("searchInput"))
		fmt.Printf("Search input: %v, length: %v", searchInput, len(searchInput))
		pageNumber := 0

		searchCookie, err := r.Cookie("search")
		if err == http.ErrNoCookie {
			cookie := &http.Cookie{
				Name:    "search",
				Value:   searchInput,
				Expires: time.Now().Add(time.Hour * 24), // expires after a day
			}
			http.SetCookie(w, cookie)

			products := middlewares.GetProductsFromSearchInput(cookie.Value, pageNumber)
			if len(products) == 0 {
				http.Error(w, "no product found for search", http.StatusBadRequest)
				return
			}

			fmt.Println("Number of products in search:", len(products))

			searchPageData := models.SearchPage{
				Products:   products,
				PageNumber: pageNumber,
			}

			if err := tpl.ExecuteTemplate(w, "search.html", searchPageData); err != nil {
				log.Fatal("ExexcuteTemplate error:", err)
			}
			return
		}

		searchCookie.Value = searchInput
		http.SetCookie(w, searchCookie)

		products := middlewares.GetProductsFromSearchInput(searchCookie.Value, pageNumber)
		if len(products) == 0 {
			http.Error(w, "no product found for search", http.StatusBadRequest)
			return
		}

		searchPageData := models.SearchPage{
			Products:   products,
			PageNumber: pageNumber,
		}

		if err := tpl.ExecuteTemplate(w, "search.html", searchPageData); err != nil {
			log.Fatal("ExexcuteTemplate error:", err)
		}
	}
}

// NextOrPreviousSearch gets the next products in a search query determined by the page number
func NextOrPreviousSearch() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		pageNumber, _ := strconv.Atoi(ps.ByName("pageNumber"))

		searchCookie, err := r.Cookie("search")
		if err == http.ErrNoCookie {
			http.NotFound(w, r)
			return
		}

		products := middlewares.GetProductsFromSearchInput(searchCookie.Value, pageNumber)
		if len(products) == 0 {
			http.Error(w, "no product found for search", http.StatusBadRequest)
			return
		}

		searchPageData := models.SearchPage{
			Products:   products,
			PageNumber: pageNumber,
		}

		if err := tpl.ExecuteTemplate(w, "search.html", searchPageData); err != nil {
			log.Fatal("ExexcuteTemplate error:", err)
		}
	}
}
