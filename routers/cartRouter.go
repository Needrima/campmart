package routers

import (
	controller "campmart/controllers"
	// "net/http"
	"github.com/julienschmidt/httprouter"
)

func RegisterCartRoutes(r *httprouter.Router) {
	r.GET("/cart", controller.CartGet())
	r.POST("/add-to-cart", controller.AddItemToCart())
	r.POST("/single/add-to-cart", controller.AddItemToCart())
	r.GET("/remove-item-from-cart/:id", controller.RemoveItemFromCart())
}
