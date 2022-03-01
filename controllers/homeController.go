package controllers

import (
	"campmart/helpers"
	"net/http"
)

var tpl = helpers.LoadTemplate()

func Get(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "index.html", nil)
}
