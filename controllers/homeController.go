package controllers

import (
	"campmart/helpers"
	"log"
	"net/http"
)

var tpl = helpers.LoadTemplate()

func Get() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := tpl.ExecuteTemplate(w, "index.html", nil); err != nil {
			log.Fatal("ExexcuteTemplate error:", err)
		}
	}
}
