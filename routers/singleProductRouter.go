package routers

import (
	controller "campmart/controllers"
	// "net/http"
	"github.com/julienschmidt/httprouter"
)

func RegisterSingleProductRoutes(r *httprouter.Router) {
	r.GET("/single-product/:id", controller.SingleProductGet())

	r.POST("/single-product/add-to-cart", controller.SingeProductItemToCart())
}
