package routers

import (
	controller "campmart/controllers"
	// "net/http"
	"github.com/julienschmidt/httprouter"
)

func BlogRoutes(r *httprouter.Router) {
	r.GET("/blog", controller.BlogGet())
}
