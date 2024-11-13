package http

import (
	"html/template"
	"net/http"
)

func Error(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("templates/errors.html"))
	t.Execute(w, nil)
}
