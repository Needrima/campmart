package helpers

import (
	"campmart/models"
	"errors"
	"html/template"
)

func LoadTemplate() *template.Template {
	tpl := template.Must(template.ParseGlob("website-templates/*.html"))

	return tpl
}

func AddToTemporaryCartDatabase(C models.CartItem) error {
	if _, ok := models.TemporaryCartDatabase[C.Id]; !ok {
		models.TemporaryCartDatabase[C.Id] = C
		return nil
	}

	return errors.New("Item already added to cart")
}
