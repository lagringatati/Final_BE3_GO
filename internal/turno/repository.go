package turno

import (
	"errors"

	"github.com/lagringatati/Final_BE3_GO/internal/domain"
	"github.com/lagringatati/Final_BE3_GO/pkg/store"
)

type Repository interface {
	// GetTurnoByID busca un turno por su id
	GetTurnoByID(id int) (domain.Turno, error)
	// CreateTurno agrega un nuevo turno
	CreateTurno(p domain.Turno) (domain.Turno, error)
	// UpdateTurno actualiza un turno
	UpdateTurno(id int, p domain.Turno) (domain.Turno, error)
	// DeleteTurno elimina un turno
	DeleteTurno(id int) error
}

type repository struct {
	storage store.StoreInterface
}

// NewRepository crea un nuevo repositorio
func NewRepository(storage store.StoreInterface) Repository {
	return &repository{storage}
}

func (r *repository) CreateTurno(p domain.Turno) (domain.Turno, error) {

	err := r.storage.CreateTurno(p)
	if err != nil {
		return domain.Turno{}, errors.New("ERROR CREANDO TURNO")
	}
	return p, nil
}

func (r *repository) GetTurnoByID(id int) (domain.Turno, error) {
	turno, err := r.storage.ReadTurno(id)
	if err != nil {
		return domain.Turno{}, errors.New("TURNO INEXISTENTE")
	}
	return turno, nil
}

func (r *repository) UpdateTurno(id int, p domain.Turno) (domain.Turno, error) {
	err := r.storage.UpdateTurno(p)
	if err != nil {
		return domain.Turno{}, errors.New("ERROR ACTUALIZANDO AL TURNO")
	}
	return p, nil
}

func (r *repository) DeleteTurno(id int) error {
	err := r.storage.DeleteTurno(id)
	if err != nil {
		return err
	}
	return nil
}
