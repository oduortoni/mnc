package http

import (
	"net/http"
	"os"
)

func Static(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Path
	file := "." + url

	stat, err := os.Stat(file)
	if err != nil {
		http.Redirect(w, r, "/error", http.StatusFound)
		return
	}

	if stat.IsDir() {
		http.Redirect(w, r, "/error", http.StatusFound)
		return
	}

	http.ServeFile(w, r, file)
}
