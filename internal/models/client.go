package models

import "time"

type Cliente struct {
	ID            int       `json:"id"`
	Nombre        string    `json:"nombre"`
	Cedula        string    `json:"cedula"`
	Email         string    `json:"email"`
	Telefono      string    `json:"telefono"`
	Membresia     string    `json:"membresia"`
	FechaRegistro time.Time `json:"fecha_registro"`
}
