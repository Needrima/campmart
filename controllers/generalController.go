package controllers

import (
	"campmart/helpers"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

var tpl = helpers.LoadTemplate()

func RedirectToHome() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		http.Redirect(w, r, "/home", http.StatusSeeOther)
	}
}
