package routers

import (
	controller "campmart/controllers"
	// "net/http"
	"github.com/julienschmidt/httprouter"
)

func RegisterHomeRoutes(r *httprouter.Router) {
	r.GET("/home", controller.HomeGet())
}
