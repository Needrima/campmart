package main

import (
	router "campmart/routers"
	"fmt"
	"log"
	"net/http"

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

	fmt.Println("Serving on port 8080. Visit localhost:8008 ....")
	server := http.Server{
		Addr:    ":8008",
		Handler: r,
	}

	log.Fatal(server.ListenAndServe())
}
