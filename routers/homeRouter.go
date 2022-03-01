package routers

import (
	controller "campmart/controllers"
	"net/http"

	"github.com/gorilla/mux"
)

func HomeRoute(r *mux.Router) {
	r.HandleFunc("/home", controller.Get()).Methods(http.MethodGet)
}
