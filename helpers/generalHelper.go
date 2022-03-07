package helpers

import (
	"campmart/models"
	"errors"
	"html/template"
)

// loads html templates
func LoadTemplate() *template.Template {
	tpl := template.Must(template.ParseGlob("website-templates/*.html"))

	return tpl
}

// add cart items to temporary database .. see type CartIte in package models
func AddToTemporaryCartDatabase(C models.CartItem) error {
	if _, ok := models.TemporaryCartDatabase[C.Id]; !ok {
		models.TemporaryCartDatabase[C.Id] = C
		return nil
	}

	return errors.New("item already added to cart")
}
