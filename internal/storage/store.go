package storage

import (
	"errors"
	"sync"
	"time"

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

func (s *ClienteStore) Create(c models.Cliente) (models.Cliente, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if c.Nombre == "" {
		return models.Cliente{}, errors.New("el nombre es obligatorio")
	}

	c.ID = s.nextID
	c.FechaRegistro = time.Now().UTC()
	s.clientes[c.ID] = c
	s.nextID++
	return c, nil
}

func (s *ClienteStore) GetAll() []models.Cliente {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	result := make([]models.Cliente, 0, len(s.clientes))
	for _, c := range s.clientes {
		result = append(result, c)
	}
	return result
}
