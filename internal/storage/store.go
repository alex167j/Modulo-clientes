package storage

import (
	"errors"
	"sync"
	"time"

	"modulo-clientes/internal/models"
)

// ── Cliente ────────────────────────────────────────────────────────────────

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
	if c.Cedula == "" {
		return models.Cliente{}, errors.New("la cedula es obligatoria")
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

func (s *ClienteStore) GetByID(id int) (models.Cliente, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	c, ok := s.clientes[id]
	if !ok {
		return models.Cliente{}, errors.New("cliente no encontrado")
	}
	return c, nil
}

func (s *ClienteStore) Update(id int, c models.Cliente) (models.Cliente, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	existing, ok := s.clientes[id]
	if !ok {
		return models.Cliente{}, errors.New("cliente no encontrado")
	}
	if c.Nombre == "" {
		return models.Cliente{}, errors.New("el nombre es obligatorio")
	}

	existing.Nombre = c.Nombre
	existing.Cedula = c.Cedula
	existing.Email = c.Email
	existing.Telefono = c.Telefono
	existing.Membresia = c.Membresia
	s.clientes[id] = existing
	return existing, nil
}

func (s *ClienteStore) Delete(id int) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if _, ok := s.clientes[id]; !ok {
		return errors.New("cliente no encontrado")
	}
	delete(s.clientes, id)
	return nil
}

// ── Reserva ────────────────────────────────────────────────────────────────

type ReservaStore struct {
	mutex    sync.RWMutex
	reservas map[int]models.Reserva
	nextID   int
}

func NewReservaStore() *ReservaStore {
	return &ReservaStore{
		reservas: make(map[int]models.Reserva),
		nextID:   1,
	}
}

func (s *ReservaStore) Create(r models.Reserva) (models.Reserva, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if r.ClienteID == 0 {
		return models.Reserva{}, errors.New("el cliente_id es obligatorio")
	}
	if r.Duracion <= 0 {
		return models.Reserva{}, errors.New("la duracion debe ser mayor a cero")
	}

	r.ID = s.nextID
	r.Estado = "activa"
	r.FechaHora = time.Now().UTC()
	s.reservas[r.ID] = r
	s.nextID++
	return r, nil
}

func (s *ReservaStore) GetAll() []models.Reserva {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	result := make([]models.Reserva, 0, len(s.reservas))
	for _, r := range s.reservas {
		result = append(result, r)
	}
	return result
}

func (s *ReservaStore) GetByID(id int) (models.Reserva, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	r, ok := s.reservas[id]
	if !ok {
		return models.Reserva{}, errors.New("reserva no encontrada")
	}
	return r, nil
}

func (s *ReservaStore) Cancelar(id int) (models.Reserva, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	r, ok := s.reservas[id]
	if !ok {
		return models.Reserva{}, errors.New("reserva no encontrada")
	}
	if r.Estado == "cancelada" {
		return models.Reserva{}, errors.New("la reserva ya esta cancelada")
	}

	r.Estado = "cancelada"
	s.reservas[id] = r
	return r, nil
}

// ── Pago ───────────────────────────────────────────────────────────────────

type PagoStore struct {
	mutex  sync.RWMutex
	pagos  map[int]models.Pago
	nextID int
}

func NewPagoStore() *PagoStore {
	return &PagoStore{
		pagos:  make(map[int]models.Pago),
		nextID: 1,
	}
}

func (s *PagoStore) Create(p models.Pago) (models.Pago, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if p.ClienteID == 0 {
		return models.Pago{}, errors.New("el cliente_id es obligatorio")
	}
	if p.Monto <= 0 {
		return models.Pago{}, errors.New("el monto debe ser mayor a cero")
	}
	if p.Concepto == "" {
		return models.Pago{}, errors.New("el concepto es obligatorio")
	}

	p.ID = s.nextID
	p.FechaHora = time.Now().UTC()
	s.pagos[p.ID] = p
	s.nextID++
	return p, nil
}

func (s *PagoStore) GetAll() []models.Pago {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	result := make([]models.Pago, 0, len(s.pagos))
	for _, p := range s.pagos {
		result = append(result, p)
	}
	return result
}
