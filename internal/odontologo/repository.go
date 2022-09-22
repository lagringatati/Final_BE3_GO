package odontologo

import (
	"errors"

	"github.com/lagringatati/Final_BE3_GO/internal/domain"
	"github.com/lagringatati/Final_BE3_GO/pkg/store"
)

type Repository interface {
	// GetOdontologoByID busca un odontologo por su id
	GetByID(id int) (domain.Odontologo, error)
	// Create agrega un nuevo odontologo
	Create(p domain.Odontologo) (domain.Odontologo, error)
	// Update actualiza un odontologo
	Update(id int, p domain.Odontologo) (domain.Odontologo, error)
	// Delete elimina un odontologo
	Delete(id int) error
}

type repository struct {
	storage store.StoreInterface
}

// NewRepository crea un nuevo repositorio
func NewRepository(storage store.StoreInterface) Repository {
	return &repository{storage}
}

func (r *repository) Create(p domain.Odontologo) (domain.Odontologo, error) {

	err := r.storage.Create(p)
	if err != nil {
		return domain.Odontologo{}, errors.New("ERROR CREANDO ODONTOLOGO")
	}
	return p, nil
}

func (r *repository) GetByID(id int) (domain.Odontologo, error) {
	odontologo, err := r.storage.Read(id)
	if err != nil {
		return domain.Odontologo{}, errors.New("ODONTOLOGO INEXISTENTE")
	}
	return odontologo, nil
}

func (r *repository) Update(id int, p domain.Odontologo) (domain.Odontologo, error) {
	/* if r.storage.Exists(p.MatriculaOdontologo) {
		return domain.Odontologo{}, errors.New("LA MATRICULA YA EXISTE")
	} */
	err := r.storage.Update(p)
	if err != nil {
		return domain.Odontologo{}, errors.New("ERROR ACTUALIZANDO AL ODONTOLOGO")
	}
	return p, nil
}

func (r *repository) Delete(id int) error {
	err := r.storage.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
