package CdN

import (
	"context"
	"database/sql"
	"fmt"
	"formulario/Controller"
	"html/template"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/microsoft/go-mssqldb"
)

type FormData struct {
	Respuestas1 []Controller.Option
	Respuestas2 []Controller.Option
	Respuestas3 []Controller.Option
	Respuestas4 []Controller.Option
}

func Taller1CdNHandler(w http.ResponseWriter, Cfg *Controller.Config) {
	// Build connection string
	var db *sql.DB

	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;", Cfg.Server2, Cfg.User2, Cfg.Password2, Cfg.Port2, Cfg.Database3)
	var err error
	// Create connection pool
	db, err = sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: ", err.Error())
	}
	ctx := context.Background()
	err = db.PingContext(ctx)
	if err != nil {
		log.Fatal(err.Error())
	}

	defer db.Close()

	// Obtener opciones de la base de datos de MYSQL

	respuestas1, err := Controller.GetOpcionesPregunta(db, "SELECT [id], [respuesta] FROM [FormCDNQuizCatRes] where idF=10 and modulo=1 and idPregunta=1 and versionf=1;")
	if err != nil {
		log.Fatal(err)
	}

	respuestas2, err := Controller.GetOpcionesPregunta(db, "SELECT [id], [respuesta] FROM [FormCDNQuizCatRes] where idF=10 and modulo=1 and idPregunta=2 and versionf=1;")
	if err != nil {
		log.Fatal(err)
	}

	respuestas3, err := Controller.GetOpcionesPregunta(db, "SELECT [id], [respuesta] FROM [FormCDNQuizCatRes] where idF=10 and modulo=1 and idPregunta=3 and versionf=1;")
	if err != nil {
		log.Fatal(err)
	}

	respuestas4, err := Controller.GetOpcionesPregunta(db, "SELECT [id], [respuesta] FROM [FormCDNQuizCatRes] where idF=10 and modulo=1 and idPregunta=4 and versionf=1;")
	if err != nil {
		log.Fatal(err)
	}

	data := FormData{
		Respuestas1: respuestas1,
		Respuestas2: respuestas2,
		Respuestas3: respuestas3,
		Respuestas4: respuestas4,
	}
	// Cargar plantilla HTML desde archivo
	tmpl, err := template.ParseFiles("public/HTML/Participantes/Salud/CdN/Taller1.html")
	if err != nil {
		log.Fatal(err)
	}

	// Renderizar plantilla HTML con opciones de preguntas select
	err = tmpl.Execute(w, data)
	if err != nil {
		log.Fatal(err)
	}
}

func Taller2CdNHandler(w http.ResponseWriter, Cfg *Controller.Config) {
	// Build connection string
	var db *sql.DB

	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;", Cfg.Server2, Cfg.User2, Cfg.Password2, Cfg.Port2, Cfg.Database3)
	var err error
	// Create connection pool
	db, err = sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: ", err.Error())
	}
	ctx := context.Background()
	err = db.PingContext(ctx)
	if err != nil {
		log.Fatal(err.Error())
	}

	defer db.Close()

	// Obtener opciones de la base de datos de MYSQL

	respuestas1, err := Controller.GetOpcionesPregunta(db, "SELECT [id], [respuesta] FROM [FormCDNQuizCatRes] where idF=10 and modulo=2 and idPregunta=1 and versionf=1;")
	if err != nil {
		log.Fatal(err)
	}

	respuestas2, err := Controller.GetOpcionesPregunta(db, "SELECT [id], [respuesta] FROM [FormCDNQuizCatRes] where idF=10 and modulo=2 and idPregunta=2 and versionf=1;")
	if err != nil {
		log.Fatal(err)
	}

	respuestas3, err := Controller.GetOpcionesPregunta(db, "SELECT [id], [respuesta] FROM [FormCDNQuizCatRes] where idF=10 and modulo=2 and idPregunta=3 and versionf=1;")
	if err != nil {
		log.Fatal(err)
	}

	data := FormData{
		Respuestas1: respuestas1,
		Respuestas2: respuestas2,
		Respuestas3: respuestas3,
	}
	// Cargar plantilla HTML desde archivo
	tmpl, err := template.ParseFiles("public/HTML/Participantes/Salud/CdN/Taller2.html")
	if err != nil {
		log.Fatal(err)
	}

	// Renderizar plantilla HTML con opciones de preguntas select
	err = tmpl.Execute(w, data)
	if err != nil {
		log.Fatal(err)
	}
}

func Taller3CdNHandler(w http.ResponseWriter, Cfg *Controller.Config) {
	// Build connection string
	var db *sql.DB

	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;", Cfg.Server2, Cfg.User2, Cfg.Password2, Cfg.Port2, Cfg.Database3)
	var err error
	// Create connection pool
	db, err = sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: ", err.Error())
	}
	ctx := context.Background()
	err = db.PingContext(ctx)
	if err != nil {
		log.Fatal(err.Error())
	}

	defer db.Close()

	// Obtener opciones de la base de datos de MYSQL

	respuestas1, err := Controller.GetOpcionesPregunta(db, "SELECT [id], [respuesta] FROM [FormCDNQuizCatRes] where idF=10 and modulo=3 and idPregunta=1 and versionf=1;")
	if err != nil {
		log.Fatal(err)
	}

	respuestas2, err := Controller.GetOpcionesPregunta(db, "SELECT [id], [respuesta] FROM [FormCDNQuizCatRes] where idF=10 and modulo=3 and idPregunta=2 and versionf=1;")
	if err != nil {
		log.Fatal(err)
	}

	respuestas3, err := Controller.GetOpcionesPregunta(db, "SELECT [id], [respuesta] FROM [FormCDNQuizCatRes] where idF=10 and modulo=3 and idPregunta=3 and versionf=1;")
	if err != nil {
		log.Fatal(err)
	}

	respuestas4, err := Controller.GetOpcionesPregunta(db, "SELECT [id], [respuesta] FROM [FormCDNQuizCatRes] where idF=10 and modulo=3 and idPregunta=4 and versionf=1;")
	if err != nil {
		log.Fatal(err)
	}

	data := FormData{
		Respuestas1: respuestas1,
		Respuestas2: respuestas2,
		Respuestas3: respuestas3,
		Respuestas4: respuestas4,
	}
	// Cargar plantilla HTML desde archivo
	tmpl, err := template.ParseFiles("public/HTML/Participantes/Salud/CdN/Taller3.html")
	if err != nil {
		log.Fatal(err)
	}

	// Renderizar plantilla HTML con opciones de preguntas select
	err = tmpl.Execute(w, data)
	if err != nil {
		log.Fatal(err)
	}
}

func Genero(w http.ResponseWriter) {

	// Cargar plantilla HTML desde archivo
	tmpl, err := template.ParseFiles("public/HTML/Participantes/Salud/CdN/GeneroRegional.html")
	if err != nil {
		log.Fatal(err)
	}

	// Renderizar plantilla HTML con opciones de preguntas select
	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func GeneroSV(w http.ResponseWriter) {

	// Cargar plantilla HTML desde archivo
	tmpl, err := template.ParseFiles("public/HTML/Participantes/Salud/CdN/GeneroSV.html")
	if err != nil {
		log.Fatal(err)
	}

	// Renderizar plantilla HTML con opciones de preguntas select
	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Fatal(err)
	}
}
