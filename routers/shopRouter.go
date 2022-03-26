package routers

import (
	controller "campmart/controllers"
	// "net/http"
	"github.com/julienschmidt/httprouter"
)

func RegisterShopRoutes(r *httprouter.Router) {
	r.GET("/shop", controller.ShopGet())
	r.GET("/shop/next", controller.NextOrPreviousPage())
	r.GET("/shop/previous", controller.NextOrPreviousPage())
	r.POST("/searchtry", controller.SearchSuggestions())
	r.POST("/search", controller.Search())
}
