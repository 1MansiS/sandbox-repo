package handlers

import (
	"net/http"
	"path/filepath"
)

func ServeFile(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("file")
	// Clean("/"+name) collapses any ".." against the leading "/" so the
	// result can never climb above root, keeping the Join below inside data/.
	safe := filepath.Clean(string(filepath.Separator) + name)
	http.ServeFile(w, r, filepath.Join("data", safe))
}
