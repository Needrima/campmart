package routers

import (
	controller "campmart/controllers"
	// "net/http"
	"github.com/julienschmidt/httprouter"
)

func ContactRoutes(r *httprouter.Router) {
	r.GET("/contact", controller.ContactGet())
}
