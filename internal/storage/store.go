package storage

import (
	"sync"

	"modulo-clientes/internal/models"
)

type ClienteStore struct {
	mutex    sync.RWMutex
	clientes map[int]models.Cliente
	nextID   int
}

func NewClienteStore() *ClienteStore {
	return &ClienteStore{
		clientes: make(map[int]models.Cliente),
		nextID:   1,
	}
}
