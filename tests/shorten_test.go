// tests/shorten_test.go
package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Viru9029/URL-shortener-golang/handlers"
	"github.com/Viru9029/URL-shortener-golang/storage"
)

func TestShortenURL(t *testing.T) {
	store := storage.NewMemoryStorage()
	handler := handlers.ShortenURL(store)

	url := "https://www.example.com"
	reqBody, _ := json.Marshal(map[string]string{"Original": url})
	req, err := http.NewRequest("POST", "/shorten", bytes.NewBuffer(reqBody))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var response map[string]string
	json.NewDecoder(rr.Body).Decode(&response)
	if response["Original"] != url {
		t.Errorf("handler returned unexpected body: got %v want %v", response["Original"], url)
	}

	shortURL, exists := store.GetShortURL(url)
	if !exists {
		t.Errorf("expected short URL to be saved")
	}
	if response["Short"] != shortURL {
		t.Errorf("handler returned wrong short URL: got %v want %v", response["Short"], shortURL)
	}
}
