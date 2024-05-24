// tests/metrics_test.go
package tests

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Viru9029/URL-shortener-golang/handlers"
	"github.com/Viru9029/URL-shortener-golang/storage"
)

func TestMetrics(t *testing.T) {
	store := storage.NewMemoryStorage()
	handler := handlers.Metrics(store)

	// Populate the store
	urls := []string{"https://www.udemy.com", "https://www.youtube.com", "https://www.wikipedia.org"}
	for i, url := range urls {
		short := string(rune('a' + i))
		for j := 0; j <= i; j++ {
			store.SaveURL(url, short)
		}
	}

	req, err := http.NewRequest("GET", "/metrics", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var response map[string]int
	json.NewDecoder(rr.Body).Decode(&response)

	expected := map[string]int{
		"https://www.udemy.com":     1,
		"https://www.youtube.com":   2,
		"https://www.wikipedia.org": 3,
	}

	for k, v := range expected {
		if response[k] != v {
			t.Errorf("handler returned wrong metrics: got %v want %v", response[k], v)
		}
	}
}
