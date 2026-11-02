package handlers

import (
	"net/http"
	"path/filepath"
)

func ServeFile(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("file")
	http.ServeFile(w, r, filepath.Join("data", name))
}
