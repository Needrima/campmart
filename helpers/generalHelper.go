package helpers

import (
	"html/template"
)

// loads html templates
func LoadTemplate() *template.Template {
	tpl := template.Must(template.ParseGlob("website-templates/*.html"))

	return tpl
}

func FoundString(items []string, item string) bool {
	for _, it := range items {
		if item == it {
			return true
		}
	}

	return false
}
