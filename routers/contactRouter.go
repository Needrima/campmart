package routers

import (
	controller "campmart/controllers"
	// "net/http"
	"github.com/julienschmidt/httprouter"
)

func RegisterContactRoutes(r *httprouter.Router) {
	r.GET("/contact", controller.ContactGet())
}
