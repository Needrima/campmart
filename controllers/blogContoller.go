package controllers

import (
	// "campmart/helpers"

	"campmart/middlewares"
	"campmart/models"
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

//BlogGet gets all blog first 3 latest blog posts serves blog.html to browser
func BlogGet() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		pageNumber := 0
		blogPosts := middlewares.GetBlogposts(pageNumber)

		if len(blogPosts) == 0 {
			http.Error(w, "somethin went wrong", http.StatusInternalServerError)
			return
		}

		blogPage := models.BlogPage{
			BlogPosts:  blogPosts,
			PageNumber: pageNumber,
		}

		if err := tpl.ExecuteTemplate(w, "blog.html", blogPage); err != nil {
			log.Fatal("ExexcuteTemplate error:", err)
		}
	}
}

// SingleBlogGet serves a single blog to single-blog.html
func SingleBlogGet() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		id := ps.ByName("id")
		blogPostAndOtherPost := middlewares.GetSinglePostAndSugestions(id)

		if err := tpl.ExecuteTemplate(w, "single-blog.html", blogPostAndOtherPost); err != nil {
			log.Fatal("ExexcuteTemplate error:", err)
		}
	}
}

func AddNewComment() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		id := ps.ByName("id")

		newCommentId, err := middlewares.AddNewCommentToPost(r, id)
		if err != nil {
			http.Error(w, "something went wrong", http.StatusInternalServerError)
			return
		}

		fmt.Println("New comment ID:", newCommentId)
		http.Redirect(w, r, "/single-blog/"+id, http.StatusSeeOther)
	}
}
