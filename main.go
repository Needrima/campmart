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

	router.HomeRoutes(r)
	router.ShopRoutes(r)
	router.BlogRoutes(r)
	router.ContactRoutes(r)
	router.AboutRoutes(r)
	router.CartRoutes(r)
	router.SingleProductRoutes(r)
	router.SingleBlogRoutes(r)
	router.CheckoutRoutes(r)

	router.GeneralRoute(r)

	// tpl := template.Must(template.ParseGlob("templates/*.html"))
	// r.GET("/", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// 	if err := tpl.ExecuteTemplate(w, "index.html", nil); err != nil {
	// 		log.Fatalf("Error parsing template: %v for route %v\n", err, r.URL.Path)
	// 	}
	// })

	// r.GET("/about", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// 	if err := tpl.ExecuteTemplate(w, "about.html", nil); err != nil {
	// 		log.Fatalf("Error parsing template: %v for route %v\n", err, r.URL.Path)
	// 	}
	// })

	// r.GET("/cart", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// 	if err := tpl.ExecuteTemplate(w, "cart.html", nil); err != nil {
	// 		log.Fatalf("Error parsing template: %v for route %v\n", err, r.URL.Path)
	// 	}
	// })

	// r.GET("/checkout", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// 	if err := tpl.ExecuteTemplate(w, "checkout.html", nil); err != nil {
	// 		log.Fatalf("Error parsing template: %v for route %v\n", err, r.URL.Path)
	// 	}
	// })

	// r.GET("/contact", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// 	if err := tpl.ExecuteTemplate(w, "contact.html", nil); err != nil {
	// 		log.Fatalf("Error parsing template: %v for route %v\n", err, r.URL.Path)
	// 	}
	// })

	// r.GET("/shop", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// 	if err := tpl.ExecuteTemplate(w, "shop.html", nil); err != nil {
	// 		log.Fatalf("Error parsing template: %v for route %v\n", err, r.URL.Path)
	// 	}
	// })

	// r.GET("/single-blog", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// 	if err := tpl.ExecuteTemplate(w, "single-blog.html", nil); err != nil {
	// 		log.Fatalf("Error parsing template: %v for route %v\n", err, r.URL.Path)
	// 	}
	// })

	// r.GET("/single-product", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// 	if err := tpl.ExecuteTemplate(w, "single-product.html", nil); err != nil {
	// 		log.Fatalf("Error parsing template: %v for route %v\n", err, r.URL.Path)
	// 	}
	// })

	// r.GET("/blog", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// 	if err := tpl.ExecuteTemplate(w, "blog.html", nil); err != nil {
	// 		log.Fatalf("Error parsing template: %v for route %v\n", err, r.URL.Path)
	// 	}
	// })

	// r.ServeFiles("/pub/*filepath", http.Dir("pub"))

	log.Fatal(http.ListenAndServe(":8080", r))
}
