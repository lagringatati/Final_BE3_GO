package store

import (
	"encoding/json"
	"errors"
	"os"

	"github.com/lagringatati/Final_BE3_GO/internal/domain"
)

type jsonStore struct {
	pathToFile string
}

// loadOdontologos carga los odontologos desde un archivo json
func (s *jsonStore) loadOdontologos() ([]domain.Odontologo, error) {
	var odontologos []domain.Odontologo
	file, err := os.ReadFile(s.pathToFile)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal([]byte(file), &odontologos)
	if err != nil {
		return nil, err
	}
	return odontologos, nil
}

// saveOdontologos guarda los odontologos en un archivo json
func (s *jsonStore) saveOdontologos(odontologos []domain.Odontologo) error {
	bytes, err := json.Marshal(odontologos)
	if err != nil {
		return err
	}
	return os.WriteFile(s.pathToFile, bytes, 0644)
}

// NewJsonStore crea un nuevo store de odontologos
func NewJsonStore(path string) StoreInterface {
	_, err := os.Stat(path)
	if err != nil {
		panic(err)
	}
	return &jsonStore{
		pathToFile: path,
	}
}

func (s *jsonStore) Read(id int) (domain.Odontologo, error) {
	odontologos, err := s.loadOdontologos()
	if err != nil {
		return domain.Odontologo{}, err
	}
	for _, odontologo := range odontologos {
		if odontologo.IdOdontologo == id {
			return odontologo, nil
		}
	}
	return domain.Odontologo{}, errors.New("ODONTOLOGO INEXISTENTE")
}

func (s *jsonStore) Create(odontologo domain.Odontologo) error {
	odontologos, err := s.loadOdontologos()
	if err != nil {
		return err
	}
	odontologo.IdOdontologo = len(odontologos) + 1
	odontologos = append(odontologos, odontologo)
	return s.saveOdontologos(odontologos)
}

func (s *jsonStore) Update(odontologo domain.Odontologo) error {
	odontologos, err := s.loadOdontologos()
	if err != nil {
		return err
	}
	for i, p := range odontologos {
		if p.IdOdontologo == odontologo.IdOdontologo {
			odontologos[i] = odontologo
			return s.saveOdontologos(odontologos)
		}
	}
	return errors.New("ERROR AL ACTUALIZAR UN ODONTOLOGO")
}

func (s *jsonStore) Delete(id int) error {
	odontologos, err := s.loadOdontologos()
	if err != nil {
		return err
	}
	for i, p := range odontologos {
		if p.IdOdontologo == id {
			odontologos = append(odontologos[:i], odontologos[i+1:]...)
			return s.saveOdontologos(odontologos)
		}
	}
	return errors.New("ERROR AL ELIMINAR UN ODONTOLOGO")
}

/* func (s *jsonStore) Exists(matriculaOdontologo string) bool {
	odontologos, err := s.loadOdontologos()
	if err != nil {
		return false
	}
	for _, p := range odontologos {
		if p.MatriculaOdontologo == matriculaOdontologo {
			return true
		}
	}
	return false
}
*/