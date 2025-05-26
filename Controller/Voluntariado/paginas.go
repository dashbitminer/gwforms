package Voluntariado

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

type Respuestas struct {
	Value      string
	Label      string
	IdPregunta string
}

type FormData struct {
	Parentesco   []Controller.Option
	Sexo         []Controller.Option
	Departamento []Controller.Option
	Respuestas   []Controller.Option
	Ultgrad      []Controller.Option
	Grado        []Controller.Option
	Perfil       []Controller.Option
	IdPais       int
}

type FormDataNew struct {
	Preguntas []Respuestas
	Solucion  []Respuestas
}

func getOpcionesPreguntaNew(db *sql.DB, Query string) ([]Respuestas, error) {
	rows, err := db.Query(Query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var opciones []Respuestas
	for rows.Next() {
		var value, label, idPregunta string
		err := rows.Scan(&value, &label, &idPregunta)
		if err != nil {
			return nil, err
		}
		opciones = append(opciones, Respuestas{Value: value, Label: label, IdPregunta: idPregunta})
	}

	return opciones, nil
}

func SatisfaccionEvento(w http.ResponseWriter, Cfg *Controller.Config) {

	// Build connection string
	var db2 *sql.DB
	connString2 := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;", Cfg.Server2, Cfg.User2, Cfg.Password2, Cfg.Port2, Cfg.Database3)

	db2, err := sql.Open("sqlserver", connString2)
	if err != nil {
		log.Fatal("Error creating connection pool: ", err.Error())
	}
	ctx2 := context.Background()
	err = db2.PingContext(ctx2)
	if err != nil {
		log.Fatal(err.Error())
	}

	defer db2.Close()

	// Obtener opciones de la base de datos de SQL
	preguntas, err := getOpcionesPreguntaNew(db2, "SELECT distinct [id], [Pregunta], [idformulario] FROM [Voluntariado].[dbo].[Preguntas_formularios] where ([idformulario]=2 and [version]=1) OR (id IN(1,3,4,5,6)) and [deleted] is null")
	if err != nil {
		log.Fatal(err)
	}

	solucion, err := getOpcionesPreguntaNew(db2, "SELECT distinct [id], [Respuesta], [idPregunta] FROM [Voluntariado].[dbo].[respuesta_vol] where ([idformulario]=2 and [version]=1) OR (idPregunta IN(1,3,19,4,5,12)) and [deleted] is null")
	if err != nil {
		log.Fatal(err)
	}

	data := FormDataNew{
		Solucion:  solucion,
		Preguntas: preguntas,
	}
	// Cargar plantilla HTML desde archivo
	tmpl, err := template.ParseFiles("public/HTML/Voluntarios/SatisfaccionEventoVoluntariado.html")
	if err != nil {
		log.Fatal(err)
	}

	// Renderizar plantilla HTML con opciones de preguntas select
	err = tmpl.Execute(w, data)
	if err != nil {
		log.Fatal(err)
	}

}

func RespuestaEncuestaVol(w http.ResponseWriter) {

	// Cargar plantilla HTML desde archivo
	tmpl, err := template.ParseFiles("public/HTML/RespuestaEncuestaEvento.html")
	if err != nil {
		log.Fatal(err)
	}

	// Renderizar plantilla HTML con opciones de preguntas select
	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Fatal(err)
	}
}
