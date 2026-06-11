package models

import "time"

type Cliente struct {
	ID             int       `json:"id"`
	Nombre         string    `json:"nombre"`
	Cedula         string    `json:"cedula"`
	Email          string    `json:"email"`
	Telefono       string    `json:"telefono"`
	Membresia      string    `json:"membresia"`
	FechaRegistro  time.Time `json:"fecha_registro"`
}

type Reserva struct {
	ID        int       `json:"id"`
	ClienteID int       `json:"cliente_id"`
	FechaHora time.Time `json:"fecha_hora"`
	Duracion  int       `json:"duracion"`
	Estado    string    `json:"estado"`
}

type Pago struct {
	ID        int       `json:"id"`
	ClienteID int       `json:"cliente_id"`
	Monto     float64   `json:"monto"`
	Concepto  string    `json:"concepto"`
	FechaHora time.Time `json:"fecha_hora"`
	Metodo    string    `json:"metodo"`
}
