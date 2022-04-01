package controllers

import (
	// "campmart/helpers"

	"campmart/middlewares"
	"campmart/models"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

//BlogGet gets all blog first 3 latest blog posts serves blog.html to browser
func BlogGet() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		pageNumber := 0
		blogPosts := middlewares.GetBlogposts(pageNumber)

		blogPage := models.BlogPage{
			BlogPosts:  blogPosts,
			PageNumber: pageNumber,
		}

		if err := tpl.ExecuteTemplate(w, "blog.html", blogPage); err != nil {
			log.Fatal("ExexcuteTemplate error:", err)
		}
	}
}
