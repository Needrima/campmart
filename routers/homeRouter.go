package routers

import (
	controller "campmart/controllers"
	"github.com/gorilla/mux"
)

func HomeRoute(r *mux.Router) {
	r.HandleFunc("/home", controller.Get)
}
