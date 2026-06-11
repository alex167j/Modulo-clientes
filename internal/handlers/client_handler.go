package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	"modulo-clientes/internal/models"
	"modulo-clientes/internal/storage"
)

func writeJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(data)
}

func getID(r *http.Request) (int, error) {
	return strconv.Atoi(chi.URLParam(r, "id"))
}

// ── Cliente Handler ────────────────────────────────────────────────────────

type ClienteHandler struct {
	store *storage.ClienteStore
}

func NewClienteHandler(store *storage.ClienteStore) *ClienteHandler {
	return &ClienteHandler{store: store}
}

// POST /api/v1/clientes
func (h *ClienteHandler) Create(w http.ResponseWriter, r *http.Request) {
	var c models.Cliente
	if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
		http.Error(w, "json invalido", http.StatusBadRequest)
		return
	}

	created, err := h.store.Create(c)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	writeJSON(w, http.StatusCreated, created)
}

// GET /api/v1/clientes
func (h *ClienteHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, h.store.GetAll())
}

// GET /api/v1/clientes/{id}
func (h *ClienteHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, err := getID(r)
	if err != nil {
		http.Error(w, "id invalido", http.StatusBadRequest)
		return
	}

	c, err := h.store.GetByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	writeJSON(w, http.StatusOK, c)
}

// PATCH /api/v1/clientes/{id}
func (h *ClienteHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, err := getID(r)
	if err != nil {
		http.Error(w, "id invalido", http.StatusBadRequest)
		return
	}

	var c models.Cliente
	if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
		http.Error(w, "json invalido", http.StatusBadRequest)
		return
	}

	updated, err := h.store.Update(id, c)
	if err != nil {
		if err.Error() == "cliente no encontrado" {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	writeJSON(w, http.StatusOK, updated)
}

// DELETE /api/v1/clientes/{id}
func (h *ClienteHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := getID(r)
	if err != nil {
		http.Error(w, "id invalido", http.StatusBadRequest)
		return
	}

	if err := h.store.Delete(id); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

// ── Reserva Handler ────────────────────────────────────────────────────────

type ReservaHandler struct {
	store *storage.ReservaStore
}

func NewReservaHandler(store *storage.ReservaStore) *ReservaHandler {
	return &ReservaHandler{store: store}
}

// POST /api/v1/reservas
func (h *ReservaHandler) Create(w http.ResponseWriter, r *http.Request) {
	var rv models.Reserva
	if err := json.NewDecoder(r.Body).Decode(&rv); err != nil {
		http.Error(w, "json invalido", http.StatusBadRequest)
		return
	}

	created, err := h.store.Create(rv)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	writeJSON(w, http.StatusCreated, created)
}

// GET /api/v1/reservas
func (h *ReservaHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, h.store.GetAll())
}

// GET /api/v1/reservas/{id}
func (h *ReservaHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, err := getID(r)
	if err != nil {
		http.Error(w, "id invalido", http.StatusBadRequest)
		return
	}

	rv, err := h.store.GetByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	writeJSON(w, http.StatusOK, rv)
}

// PATCH /api/v1/reservas/{id}/cancelar
func (h *ReservaHandler) Cancelar(w http.ResponseWriter, r *http.Request) {
	id, err := getID(r)
	if err != nil {
		http.Error(w, "id invalido", http.StatusBadRequest)
		return
	}

	rv, err := h.store.Cancelar(id)
	if err != nil {
		if err.Error() == "reserva no encontrada" {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	writeJSON(w, http.StatusOK, rv)
}

// ── Pago Handler ───────────────────────────────────────────────────────────

type PagoHandler struct {
	store *storage.PagoStore
}

func NewPagoHandler(store *storage.PagoStore) *PagoHandler {
	return &PagoHandler{store: store}
}

// POST /api/v1/pagos
func (h *PagoHandler) Create(w http.ResponseWriter, r *http.Request) {
	var p models.Pago
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, "json invalido", http.StatusBadRequest)
		return
	}

	created, err := h.store.Create(p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	writeJSON(w, http.StatusCreated, created)
}

// GET /api/v1/pagos
func (h *PagoHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, h.store.GetAll())
}
