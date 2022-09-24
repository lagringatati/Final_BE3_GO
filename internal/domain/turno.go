package domain

type Turno struct {
	IdTurno          int    `json:"idTurno"`
	DescripcionTurno string `json:"descripcionTurno" binding:"required"`
	FechaTurno       string `json:"fechaTurno" binding:"required"`
	IdOdontologo     int    `json:"idOdontologo" binding:"required"`
	IdPaciente       int    `json:"idPaciente" binding:"required"`
}
