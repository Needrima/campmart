package routers

import (
	controller "campmart/controllers"
	// "net/http"
	"github.com/julienschmidt/httprouter"
)

func RegisterBlogRoutes(r *httprouter.Router) {
	r.GET("/blog", controller.BlogGet())
	r.GET("/single-blog/:id", controller.SingleBlogGet())
	r.POST("/comment/:id", controller.AddNewComment())
}
