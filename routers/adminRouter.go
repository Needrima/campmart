package routers

import (
	controller "campmart/controllers"

	"github.com/julienschmidt/httprouter"
)

func RegisterAdminRoute(r *httprouter.Router) {
	r.GET("/admin/new-product", controller.NewProductGet())
	r.POST("/admin/new-product", controller.AddNewProduct())
	r.GET("/admin/new-blog", controller.NewBlogpostGet())
	r.POST("/admin/new-blog", controller.AddNewBlogpost())
}
