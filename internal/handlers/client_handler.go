package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	"modulo-clientes/internal/storage"
)

type ClienteHandler struct {
	store *storage.ClienteStore
}

func NewClienteHandler(store *storage.ClienteStore) *ClienteHandler {
	return &ClienteHandler{store: store}
}

func writeJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(data)
}

func getID(r *http.Request) (int, error) {
	return strconv.Atoi(chi.URLParam(r, "id"))
}
