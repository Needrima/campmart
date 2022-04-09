package controllers

import (
	// "campmart/helpers"
	"campmart/helpers"
	"campmart/middlewares"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"
)

// ContactGet serves contact.html to browser
func ContactGet() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		if err := tpl.ExecuteTemplate(w, "contact.html", nil); err != nil {
			log.Fatal("ExexcuteTemplate error:", err)
		}
	}
}

func SendUserMsg() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		name, exp := strings.TrimSpace(r.FormValue("senderName")), `^[a-zA-Z\s]{3,}$`
		if !helpers.ValidFormInput(name, exp) {
			http.Error(w, "invalid name, only english alphabets with atleast 3 chracters", http.StatusBadRequest)
			return
		}

		email, exp := strings.TrimSpace(r.FormValue("senderEmail")), `^[a-zA-Z0-9.!#$%&'*+/=?^_{|}~-]+@[a-zA-Z0-9-]+(?:\.[a-zA-Z0-9-]+)*$`
		if !helpers.ValidFormInput(email, exp) {
			http.Error(w, "invalid email, check email and try again", http.StatusBadRequest)
			return
		}

		subject, exp := strings.TrimSpace(r.FormValue("msgSubject")), `.*`
		if !helpers.ValidFormInput(subject, exp) {
			http.Error(w, "invalid subject, only english alphabets with atleast 3 chracters", http.StatusBadRequest)
			return
		}

		message, exp := strings.TrimSpace(r.FormValue("msg")), `.*`
		if !helpers.ValidFormInput(message, exp) {
			http.Error(w, "message body should be only english alphabets with atleast 3 chracters", http.StatusBadRequest)
			return
		}

		contactMsg := struct {
			Name, Email, Subject, Message string
		}{name, email, subject, message}

		mailSubject := fmt.Sprintf("%s sent a new contact mail", name)
		if err := middlewares.SendMail("oyebodeamirdeen@gmail.com", "sentContactmail.html", mailSubject, contactMsg); err != nil {
			http.Error(w, "something went wrong, try again later", http.StatusInternalServerError)
			return
		}

		if err := tpl.ExecuteTemplate(w, "contact.html", "mail sent successfully"); err != nil {
			log.Fatal("ExexcuteTemplate error:", err)
		}
	}
}
