package helpers

import (
	"html/template"
)

func LoadTemplate() *template.Template {
	tpl := template.Must(template.ParseGlob("website-templates/*.html"))

	return tpl
}
