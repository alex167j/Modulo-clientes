package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"

	"modulo-clientes/internal/handlers"
	"modulo-clientes/internal/storage"
)

func main() {
	clienteStore := storage.NewClienteStore()
	reservaStore := storage.NewReservaStore()
	pagoStore    := storage.NewPagoStore()

	clientes := handlers.NewClienteHandler(clienteStore)
	reservas := handlers.NewReservaHandler(reservaStore)
	pagos    := handlers.NewPagoHandler(pagoStore)

	r := chi.NewRouter()

	r.Route("/api/v1", func(r chi.Router) {

		r.Route("/clientes", func(r chi.Router) {
			r.Post("/", clientes.Create)
			r.Get("/", clientes.GetAll)
			r.Get("/{id}", clientes.GetByID)
			r.Patch("/{id}", clientes.Update)
			r.Delete("/{id}", clientes.Delete)
		})

		r.Route("/reservas", func(r chi.Router) {
			r.Post("/", reservas.Create)
			r.Get("/", reservas.GetAll)
			r.Get("/{id}", reservas.GetByID)
			r.Patch("/{id}/cancelar", reservas.Cancelar)
		})

		r.Route("/pagos", func(r chi.Router) {
			r.Post("/", pagos.Create)
			r.Get("/", pagos.GetAll)
		})
	})

	log.Println("Servidor iniciado en http://localhost:8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}
}
