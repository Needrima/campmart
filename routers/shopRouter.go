package routers

import (
	controller "campmart/controllers"
	// "net/http"
	"github.com/julienschmidt/httprouter"
)

func RegisterShopRoutes(r *httprouter.Router) {
	r.GET("/shop", controller.ShopGet())
	r.GET("/shop/next", controller.ChangePage())
	r.GET("/shop/previous", controller.ChangePage())
	r.POST("/searchtry", controller.SearchSuggestions())
}
