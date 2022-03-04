package controllers

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func HomeGet() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		if err := tpl.ExecuteTemplate(w, "index.html", nil); err != nil {
			log.Fatal("ExexcuteTemplate error:", err)
		}
	}
}
