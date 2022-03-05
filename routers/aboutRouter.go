package routers

import (
	controller "campmart/controllers"
	// "net/http"
	"github.com/julienschmidt/httprouter"
)

func RegisterAboutRoutes(r *httprouter.Router) {
	r.GET("/about", controller.AboutGet())
}
