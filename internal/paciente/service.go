package paciente

import (
	"github.com/lagringatati/Final_BE3_GO/internal/domain"
)

type Service interface {
	// GetPacienteByID busca un paciente por su id
	GetPacienteByID(id int) (domain.Paciente, error)
	// CreatePaciente agrega un nuevo paciente
	CreatePaciente(p domain.Paciente) (domain.Paciente, error)
	// DeletePaciente elimina un paciente
	DeletePaciente(id int) error
	// UpdatePaciente actualiza un paciente
	UpdatePaciente(id int, p domain.Paciente) (domain.Paciente, error)
}

type service struct {
	r Repository
}

// NewService crea un nuevo servicio
func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) CreatePaciente(p domain.Paciente) (domain.Paciente, error) {
	p, err := s.r.CreatePaciente(p)
	if err != nil {
		return domain.Paciente{}, err
	}
	return p, nil
}

func (s *service) GetPacienteByID(id int) (domain.Paciente, error) {
	p, err := s.r.GetPacienteByID(id)
	if err != nil {
		return domain.Paciente{}, err
	}
	return p, nil
}

func (s *service) UpdatePaciente(id int, u domain.Paciente) (domain.Paciente, error) {
	p, err := s.r.GetPacienteByID(id)
	if err != nil {
		return domain.Paciente{}, err
	}
	if u.NombrePaciente != "" {
		p.NombrePaciente = u.NombrePaciente
	}
	if u.ApellidoPaciente != "" {
		p.ApellidoPaciente = u.ApellidoPaciente
	}
	if u.DomicilioPaciente != "" {
		p.DomicilioPaciente = u.DomicilioPaciente
	}
	if u.DniPaciente != "" {
		p.DniPaciente = u.DniPaciente
	}
	if u.FechaDeAltaPaciente != "" {
		p.FechaDeAltaPaciente = u.FechaDeAltaPaciente
	}
	p, err = s.r.UpdatePaciente(id, p)
	if err != nil {
		return domain.Paciente{}, err
	}
	return p, nil
}

func (s *service) DeletePaciente(id int) error {
	err := s.r.DeletePaciente(id)
	if err != nil {
		return err
	}
	return nil
}
