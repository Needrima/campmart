package helpers

import (
	"html/template"
	"regexp"
)

func incrementPageNumber(num int) int {
	num++
	return num
}

func decrementPageNumber(num int) int {
	if num == 0 {
		return num
	}

	num--
	return num
}

func cutContent(content string) string {
	return content[:203]
}

var templateFuncMap = template.FuncMap{
	"inc": incrementPageNumber,
	"dec": decrementPageNumber,
	"cut": cutContent,
}

// LoadTemplate loads html templates
func LoadTemplate() *template.Template {
	tpl := template.Must(template.New("").Funcs(templateFuncMap).ParseGlob("website-templates/*.html"))

	return tpl
}

// FoundString checks if a string exists in a slice of strings
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
