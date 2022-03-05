package routers

import (
	controller "campmart/controllers"
	// "net/http"
	"github.com/julienschmidt/httprouter"
)

func RegisterSingleProductRoutes(r *httprouter.Router) {
	r.GET("/single-product", controller.SingleProductGet())
}
