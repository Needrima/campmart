package controllers

import (
	// "campmart/helpers"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// SingleBlogGet serves a single blog to single-blog.html
func SingleBlogGet() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		if err := tpl.ExecuteTemplate(w, "single-blog.html", nil); err != nil {
			log.Fatal("ExexcuteTemplate error:", err)
		}
	}
}
