package handler

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lagringatati/Final_BE3_GO/internal/domain"
	"github.com/lagringatati/Final_BE3_GO/internal/turno"
	"github.com/lagringatati/Final_BE3_GO/pkg/web"
)

type turnoHandler struct {
	s turno.Service
}

// NewTurnoHandler crea un nuevo controller de turnos
func NewTurnoHandler(s turno.Service) *turnoHandler {
	return &turnoHandler{
		s: s,
	}
}

// POST --> CreateTurno crea un nuevo turno
func (h *turnoHandler) CreateTurno() gin.HandlerFunc {
	return func(c *gin.Context) {
		var turno domain.Turno

		err := c.ShouldBindJSON(&turno)
		if err != nil {
			web.Failure(c, 400, errors.New("JSON DE TURNO INVALIDO"))
			return
		}
		p, err := h.s.CreateTurno(turno)
		if err != nil {
			web.Failure(c, 400, err)
			return
		}
		web.Success(c, 201, p, "TURNO CREADO CORRECTAMENTE")
	}
}

// GET --> GetTurnoById obtiene un turno por id
func (h *turnoHandler) GetTurnoByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("idTurno")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("ID INVALIDO"))
			return
		}
		turno, err := h.s.GetTurnoByID(id)
		if err != nil {
			web.Failure(c, 404, errors.New("TURNO NO ENCONTRADO"))
			return
		}
		web.Success(c, 200, turno, "TURNO OBTENIDO POR ID")
	}
}

// validateFieldsTurno valida que los campos sean correctos
func validateFieldsTurno(turno *domain.Turno) (bool, error) {
	switch {
	case turno.DescripcionTurno == "" || turno.FechaTurno == "" || turno.IdOdontologo == "" || turno.IdPaciente == "":
		return false, errors.New("LOS CAMPOS NO PUEDEN SER VACIOS")
	}
	return true, nil
}

// PUT --> UpdateTurno actualiza un turno
func (h *turnoHandler) UpdateTurno() gin.HandlerFunc {
	return func(c *gin.Context) {
		/* 	token := c.GetHeader("TOKEN")
		if token == "" {
			web.Failure(c, 401, errors.New("TOKEN INCORRECTO"))
			return
		}
		if token != os.Getenv("TOKEN") {
			web.Failure(c, 401, errors.New("TOKEN INVALIDO"))
			return
		} */
		idParam := c.Param("idTurno")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("ID INVALIDO"))
			return
		}
		_, err = h.s.GetTurnoByID(id)
		if err != nil {
			web.Failure(c, 404, errors.New("ID DE TURNO INEXISTENTE"))
			return
		}
		if err != nil {
			web.Failure(c, 409, err)
			return
		}
		var turno domain.Turno
		err = c.ShouldBindJSON(&turno)
		if err != nil {
			web.Failure(c, 400, errors.New("JSON INVALIDO"))
			return
		}
		valid, err := validateFieldsTurno(&turno)
		if !valid {
			web.Failure(c, 400, err)
			return
		}
		p, err := h.s.UpdateTurno(id, turno)
		if err != nil {
			web.Failure(c, 409, err)
			return
		}
		web.Success(c, 200, p, "TURNO ACTUALIZADO CON EXITO")
	}
}

// PATCH --> UpdateTurnoForField actualiza un turno por alguno de sus campos
func (h *turnoHandler) UpdateTurnoForField() gin.HandlerFunc {
	type Request struct {
		DescripcionTurno string `json:"descripcionTurno,omitempty"`
		FechaTurno       string `json:"fechaTurno,omitempty"`
		IdOdontologo     string `json:"idOdontologo,omitempty"`
		IdPaciente       string `json:"idPaciente,omitempty"`
	}
	return func(c *gin.Context) {
		/* token := c.GetHeader("TOKEN")
		if token == "" {
			web.Failure(c, 401, errors.New("TOKEN INCORRECTO"))
			return
		}
		if token != os.Getenv("TOKEN") {
			web.Failure(c, 401, errors.New("TOKEN INVALIDO"))
			return
		} */
		var r Request
		idParam := c.Param("idTurno")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("ID INVALIDO"))
			return
		}
		_, err = h.s.GetTurnoByID(id)
		if err != nil {
			web.Failure(c, 404, errors.New("ID DE TURNO INEXISTENTE"))
			return
		}
		if err := c.ShouldBindJSON(&r); err != nil {
			web.Failure(c, 400, errors.New("JSON INVALIDO"))
			return
		}
		update := domain.Turno{
			DescripcionTurno: r.DescripcionTurno,
			FechaTurno:       r.FechaTurno,
			IdOdontologo:     r.IdOdontologo,
			IdPaciente:       r.IdPaciente,
		}
		p, err := h.s.UpdateTurno(id, update)
		if err != nil {
			web.Failure(c, 409, err)
			return
		}
		web.Success(c, 200, p, "TURNO ACTUALIZADO CON EXITO")
	}
}

// DEL --> DeleteTurno elimina un turno
func (h *turnoHandler) DeleteTurno() gin.HandlerFunc {
	//requiredToken := os.Getenv("API_TOKEN")

	return func(c *gin.Context) {

		/* 	token := c.GetHeader("api_token")
		if token == "" {
			web.Failure(c, 401, errors.New("TOKEN INCORRECTO"))
			return
		}
		if token != requiredToken {
			//c.AbortWithStatusJSON(401, gin.H{"error": "Invalid API token"})
			web.Failure(c, 401, errors.New("TOKEN INVALIDO"))
			return
		} */
		idParam := c.Param("idTurno")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("ID INVALIDO"))
			return
		}
		_, err = h.s.GetTurnoByID(id)
		if err != nil {
			web.Failure(c, 404, errors.New("ID DE TURNO INEXISTENTE"))
			return
		}
		err = h.s.DeleteTurno(id)
		if err != nil {
			web.Failure(c, 404, err)
			return
		}
		web.Success(c, 200, nil, "TURNO ELIMINADO CORRECTAMENTE")
	}
}

//POST Y GET ESPECIALES

// obtenerIdOdontologo obtiene el id del odontologo mediante su matricula
/* func (s *jsonStore) obtenerIdOdontologo(matricula string) (int, error) {
	odontologos, err := s.loadOdontologos()
	if err != nil {
		return 0, err
	}
	for _, odontologo := range odontologos {
		if odontologo.MatriculaOdontologo == matricula {
			return odontologo.IdOdontologo, nil
		}
	}
	return 0, errors.New("ODONTOLOGO INEXISTENTE")
}
*/
// obtenerIdPaciente obtiene el id del paciente mediante su DNI
/* func (s *jsonStore) obtenerIdPaciente(dni string) (int, error) {
	pacientes, err := s.loadPacientes()
	if err != nil {
		return 0, err
	}
	for _, paciente := range pacientes {
		if paciente.DniPaciente == dni {
			return paciente.IdPaciente, nil
		}
	}
	return 0, errors.New("ODONTOLOGO INEXISTENTE")
} */

// POST --> CreateTurnoSpecial crea un nuevo turno con el DNI del paciente y la matricula del odontologo
/* func (h *turnoHandler) CreateTurnoSpecial() gin.HandlerFunc {

	return func(c *gin.Context) {

		matO := c.Param("matriculaOdontologo")
		m, err := strconv.Atoi(matO)
		if err != nil {
			web.Failure(c, 400, errors.New("MATRICULA INVALIDA"))
			return
		}

		idO, err = h.s.obtenerIdOdontologo(m)
		if err != nil {
			web.Failure(c, 404, errors.New("MATRICULA INEXISTENTE"))
			return
		}

		dniP := c.Param("dniPaciente")
		dn, err := strconv.Atoi(dniP)
		if err != nil {
			web.Failure(c, 400, errors.New("DNI INVALIDO"))
			return
		}

		idP, err = h.s.obtenerIdPaciente(dn)
		if err != nil {
			web.Failure(c, 404, errors.New("DNI INEXISTENTE"))
			return
		}

		var turno domain.Turno

		turno.IdPaciente:=idP
		turno.IdOdontologo:=idO

		err := c.ShouldBindJSON(&turno)
		if err != nil {
			web.Failure(c, 400, errors.New("JSON DE TURNO INVALIDO"))
			return
		}

		p, err := h.s.CreateTurnoSpecial(turno)
		if err != nil {
			web.Failure(c, 400, err)
			return
		}
		web.Success(c, 201, p, "TURNO CREADO CORRECTAMENTE")
	}
}
*/
// GET --> GetTurnoByDni obtiene un turno por DNI del paciente
/* func (h *turnoHandler) GetTurnoByDni() gin.HandlerFunc {
	return func(c *gin.Context) {

		dniParam := c.Param("dniPaciente")
		dn, err := strconv.Atoi(dniParam)
		if err != nil {
			web.Failure(c, 400, errors.New("DNI INVALIDO"))
			return
		}

		//busco TODOS los turnos y hago un for buscando los q coincidan con el dni
		turno, err := h.s.GetTurnoByID(id)
		if err != nil {
			web.Failure(c, 404, errors.New("TURNO NO ENCONTRADO"))
			return
		}
		web.Success(c, 200, turno, "TURNO OBTENIDO POR ID")
	}
} */
