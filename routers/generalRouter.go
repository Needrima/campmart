package routers

import (
	controller "campmart/controllers"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func GeneralRoute(r *httprouter.Router) {
	r.GET("/", controller.RedirectToHome())
	r.GET("/admin/new-product", controller.NewProduct())

	r.ServeFiles("/website-pub/*filepath", http.Dir("website-pub"))
}
