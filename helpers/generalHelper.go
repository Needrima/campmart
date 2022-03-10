package helpers

import (
	"html/template"
	"regexp"
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

func ValidFormInput(input, exp string) bool {
	return regexp.MustCompile(exp).MatchString(input)
}
