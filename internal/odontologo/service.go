package odontologo

import (
	"github.com/lagringatati/Final_BE3_GO/internal/domain"
)

type Service interface {
	// GetByID busca un odontologo por su id
	GetByID(id int) (domain.Odontologo, error)
	// Create agrega un nuevo odontologo
	Create(p domain.Odontologo) (domain.Odontologo, error)
	// Delete elimina un odontologo
	Delete(id int) error
	// Update actualiza un odontologo
	Update(id int, p domain.Odontologo) (domain.Odontologo, error)
}

type service struct {
	r Repository
}

// NewService crea un nuevo servicio
func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) Create(p domain.Odontologo) (domain.Odontologo, error) {
	p, err := s.r.Create(p)
	if err != nil {
		return domain.Odontologo{}, err
	}
	return p, nil
}

func (s *service) GetByID(id int) (domain.Odontologo, error) {
	p, err := s.r.GetByID(id)
	if err != nil {
		return domain.Odontologo{}, err
	}
	return p, nil
}

func (s *service) Update(id int, u domain.Odontologo) (domain.Odontologo, error) {
	p, err := s.r.GetByID(id)
	if err != nil {
		return domain.Odontologo{}, err
	}
	if u.NombreOdontologo != "" {
		p.NombreOdontologo = u.NombreOdontologo
	}
	if u.ApellidoOdontologo != "" {
		p.ApellidoOdontologo = u.ApellidoOdontologo
	}
	if u.MatriculaOdontologo != "" {
		p.MatriculaOdontologo = u.MatriculaOdontologo
	}
	p, err = s.r.Update(id, p)
	if err != nil {
		return domain.Odontologo{}, err
	}
	return p, nil
}

func (s *service) Delete(id int) error {
	err := s.r.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
