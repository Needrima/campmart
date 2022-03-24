package main

import (
	router "campmart/routers"
	"fmt"
	"log"
	"net/http"
	"os"

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

	fmt.Println("Serving on port 8008. Visit localhost:8008 ....")
	port := os.Getenv("PORT")
	if port == "" {
		port = "8008"
	}

	server := http.Server{
		Addr:    ":" + port,
		Handler: r,
	}

	log.Fatal(server.ListenAndServe())
}
