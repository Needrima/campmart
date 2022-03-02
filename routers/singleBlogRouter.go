package routers

import (
	controller "campmart/controllers"
	// "net/http"
	"github.com/julienschmidt/httprouter"
)

func SingleBlogRoutes(r *httprouter.Router) {
	r.GET("/single-blog", controller.SingleBlogGet())
}
