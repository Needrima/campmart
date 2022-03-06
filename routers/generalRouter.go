package routers

import (
	controller "campmart/controllers"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func RegisterGeneralRoute(r *httprouter.Router) {
	r.GET("/", controller.RedirectToHome())
	r.POST("/add-to-cart", controller.AddItemToCart())

	r.ServeFiles("/website-pub/*filepath", http.Dir("website-pub"))
}
