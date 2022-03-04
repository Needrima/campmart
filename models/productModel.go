package models

type Product struct {
	Name        string   `json:"name"`
	Price       int      `json:"price"`
	Types       []string `json:"type"` //comma seperated values
	ImageNames  []string `json:"imagenames"`
	Description string   `json:"description"`
	Properties  []string `json:"properties"`
	Rating      int      `json:"rating"`
	Brand       string   `json:"brand"`
}
