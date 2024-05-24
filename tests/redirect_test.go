// tests/redirect_test.go
package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Viru9029/URL-shortener-golang/handlers"
	"github.com/Viru9029/URL-shortener-golang/storage"
	"github.com/gorilla/mux"
)

func TestRedirectURL(t *testing.T) {
	store := storage.NewMemoryStorage()
	handler := handlers.RedirectURL(store)

	originalURL := "https://www.example.com"
	shortURL := "abcd"
	store.SaveURL(originalURL, shortURL)

	req, err := http.NewRequest("GET", "/"+shortURL, nil)
	if err != nil {
		t.Fatal(err)
	}

	req = mux.SetURLVars(req, map[string]string{"shortURL": shortURL})

	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusFound {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusFound)
	}

	if location := rr.Header().Get("Location"); location != originalURL {
		t.Errorf("handler returned wrong redirect URL: got %v want %v", location, originalURL)
	}
}
