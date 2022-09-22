package store

import "github.com/lagringatati/Final_BE3_GO/internal/domain"

type StoreInterface interface {
	// Read devuelve un odontologo por su id
	Read(id int) (domain.Odontologo, error)
	// Create agrega un nuevo odontologo
	Create(odontologo domain.Odontologo) error
	// Update actualiza un odontologo
	Update(odontologo domain.Odontologo) error
	// Delete elimina un odontologo
	Delete(id int) error
	// Exists verifica si un odontologo existe
	//Exists(matriculaOdontologo string) bool
}
