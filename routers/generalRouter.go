package routers

import (
	"net/http"

	"github.com/gorilla/mux"
)

func StaticFilesRoute(r *mux.Router) {
	fs := http.FileServer(http.Dir("campmart-website/pub"))
	r.Handle("/campmart-website/pub/", http.StripPrefix("/campmart-website/pub/", fs))
}
