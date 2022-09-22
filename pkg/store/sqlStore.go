package store

import (
	"database/sql"
	"fmt"

	"github.com/lagringatati/Final_BE3_GO/internal/domain"
)

type sqlStore struct {
	db *sql.DB
}

func NewSqlStore(db *sql.DB) StoreInterface {
	return &sqlStore{
		db: db,
	}
}

func (s *sqlStore) Create(odontologo domain.Odontologo) error {
	query := "INSERT INTO odontologos (idOdontologo, nombreOdontologo, apellidoOdontologo, matriculaOdontologo) VALUES (?, ?, ?, ?);"
	stmt, err := s.db.Prepare(query)
	if err != nil {
		return err
	}
	fmt.Println(odontologo)
	res, err := stmt.Exec(odontologo.IdOdontologo, odontologo.NombreOdontologo, odontologo.ApellidoOdontologo, odontologo.MatriculaOdontologo)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println(res)
	_, err = res.RowsAffected()
	if err != nil {
		return err
	}
	return nil
}

func (s *sqlStore) Read(id int) (domain.Odontologo, error) {
	var odontologo domain.Odontologo
	query := "SELECT * FROM odontologos WHERE idOdontologo = ?;"
	row := s.db.QueryRow(query, id)
	err := row.Scan(&odontologo.IdOdontologo, &odontologo.NombreOdontologo, &odontologo.ApellidoOdontologo, &odontologo.MatriculaOdontologo)
	if err != nil {
		return domain.Odontologo{}, err
	}
	return odontologo, nil
}

func (s *sqlStore) Update(odontologo domain.Odontologo) error {
	query := "UPDATE odontologos SET nombreOdontologo = ?, apellidoOdontologo = ?, matriculaOdontologo = ? WHERE idOdontologo = ?;"
	stmt, err := s.db.Prepare(query)
	if err != nil {
		return err
	}
	res, err := stmt.Exec(odontologo.NombreOdontologo, odontologo.ApellidoOdontologo, odontologo.MatriculaOdontologo, odontologo.IdOdontologo)
	if err != nil {
		return err
	}
	_, err = res.RowsAffected()
	if err != nil {
		return err
	}
	return nil
}

func (s *sqlStore) Delete(id int) error {
	query := "DELETE FROM odontologos WHERE idOdontologo = ?;"
	stmt, err := s.db.Prepare(query)
	if err != nil {
		return err
	}
	res, err := stmt.Exec(id)
	if err != nil {
		return err
	}
	_, err = res.RowsAffected()
	if err != nil {
		return err
	}
	return nil
}

/* func (s *sqlStore) Exists(matriculaOdontologo string) bool {
	var exists bool
	var id int
	query := "SELECT idOdontologo FROM odontologos WHERE matriculaOdontologo = ?;"
	row := s.db.QueryRow(query, matriculaOdontologo)
	err := row.Scan(&id)
	if err != nil {
		return false
	}
	if id > 0 {
		exists = true
	}
	return exists
}
*/
