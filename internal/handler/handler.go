package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"url-shortener/internal/shortener"
	"url-shortener/internal/storage"
)

type URLHandler struct {
	store storage.URLStore
}

func NewURLHandler(store storage.URLStore) *URLHandler {
	return &URLHandler{store: store}
}

func (h *URLHandler) HandleShorten(w http.ResponseWriter, r *http.Request) {
	// 1. decode the incoming JSON
	var req ShortenRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	// 2. validate the URL
	if err := req.Validate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// 3. generate a key
	key := shortener.GenerateKey(6)

	// 4. save to storage
	if err := h.store.Save(key, req.URL); err != nil {
		http.Error(w, "failed to save url", http.StatusInternalServerError)
		return
	}

	// 5. forming a response
	shortURL := fmt.Sprintf("http://localhost:9091/r/%s", key)
	resp := ShortenResponse{
		ShortURL: shortURL,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)
}

func (h *URLHandler) HandleRedirect(w http.ResponseWriter, r *http.Request) {
	// get the key from the path
	key := r.PathValue("key")
	if key == "" {
		http.Error(w, "key is required", http.StatusBadRequest)
		return
	}

	// search for a record
	record, err := h.store.Get(key)
	if err != nil {
		http.Error(w, "url not found", http.StatusNotFound)
		return
	}

	// increase the visitor counter
	h.store.IncrementVisits(key)

	// make a redirect
	http.Redirect(w, r, record.Original, http.StatusFound)
}

func (h *URLHandler) HandleStats(w http.ResponseWriter, r *http.Request) {
	// get the key from the path
	key := r.PathValue("key")
	if key == "" {
		http.Error(w, "key is required", http.StatusBadRequest)
		return
	}

	// search for a record
	record, err := h.store.Get(key)
	if err != nil {
		http.Error(w, "url not found", http.StatusNotFound)
		return
	}

	// forming a response
	resp := StatsResponse{
		Original:  record.Original,
		CreatedAt: record.Created,
		Visits:    record.Visits,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)

}
