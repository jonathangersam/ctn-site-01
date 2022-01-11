package home

import (
	"html/template"
	"net/http"
)

type data struct {
	Name string
}

func Handler() http.HandlerFunc {
	tmpl := template.Must(template.ParseFiles("web/templates/home.html"))

	d := data{
		Name: "Jonathan Gersam S. Lopez",
	}

	return func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w, d)
	}
}
