package main

import (
	"log"
	"net/http"

	router "campmart/routers"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	router.StaticFilesRoute(r)
	router.HomeRoute(r)
	// router.AdminRoute(r)
	// router.BlogRoute(r)
	// router.CartRoute(r)
	// router.ShopRoute(r)

	log.Fatal(http.ListenAndServe(":8080", r))
}
