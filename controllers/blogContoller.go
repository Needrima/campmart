package controllers

import (
	// "campmart/helpers"

	"campmart/middlewares"
	"campmart/models"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

//BlogGet gets the newest 3 blog posts in database sorted by time added serves to browser
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

// NextOrPreviousBlogPage get the next 3 blog posts or previous three blog post from the database.
// The action to execute i.e next or previou is determined by the page number.
// See type BlogPage in package models
func NextOrPreviousBlogPage() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		pageNumber, _ := strconv.Atoi(ps.ByName("pageNumber"))
		if pageNumber <= 0 {
			http.Redirect(w, r, "/blog", http.StatusSeeOther)
			return
		}

		blogPosts := middlewares.GetBlogposts(pageNumber)

		if len(blogPosts) == 0 {
			http.Redirect(w, r, "/blog", http.StatusSeeOther)
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

// AddNewComment adds a new comment to a blog post
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
