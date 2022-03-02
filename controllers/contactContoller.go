package controllers

import (
	// "campmart/helpers"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// var tpl = helpers.LoadTemplate()

func ContactGet() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		if err := tpl.ExecuteTemplate(w, "contact.html", nil); err != nil {
			log.Fatal("ExexcuteTemplate error:", err)
		}
	}
}
