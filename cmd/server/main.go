package main

import (
	"fmt"
	"net/http"
	"url-shortener/internal/handler"
	"url-shortener/internal/storage"
)

func main() {
	store := storage.NewMapStore()

	handler := handler.NewURLHandler(store)

	mux := http.NewServeMux()

	mux.HandleFunc("POST /shorten", handler.HandleShorten)
	mux.HandleFunc("GET /r/{key}", handler.HandleRedirect)
	mux.HandleFunc("GET /stats/{key}", handler.HandleStats)

	fmt.Println("Server is starting on :9091...")
	err := http.ListenAndServe(":9091", mux)
	if err != nil {
		fmt.Println("Server failed to start:", err)
	}
}
