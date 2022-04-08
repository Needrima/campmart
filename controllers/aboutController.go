package controllers

import (
	"campmart/helpers"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

var tpl = helpers.LoadTemplate()

// AboutGet serves the about.html page to browser
func AboutGet() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		if err := tpl.ExecuteTemplate(w, "about.html", nil); err != nil {
			log.Fatal("ExexcuteTemplate error:", err)
		}
	}
}
