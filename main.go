package main

import (
	router "campmart/routers"
	"log"
	"net/http"

	// "github.com/gorilla/mux"
	"github.com/julienschmidt/httprouter"
)

func main() {
	r := httprouter.New()

	router.GeneralRoute(r)

	router.HomeRoutes(r)
	router.ShopRoutes(r)
	router.BlogRoutes(r)
	router.ContactRoutes(r)
	router.AboutRoutes(r)
	router.CartRoutes(r)
	router.SingleProductRoutes(r)
	router.SingleBlogRoutes(r)
	router.CheckoutRoutes(r)

	log.Fatal(http.ListenAndServe(":8080", r))
}
