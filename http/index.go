package http

import (
	"html/template"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Redirect(w, r, "/error", http.StatusFound)
		return
	}
	t := template.Must(template.ParseFiles("templates/index.html"))
	t.Execute(w, nil)
}
