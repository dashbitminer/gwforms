package Salud

import (
	"context"
	"database/sql"
	"fmt"
	"formulario/Controller"
	"html/template"
	"log"
	"net/http"
)

type FormData struct {
	Respuestas2  []Controller.Option
	Parentesco   []Controller.Option
	Sexo         []Controller.Option
	Departamento []Controller.Option
	Respuestas   []Controller.Option
	Ultgrad      []Controller.Option
	Grado        []Controller.Option
	Turno        []Controller.Option
	Seccion      []Controller.Option
	Depsede      []Controller.Option
	Perfil       []Controller.Option
}

func Consentimientomx(w http.ResponseWriter, Cfg *Controller.Config) {

	// Cargar plantilla HTML desde archivo
	tmpl, err := template.ParseFiles("public/HTML/Participantes/Salud/ConsentimientoMx.html")
	if err != nil {
		log.Fatal(err)
	}

	// Renderizar plantilla HTML con opciones de preguntas select
	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func Consentimientogt(w http.ResponseWriter, Cfg *Controller.Config) {

	// Cargar plantilla HTML desde archivo
	tmpl, err := template.ParseFiles("public/HTML/Participantes/Salud/ConsentimientoVirtual.html")
	if err != nil {
		log.Fatal(err)
	}

	// Renderizar plantilla HTML con opciones de preguntas select
	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func Saludhnnew(w http.ResponseWriter, Cfg *Controller.Config) {
	// Build connection string
	var db *sql.DB

	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", Cfg.User, Cfg.Password, Cfg.Server, Cfg.Port, Cfg.Database)
	var err error
	// Create connection pool
	db, err = sql.Open("mysql", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: 1", err.Error())
	}
	defer db.Close()
	ctx := context.Background()
	err = db.PingContext(ctx)
	if err != nil {
		log.Fatal(err.Error())
	}

	// Obtener opciones de la base de datos de MYSQL
	parentesco, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM family_relationships where id not in(15)")
	if err != nil {
		log.Fatal(err)
	}
	sexo, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM genres")
	if err != nil {
		log.Fatal(err)
	}

	departamento, err := Controller.GetOpcionesPregunta(db, "SELECT distinct idd,departamento FROM admin_main_gwdata.geografia where idp=3")
	if err != nil {
		log.Fatal(err)
	}

	respuestas, err := Controller.GetOpcionesPregunta(db, "SELECT id, label FROM F_respuesta")
	if err != nil {
		log.Fatal(err)
	}

	respuestas2, err := Controller.GetOpcionesPregunta(db, "SELECT id,name FROM admin_main_gwdata.institutional_people where active_at is not null and id not in(13) order by name")
	if err != nil {
		log.Fatal(err)
	}

	ultgrad, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM last_grades where active_at is not null")
	if err != nil {
		log.Fatal(err)
	}

	grado, err := Controller.GetOpcionesPregunta(db, "SELECT id,name FROM levels where orden is not null order by orden")
	if err != nil {
		log.Fatal(err)
	}
	seccion, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM section_schools where activo=1 ")
	if err != nil {
		log.Fatal(err)
	}
	turno, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM school_beneficiaries_turns")
	if err != nil {
		log.Fatal(err)
	}

	depsede, err := Controller.GetOpcionesPregunta(db, "SELECT Distinct Iddep,Departamento FROM VSedesSalud where IdPais=3 order by Departamento")
	if err != nil {
		log.Fatal(err)
	}

	perfil, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM admin_main_gwdata.type_beneficiarios where id in (10)")
	if err != nil {
		log.Fatal(err)
	}

	data := FormData{
		Parentesco:   parentesco,
		Sexo:         sexo,
		Departamento: departamento,
		Respuestas:   respuestas,
		Ultgrad:      ultgrad,
		Grado:        grado,
		Turno:        turno,
		Seccion:      seccion,
		Depsede:      depsede,
		Perfil:       perfil,
		Respuestas2:  respuestas2,
	}
	// Cargar plantilla HTML desde archivo
	tmpl, err := template.ParseFiles("public/HTML/Participantes/Salud/saludhnnew.html")
	if err != nil {
		log.Fatal(err)
	}

	// Renderizar plantilla HTML con opciones de preguntas select
	err = tmpl.Execute(w, data)
	if err != nil {
		log.Fatal(err)
	}
}
