package handler

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lagringatati/Final_BE3_GO/internal/domain"
	"github.com/lagringatati/Final_BE3_GO/internal/paciente"
	"github.com/lagringatati/Final_BE3_GO/pkg/web"
)

type pacienteHandler struct {
	s paciente.Service
}

// NewPacienteHandler crea un nuevo controller de pacientes
func NewPacienteHandler(s paciente.Service) *pacienteHandler {
	return &pacienteHandler{
		s: s,
	}
}

// POST --> CreatePaciente crea un nuevo paciente
func (h *pacienteHandler) CreatePaciente() gin.HandlerFunc {
	return func(c *gin.Context) {
		var paciente domain.Paciente

		err := c.ShouldBindJSON(&paciente)
		if err != nil {
			web.Failure(c, 400, errors.New("JSON DE PACIENTE INVALIDO"))
			return
		}
		p, err := h.s.CreatePaciente(paciente)
		if err != nil {
			web.Failure(c, 400, err)
			return
		}
		web.Success(c, 201, p, "PACIENTE CREADO CORRECTAMENTE")
	}
}

// GET --> GetPacienteById obtiene un paciente por id
func (h *pacienteHandler) GetPacienteByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("idPaciente")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("ID INVALIDO"))
			return
		}
		paciente, err := h.s.GetPacienteByID(id)
		if err != nil {
			web.Failure(c, 404, errors.New("PACIENTE NO ENCONTRADO"))
			return
		}
		web.Success(c, 200, paciente, "PACIENTE OBTENIDO POR ID")
	}
}

// validateFieldsPaciente valida que los campos sean correctos
func validateFieldsPaciente(paciente *domain.Paciente) (bool, error) {
	switch {
	case paciente.NombrePaciente == "" || paciente.ApellidoPaciente == "" || paciente.DomicilioPaciente == "" || paciente.DniPaciente == "" || paciente.FechaDeAltaPaciente == "" /* || paciente.FechaDeAltaPaciente > time.Now() */ :
		return false, errors.New("LOS CAMPOS NO PUEDEN SER VACIOS NI LA FECHA DE ALTA POSTERIOR A HOY")
	}
	return true, nil
}

// PUT --> UpdatePaciente actualiza un paciente
func (h *pacienteHandler) UpdatePaciente() gin.HandlerFunc {
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
		idParam := c.Param("idPaciente")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("ID INVALIDO"))
			return
		}
		_, err = h.s.GetPacienteByID(id)
		if err != nil {
			web.Failure(c, 404, errors.New("ID DE PACIENTE INEXISTENTE"))
			return
		}
		if err != nil {
			web.Failure(c, 409, err)
			return
		}
		var paciente domain.Paciente
		err = c.ShouldBindJSON(&paciente)
		if err != nil {
			web.Failure(c, 400, errors.New("JSON INVALIDO"))
			return
		}
		valid, err := validateFieldsPaciente(&paciente)
		if !valid {
			web.Failure(c, 400, err)
			return
		}
		p, err := h.s.UpdatePaciente(id, paciente)
		if err != nil {
			web.Failure(c, 409, err)
			return
		}
		web.Success(c, 200, p, "PACIENTE ACTUALIZADO CON EXITO")
	}
}

// PATCH --> UpdatePacienteForField actualiza un paciente por alguno de sus campos
func (h *pacienteHandler) UpdatePacienteForField() gin.HandlerFunc {
	type Request struct {
		NombrePaciente      string `json:"nombrePaciente,omitempty"`
		ApellidoPaciente    string `json:"apellidoPaciente,omitempty"`
		DomicilioPaciente   string `json:"domicilioPaciente,omitempty"`
		DniPaciente         string `json:"dniPaciente,omitempty"`
		FechaDeAltaPaciente string `json:"fechaDeAltaPacientePaciente,omitempty"`
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
		idParam := c.Param("idPaciente")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("ID INVALIDO"))
			return
		}
		_, err = h.s.GetPacienteByID(id)
		if err != nil {
			web.Failure(c, 404, errors.New("ID DE PACIENTE INEXISTENTE"))
			return
		}
		if err := c.ShouldBindJSON(&r); err != nil {
			web.Failure(c, 400, errors.New("JSON INVALIDO"))
			return
		}
		update := domain.Paciente{
			NombrePaciente:      r.NombrePaciente,
			ApellidoPaciente:    r.ApellidoPaciente,
			DomicilioPaciente:   r.DomicilioPaciente,
			DniPaciente:         r.DniPaciente,
			FechaDeAltaPaciente: r.FechaDeAltaPaciente,
		}
		p, err := h.s.UpdatePaciente(id, update)
		if err != nil {
			web.Failure(c, 409, err)
			return
		}
		web.Success(c, 200, p, "PACIENTE ACTUALIZADO CON EXITO")
	}
}

// DEL --> DeletePaciente elimina un paciente
func (h *pacienteHandler) DeletePaciente() gin.HandlerFunc {
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
		idParam := c.Param("idPaciente")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("ID INVALIDO"))
			return
		}
		_, err = h.s.GetPacienteByID(id)
		if err != nil {
			web.Failure(c, 404, errors.New("ID DE PACIENTE INEXISTENTE"))
			return
		}
		err = h.s.DeletePaciente(id)
		if err != nil {
			web.Failure(c, 404, err)
			return
		}
		web.Success(c, 200, nil, "PACIENTE ELIMINADO CORRECTAMENTE")
	}
}
