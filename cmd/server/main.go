package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	/*"encoding/json"
	"log"*/

	"github.com/gin-gonic/gin"
	"github.com/lagringatati/Final_BE3_GO/cmd/server/handler"
	"github.com/lagringatati/Final_BE3_GO/internal/odontologo"
	"github.com/lagringatati/Final_BE3_GO/internal/paciente"
	"github.com/lagringatati/Final_BE3_GO/internal/turno"
	"github.com/lagringatati/Final_BE3_GO/pkg/store"
)

/*type turno struct {
	IdTurno      int
	IdPaciente   int
	IdOdontologo int
	FechaTurno   string
	HoraTurno    string
	Descripcion  string
}
*/

func main() {

	db, err := sql.Open("mysql", "root:root@/my_db")
	if err != nil {
		panic(err.Error())
	}
	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	storage := store.NewSqlStore(db)

	repo := odontologo.NewRepository(storage)
	service := odontologo.NewService(repo)
	odontologoHandler := handler.NewOdontologoHandler(service)

	repoPaciente := paciente.NewRepository(storage)
	servicePaciente := paciente.NewService(repoPaciente)
	pacienteHandler := handler.NewPacienteHandler(servicePaciente)

	repoTurno := turno.NewRepository(storage)
	serviceTurno := turno.NewService(repoTurno)
	turnoHandler := handler.NewTurnoHandler(serviceTurno)

	// inicializo el engine utilizando gin
	engine := gin.Default()

	// declaro un puerto de entrada seguro para la app (en este caso el local)
	engine.SetTrustedProxies([]string{"127.0.0.1"})

	// valido que la conexion se haya realizado correctamente
	engine.GET("/api/v1/ping", func(c *gin.Context) { c.String(200, "pong") })

	// determino las uris de odontologo
	odontologos := engine.Group("/api/v1/odontologos")
	{
		odontologos.POST("", odontologoHandler.CreateOdontologo())                       //crear odontologo
		odontologos.GET(":idOdontologo", odontologoHandler.GetOdontologoByID())          //obtener odontologo por id
		odontologos.PUT(":idOdontologo", odontologoHandler.UpdateOdontologo())           //actualizar un odontologo
		odontologos.PATCH(":idOdontologo", odontologoHandler.UpdateOdontologoForField()) //actualizar un odontologo por alguno de sus campos
		odontologos.DELETE(":idOdontologo", odontologoHandler.DeleteOdontologo())        //eliminar un odontologo

	}

	// determino las uris de paciente
	pacientes := engine.Group("/api/v1/pacientes")
	{
		pacientes.POST("", pacienteHandler.CreatePaciente())                     //crear paciente
		pacientes.GET(":idPaciente", pacienteHandler.GetPacienteByID())          //obtener paciente por id
		pacientes.PUT(":idPaciente", pacienteHandler.UpdatePaciente())           //actualizar un paciente
		pacientes.PATCH(":idPaciente", pacienteHandler.UpdatePacienteForField()) //actualizar un paciente por alguno de sus campos
		pacientes.DELETE(":idPaciente", pacienteHandler.DeletePaciente())        //eliminar un paciente

	}

	// determino las uris de turno
	turnos := engine.Group("/api/v1/turnos")
	{
		turnos.POST("", turnoHandler.CreateTurno())                  //crear turno
		turnos.GET(":idTurno", turnoHandler.GetTurnoByID())          //obtener turno por id
		turnos.PUT(":idTurno", turnoHandler.UpdateTurno())           //actualizar un turno
		turnos.PATCH(":idTurno", turnoHandler.UpdateTurnoForField()) //actualizar un turno por alguno de sus campos
		turnos.DELETE(":idTurno", turnoHandler.DeleteTurno())        //eliminar un turno
		// turnos.POST(":dniPaciente", ":matriculaOdontologo", turnoHandler.CreateTurnoSpecial()) //crear turno "especial"
		// turnos.GET(":dniPaciente", turnoHandler.GetTurnoByDni())                               //obtener turno por DNI del paciente

	}

	// defino el puerto en el cual va a correr el servidor
	engine.Run(":8080")

	// para poder detener el proceso por consola utilizar Ctrl + C
	// defer db.Close()

}
