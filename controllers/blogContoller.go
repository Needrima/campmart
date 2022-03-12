package controllers

import (
	// "campmart/helpers"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

//BlogGet gets all blog first 3 latest blog posts serves blog.html to browser
func BlogGet() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		if err := tpl.ExecuteTemplate(w, "blog.html", nil); err != nil {
			log.Fatal("ExexcuteTemplate error:", err)
		}
	}
}
