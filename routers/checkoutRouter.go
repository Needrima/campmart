package routers

import (
	controller "campmart/controllers"
	// "net/http"
	"github.com/julienschmidt/httprouter"
)

func RegisterCheckoutRoutes(r *httprouter.Router) {
	r.GET("/checkout", controller.CheckoutGet())
}
