// handlers/shorten.go
package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Viru9029/URL-shortener-golang/models"
	"github.com/Viru9029/URL-shortener-golang/storage"
	"github.com/Viru9029/URL-shortener-golang/utils"
)

func ShortenURL(storage *storage.MemoryStorage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var request models.URL
		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		short, exists := storage.GetShortURL(request.Original)
		if !exists {
			short = utils.GenerateShortURL(request.Original)
			storage.SaveURL(request.Original, short)
		}

		response := models.URL{Original: request.Original, Short: short}
		json.NewEncoder(w).Encode(response)
	}
}
