// handlers/redirect.go
package handlers

import (
	"net/http"

	"github.com/Viru9029/URL-shortener-golang/storage"
	"github.com/gorilla/mux"
)

func RedirectURL(storage *storage.MemoryStorage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		shortURL := vars["shortURL"]

		original, exists := storage.GetOriginalURL(shortURL)
		if !exists {
			http.NotFound(w, r)
			return
		}

		http.Redirect(w, r, original, http.StatusFound)
	}
}
