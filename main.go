// main.go
package main

import (
	"net/http"

	"github.com/Viru9029/URL-shortener-golang/handlers"
	"github.com/Viru9029/URL-shortener-golang/storage"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	storage := storage.NewMemoryStorage()

	r.HandleFunc("/shorten", handlers.ShortenURL(storage)).Methods("POST")
	r.HandleFunc("/{shortURL}", handlers.RedirectURL(storage)).Methods("GET")
	r.HandleFunc("/metrics", handlers.Metrics(storage)).Methods("GET")

	http.ListenAndServe(":8080", r)
}
