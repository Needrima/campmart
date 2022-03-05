package routers

import (
	controller "campmart/controllers"

	"github.com/julienschmidt/httprouter"
)

func RegisterAdminRoute(r *httprouter.Router) {
	r.GET("/admin/new-product", controller.NewProductGet())

	r.POST("/admin/new-product", controller.AddNewProduct())
}
