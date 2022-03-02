package main

import (
	"html/template"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	tpl := template.Must(template.ParseGlob("templates/*.html"))

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if err := tpl.ExecuteTemplate(w, "index.html", nil); err != nil {
			log.Fatalf("Error parsing template: %v for route %v\n", err, r.URL.Path)
		}
	})

	r.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		if err := tpl.ExecuteTemplate(w, "about.html", nil); err != nil {
			log.Fatalf("Error parsing template: %v for route %v\n", err, r.URL.Path)
		}
	})

	r.HandleFunc("/cart", func(w http.ResponseWriter, r *http.Request) {
		if err := tpl.ExecuteTemplate(w, "cart.html", nil); err != nil {
			log.Fatalf("Error parsing template: %v for route %v\n", err, r.URL.Path)
		}
	})

	r.HandleFunc("/checkout", func(w http.ResponseWriter, r *http.Request) {
		if err := tpl.ExecuteTemplate(w, "checkout.html", nil); err != nil {
			log.Fatalf("Error parsing template: %v for route %v\n", err, r.URL.Path)
		}
	})

	r.HandleFunc("/contact", func(w http.ResponseWriter, r *http.Request) {
		if err := tpl.ExecuteTemplate(w, "contact.html", nil); err != nil {
			log.Fatalf("Error parsing template: %v for route %v\n", err, r.URL.Path)
		}
	})

	r.HandleFunc("/shop", func(w http.ResponseWriter, r *http.Request) {
		if err := tpl.ExecuteTemplate(w, "shop.html", nil); err != nil {
			log.Fatalf("Error parsing template: %v for route %v\n", err, r.URL.Path)
		}
	})

	r.HandleFunc("/single-blog", func(w http.ResponseWriter, r *http.Request) {
		if err := tpl.ExecuteTemplate(w, "single-blog.html", nil); err != nil {
			log.Fatalf("Error parsing template: %v for route %v\n", err, r.URL.Path)
		}
	})

	r.HandleFunc("/single-product", func(w http.ResponseWriter, r *http.Request) {
		if err := tpl.ExecuteTemplate(w, "single-product.html", nil); err != nil {
			log.Fatalf("Error parsing template: %v for route %v\n", err, r.URL.Path)
		}
	})

	r.HandleFunc("/blog", func(w http.ResponseWriter, r *http.Request) {
		if err := tpl.ExecuteTemplate(w, "blog.html", nil); err != nil {
			log.Fatalf("Error parsing template: %v for route %v\n", err, r.URL.Path)
		}
	})

	r.Handle("/pub/", http.StripPrefix("/pub/", http.FileServer(http.Dir("pub"))))

	log.Fatal(http.ListenAndServe(":8080", r))
}
