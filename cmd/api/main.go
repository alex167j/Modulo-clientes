package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"

	"modulo-clientes/internal/handlers"
	"modulo-clientes/internal/models"
	"modulo-clientes/internal/storage"
)

func main() {
	clienteStore := storage.NewClienteStore()
	reservaStore := storage.NewReservaStore()
	pagoStore    := storage.NewPagoStore()

	// Datos de ejemplo
	clienteStore.Create(models.Cliente{
		Nombre:    "Carlos Mendoza",
		Cedula:    "1304567890",
		Email:     "carlos.mendoza@gmail.com",
		Telefono:  "0991234567",
		Membresia: "mensual",
	})
	clienteStore.Create(models.Cliente{
		Nombre:    "Ana Loor Pincay",
		Cedula:    "1312345678",
		Email:     "ana.loor@hotmail.com",
		Telefono:  "0987654321",
		Membresia: "anual",
	})
	clienteStore.Create(models.Cliente{
		Nombre:    "Pedro Cedeño",
		Cedula:    "1398765432",
		Email:     "pedro.cedeno@yahoo.com",
		Telefono:  "0976543210",
		Membresia: "mensual",
	})

	reservaStore.Create(models.Reserva{
		ClienteID: 1,
		Duracion:  60,
	})
	reservaStore.Create(models.Reserva{
		ClienteID: 2,
		Duracion:  90,
	})

	pagoStore.Create(models.Pago{
		ClienteID: 1,
		Monto:     25.00,
		Concepto:  "Membresia mensual",
		Metodo:    "efectivo",
	})
	pagoStore.Create(models.Pago{
		ClienteID: 2,
		Monto:     150.00,
		Concepto:  "Membresia anual",
		Metodo:    "transferencia",
	})

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
