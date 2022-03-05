package routers

import (
	controller "campmart/controllers"
	// "net/http"
	"github.com/julienschmidt/httprouter"
)

func RegisterShopRoutes(r *httprouter.Router) {
	r.GET("/shop", controller.ShopGet())
}
