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

	router.RegisterGeneralRoute(r) // general routes like serving static files
	router.RegisterAdminRoute(r)   // admin routes like adding new products

	router.RegisterHomeRoutes(r)
	router.RegisterShopRoutes(r)
	router.RegisterBlogRoutes(r)
	router.RegisterContactRoutes(r)
	router.RegisterAboutRoutes(r)
	router.RegisterCartRoutes(r)
	router.RegisterSingleProductRoutes(r)
	router.RegisterSingleBlogRoutes(r)
	router.RegisterCheckoutRoutes(r)

	log.Fatal(http.ListenAndServe(":8080", r))
}
