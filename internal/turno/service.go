package turno

import (
	"github.com/lagringatati/Final_BE3_GO/internal/domain"
)

type Service interface {
	// GetTurnoByID busca un turno por su id
	GetTurnoByID(id int) (domain.Turno, error)
	// CreateTurno agrega un nuevo turno
	CreateTurno(p domain.Turno) (domain.Turno, error)
	// DeleteTurno elimina un turno
	DeleteTurno(id int) error
	// UpdateTurno actualiza un turno
	UpdateTurno(id int, p domain.Turno) (domain.Turno, error)
}

type service struct {
	r Repository
}

// NewService crea un nuevo servicio
func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) CreateTurno(p domain.Turno) (domain.Turno, error) {
	p, err := s.r.CreateTurno(p)
	if err != nil {
		return domain.Turno{}, err
	}
	return p, nil
}

func (s *service) GetTurnoByID(id int) (domain.Turno, error) {
	p, err := s.r.GetTurnoByID(id)
	if err != nil {
		return domain.Turno{}, err
	}
	return p, nil
}

func (s *service) UpdateTurno(id int, u domain.Turno) (domain.Turno, error) {
	p, err := s.r.GetTurnoByID(id)
	if err != nil {
		return domain.Turno{}, err
	}
	if u.DescripcionTurno != "" {
		p.DescripcionTurno = u.DescripcionTurno
	}
	if u.FechaTurno != "" {
		p.FechaTurno = u.FechaTurno
	}
	if u.IdOdontologo != "" {
		p.IdOdontologo = u.IdOdontologo
	}
	if u.IdPaciente != "" {
		p.IdPaciente = u.IdPaciente
	}
	p, err = s.r.UpdateTurno(id, p)
	if err != nil {
		return domain.Turno{}, err
	}
	return p, nil
}

func (s *service) DeleteTurno(id int) error {
	err := s.r.DeleteTurno(id)
	if err != nil {
		return err
	}
	return nil
}
