package routers

import (
	controller "campmart/controllers"
	// "net/http"
	"github.com/julienschmidt/httprouter"
)

func ShopRoutes(r *httprouter.Router) {
	r.GET("/shop", controller.ShopGet())
}
