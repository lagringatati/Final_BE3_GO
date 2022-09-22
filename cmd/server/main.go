package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	/*"encoding/json"
	"log"*/

	"github.com/gin-gonic/gin"
	"github.com/lagringatati/Final_BE3_GO/cmd/server/handler"
	"github.com/lagringatati/Final_BE3_GO/internal/odontologo"
	"github.com/lagringatati/Final_BE3_GO/pkg/store"
)

/*type paciente struct {
	IdPaciente  int
	Nombre      string
	Apellido    string
	Domicilio   string
	DNI         int
	FechaDeAlta string // date
}

type turno struct {
	IdTurno      int
	IdPaciente   int
	IdOdontologo int
	FechaTurno   string
	HoraTurno    string
	Descripcion  string
}
*/
/*
	func printOdontologo(o odontologo) {
		h := "*"
		log.Print(strings.Repeat(h, 60))
		log.Print(o)
		log.Print(strings.Repeat(h, 60))
	}
*/
func conectDDBB(conexion *sql.DB) {
	Driver := "mysql"
	Usuario := "root"
	Contrasenia := ""
	Nombre := "sistema"

	conexion, err := sql.Open(Driver, Usuario+":"+Contrasenia+"@tcp(127.0.0.1)/"+Nombre)
	if err != nil {
		panic(err.Error())
	}
	//return conexion
}
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

	// inicializo el engine utilizando gin
	engine := gin.Default()

	// declaro un puerto de entrada seguro para la app (en este caso el local)
	engine.SetTrustedProxies([]string{"127.0.0.1"})

	// valido que la conexion se haya realizado correctamente
	engine.GET("/api/v1/ping", func(c *gin.Context) { c.String(200, "pong") })

	odontologos := engine.Group("/api/v1/odontologos")
	{
		odontologos.POST("", odontologoHandler.CreateOdontologo())                       //crear odontologo
		odontologos.GET(":idOdontologo", odontologoHandler.GetOdontologoByID())          //obtener odontologo por id
		odontologos.PUT(":idOdontologo", odontologoHandler.UpdateOdontologo())           //actualizar un odontologo
		odontologos.PATCH(":idOdontologo", odontologoHandler.UpdateOdontologoForField()) //actualiza un odontologo por alguno de sus campos
		odontologos.DELETE(":idOdontologo", odontologoHandler.DeleteOdontologo())        //elimina un odontologo

	}

	// defino el router de chequeo (para que haga la config del router)
	//handlers.NewHealthRouter(engine)

	// inicializo los odontologos
	//odontologoList := store.LoadOdontologos("odontologos.json")

	// defino el handler de odontologos
	//handlers.NewOdontologoRouter(engine, odontologoList)

	// handler.NewOdontologoRouter(engine, *handler.InitOdontologoHandler())
	// hago un mock pero despues tengo q tenerlo en una DDBB
	/*jsonData := `{
		"IdOdontologo"	: 1,
		"Nombre"	: "Sergio",
		"Apellido"	: "Macor",
		"Matricula"	: "MP3142"
	}`*/

	//var o odontologo
	// var p paciente
	// var t turno

	// el unMarchall convierte un arreglo de byte (json) en una estructura que le paso, en este caso "o"
	//if err := json.Unmarshal([]byte(jsonData /*nombre de lo que queremos convertir*/), &o /*o, p, t o sea, a lo que queremos convertir el json*/); err != nil {
	//	log.Fatal("Error fatal en conversion de json")
	//	log.Fatal(err)
	//}

	//printOdontologo(o)

	// defino el handler GET de "/odontologo"
	/*router.GET("/odontologo", func(c *gin.Context) {
		c.JSON(200, gin.H{"odontologo": o})
	})*/

	// defino el puerto en el cual va a correr el servidor
	engine.Run(":8080")

	// para poder detener el proceso por consola utilizar Ctrl + C
	//defer db.Close()

}
