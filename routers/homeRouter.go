package routers

import (
	controller "campmart/controllers"
	// "net/http"
	"github.com/julienschmidt/httprouter"
)

func HomeRoutes(r *httprouter.Router) {
	r.GET("/home", controller.HomeGet())
}