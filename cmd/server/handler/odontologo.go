package handler

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lagringatati/Final_BE3_GO/internal/domain"
	"github.com/lagringatati/Final_BE3_GO/internal/odontologo"
	"github.com/lagringatati/Final_BE3_GO/pkg/web"
)

type odontologoHandler struct {
	s odontologo.Service
}

// NewOdontologoHandler crea un nuevo controller de odontologos
func NewOdontologoHandler(s odontologo.Service) *odontologoHandler {
	return &odontologoHandler{
		s: s,
	}
}

// POST --> CreateOdontologo crea un nuevo odontologo
func (h *odontologoHandler) CreateOdontologo() gin.HandlerFunc {
	return func(c *gin.Context) {
		var odontologo domain.Odontologo

		err := c.ShouldBindJSON(&odontologo)
		if err != nil {
			web.Failure(c, 400, errors.New("JSON DE ODONTOLOGO INVALIDO"))
			return
		}
		p, err := h.s.Create(odontologo)
		if err != nil {
			web.Failure(c, 400, err)
			return
		}
		web.Success(c, 201, p, "ODONTOLOGO CREADO CORRECTAMENTE")
	}
}

// GET --> GetOdontologoById obtiene un odontologo por id
func (h *odontologoHandler) GetOdontologoByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("idOdontologo")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("ID INVALIDO"))
			return
		}
		odontologo, err := h.s.GetByID(id)
		if err != nil {
			web.Failure(c, 404, errors.New("ODONTOLOGO NO ENCONTRADO"))
			return
		}
		web.Success(c, 200, odontologo, "ODONTOLOGO OBTENIDO POR ID")
	}
}

// validateEmptys valida que los campos sean correctos
func validateEmptys(odontologo *domain.Odontologo) (bool, error) {
	switch {
	case odontologo.NombreOdontologo == "" || odontologo.ApellidoOdontologo == "" || odontologo.MatriculaOdontologo == "":
		return false, errors.New("LOS CAMPOS NO PUEDEN SER VACIOS")
	}
	return true, nil
}

// PUT --> UpdateOdontologo actualiza un odontologo
func (h *odontologoHandler) UpdateOdontologo() gin.HandlerFunc {
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
		idParam := c.Param("idOdontologo")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("ID INVALIDO"))
			return
		}
		_, err = h.s.GetByID(id)
		if err != nil {
			web.Failure(c, 404, errors.New("ID DE ODONTOLOGO INEXISTENTE"))
			return
		}
		if err != nil {
			web.Failure(c, 409, err)
			return
		}
		var odontologo domain.Odontologo
		err = c.ShouldBindJSON(&odontologo)
		if err != nil {
			web.Failure(c, 400, errors.New("JSON INVALIDO"))
			return
		}
		/* valid, err := validateEmptys(&odontologo)
		if !valid {
			web.Failure(c, 400, err)
			return
		} */
		p, err := h.s.Update(id, odontologo)
		if err != nil {
			web.Failure(c, 409, err)
			return
		}
		web.Success(c, 200, p, "ODONTOLOGO ACTUALIZADO CON EXITO")
	}
}

// PATCH --> UpdateOdontologoForField actualiza un odontologo por alguno de sus campos
func (h *odontologoHandler) UpdateOdontologoForField() gin.HandlerFunc {
	type Request struct {
		NombreOdontologo    string `json:"nombreOdontologo,omitempty"`
		ApellidoOdontologo  string `json:"apellidoOdontologo,omitempty"`
		MatriculaOdontologo string `json:"matriculaOdontologo,omitempty"`
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
		idParam := c.Param("idOdontologo")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("ID INVALIDO"))
			return
		}
		_, err = h.s.GetByID(id)
		if err != nil {
			web.Failure(c, 404, errors.New("ODONTOLOGO INEXISTENTE"))
			return
		}
		if err := c.ShouldBindJSON(&r); err != nil {
			web.Failure(c, 400, errors.New("JSON INVALIDO"))
			return
		}
		update := domain.Odontologo{
			NombreOdontologo:    r.NombreOdontologo,
			ApellidoOdontologo:  r.ApellidoOdontologo,
			MatriculaOdontologo: r.MatriculaOdontologo,
		}
		p, err := h.s.Update(id, update)
		if err != nil {
			web.Failure(c, 409, err)
			return
		}
		web.Success(c, 200, p, "ODONTOLOGO ACTUALIZADO CON EXITO")
	}
}

// DEL --> DeleteOdontologo elimina un odontologo
func (h *odontologoHandler) DeleteOdontologo() gin.HandlerFunc {
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
		idParam := c.Param("idOdontologo")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("ID INVALIDO"))
			return
		}
		_, err = h.s.GetByID(id)
		if err != nil {
			web.Failure(c, 404, errors.New("ID DE ODONTOLOGO INEXISTENTE"))
			return
		}
		err = h.s.Delete(id)
		if err != nil {
			web.Failure(c, 404, err)
			return
		}
		web.Success(c, 200, nil, "ODONTOLOGO ELIMINADO CORRECTAMENTE")
	}
}
