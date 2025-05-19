package handler

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/borisdvlpr/epicshelter/internal/service"
)

type ApiHandler struct {
	svc *service.ApiService
}

func NewApiHandler(svc *service.ApiService) *ApiHandler {
	return &ApiHandler{
		svc: svc,
	}
}

func (h *ApiHandler) GetTodo(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := chi.URLParam(r, "id")
	cacheKey := fmt.Sprintf("todo:%s", id)

	// check cache
	if cachedData, err := h.svc.GetCache(ctx, cacheKey); err == nil {
		h.writeJSONResponse(w, http.StatusOK, "HIT", cachedData)
		log.Printf("%s: cache hit", cacheKey)
		return
	} else {
		log.Printf("%s: cache miss or error: %v", cacheKey, err)
	}

	// fetch from external API
	url := fmt.Sprintf("https://jsonplaceholder.typicode.com/todos/%s", id)
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Error fetching API: %v", err)
		http.Error(w, "Error fetching API", http.StatusNotFound)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error reading response body: %v", err)
		http.Error(w, "Error reading response body", http.StatusInternalServerError)
		return
	}

	// cache the response
	if err := h.svc.SetCache(ctx, cacheKey, body); err != nil {
		log.Printf("Error caching response: %v", err)
	} else {
		log.Printf("%s: added to cache", cacheKey)
	}

	// write response
	h.writeJSONResponse(w, http.StatusOK, "MISS", body)
}

// Helper function to write JSON responses
func (h *ApiHandler) writeJSONResponse(w http.ResponseWriter, status int, cacheStatus string, body []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Cache", cacheStatus)
	w.WriteHeader(status)
	w.Write(body)
}
