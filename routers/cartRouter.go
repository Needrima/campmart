package routers

import (
	controller "campmart/controllers"
	// "net/http"
	"github.com/julienschmidt/httprouter"
)

func CartRoutes(r *httprouter.Router) {
	r.GET("/cart", controller.CartGet())
}
