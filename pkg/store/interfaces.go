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
	// ReadPaciente devuelve un paciente por su id
	ReadPaciente(id int) (domain.Paciente, error)
	// CreatePaciente agrega un nuevo paciente
	CreatePaciente(paciente domain.Paciente) error
	// UpdatePaciente actualiza un paciente
	UpdatePaciente(paciente domain.Paciente) error
	// DeletePaciente elimina un paciente
	DeletePaciente(id int) error
	// ExistsPaciente verifica si un paciente existe
	//ExistsPaciente(dniPaciente string) bool
}
