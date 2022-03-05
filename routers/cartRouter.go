package routers

import (
	controller "campmart/controllers"
	// "net/http"
	"github.com/julienschmidt/httprouter"
)

func RegisterCartRoutes(r *httprouter.Router) {
	r.GET("/cart", controller.CartGet())
}
