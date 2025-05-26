package Sanamente

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
	Parentesco   []Controller.Option
	Sexo         []Controller.Option
	Departamento []Controller.Option
	Respuestas   []Controller.Option
	Ultgrad      []Controller.Option
	Grado        []Controller.Option
	Perfil       []Controller.Option
	IdPais       int
	Turno        []Controller.Option
	Seccion      []Controller.Option
	Depsede      []Controller.Option
	Respuestas2  []Controller.Option
}

func SanamenteInscripcionPNCsv(w http.ResponseWriter, r *http.Request, Cfg *Controller.Config) {
	// Build connection string
	/* 	var db *sql.DB

	   	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;", Cfg.Server2, Cfg.User2, Cfg.Password2, Cfg.Port2, Cfg.Database4)
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

	   	type FormData struct {
	   		Respuestas1 []Controller.Option
	   		Respuestas2 []Controller.Option
	   		Respuestas3 []Controller.Option
	   		Respuestas4 []Controller.Option
	   		Respuestas5 []Controller.Option
	   		Respuestas6 []Controller.Option
	   	}
	   	// Obtener opciones de la base de datos de MYSQL

	   	respuestas1, err := Controller.GetOpcionesPregunta(db, "SELECT [id],[name] FROM [admin_main_gwdata].[genres];")
	   	if err != nil {
	   		http.Error(w, err.Error(), http.StatusInternalServerError)
	   		return
	   	}

	   	respuestas2, err := Controller.GetOpcionesPregunta(db, "SELECT DISTINCT [idd],[departamento] FROM [geografia] where idp=7;")
	   	if err != nil {
	   		http.Error(w, err.Error(), http.StatusInternalServerError)
	   		return
	   	}

	   	respuestas3, err := Controller.GetOpcionesPregunta(db, "SELECT [id],[name] FROM [admin_main_gwdata].[last_grades] where active_at is not null;")
	   	if err != nil {
	   		http.Error(w, err.Error(), http.StatusInternalServerError)
	   		return
	   	}

	   	respuestas4, err := Controller.GetOpcionesPregunta(db, "SELECT [id],[name] FROM [admin_main_gwdata].[levels] where orden is not null order by orden;")
	   	if err != nil {
	   		http.Error(w, err.Error(), http.StatusInternalServerError)
	   		return
	   	}

	   	respuestas5, err := Controller.GetOpcionesPregunta(db, "SELECT id,name FROM [admin_main_gwdata].[sanamente_subtipos] where id_institucional=1 and id not in (12) order by name")
	   	if err != nil {
	   		http.Error(w, err.Error(), http.StatusInternalServerError)
	   		return
	   	}

	   	respuestas6, err := Controller.GetOpcionesPregunta(db, "SELECT id,name FROM [admin_main_gwdata].[sanamente_subtipos] where id_institucional=1 and id not in (12) order by name")
	   	if err != nil {
	   		http.Error(w, err.Error(), http.StatusInternalServerError)
	   		return
	   	}

	   	data := FormData{
	   		Respuestas1: respuestas1,
	   		Respuestas2: respuestas2,
	   		Respuestas3: respuestas3,
	   		Respuestas4: respuestas4,
	   		Respuestas5: respuestas5,
	   		Respuestas6: respuestas6,
	   	}
	   	// Cargar plantilla HTML desde archivo
	   	tmpl, err := template.ParseFiles("public/HTML/sanamente/sv/pncinscripcionsv.html")
	   	if err != nil {
	   		http.Error(w, err.Error(), http.StatusInternalServerError)
	   		return
	   	}

	   	// Renderizar plantilla HTML con opciones de preguntas select
	   	err = tmpl.Execute(w, data)
	   	if err != nil {
	   		http.Error(w, err.Error(), http.StatusInternalServerError)
	   		return
	   	} */

	// Redirigir a la página HTML por encuesta cerrada
	http.ServeFile(w, r, "public/HTML/sanamente/RespuestaSanamenteRecursoCerrado.html")
}

func ModuloasvHandlerPNC(w http.ResponseWriter, Cfg *Controller.Config) {
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
	parentesco, err := Controller.GetOpcionesPregunta(db, "SELECT id,Respuesta FROM admin_main_gwdata.Sanamente_quiz_P where modulo=1 and idPregunta=1 and Version=1;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	sexo, err := Controller.GetOpcionesPregunta(db, "SELECT id,Respuesta FROM admin_main_gwdata.Sanamente_quiz_P where modulo=1 and idPregunta=2 and Version=1;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	departamento, err := Controller.GetOpcionesPregunta(db, "SELECT Distinct Iddep,Departamento FROM VSedessanamente where categorizacion=4 and IdPais=7 order by Departamento")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respuestas, err := Controller.GetOpcionesPregunta(db, "SELECT id,Respuesta FROM admin_main_gwdata.Sanamente_quiz_P where modulo=1 and idPregunta=3 and Version=1;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ultgrad, err := Controller.GetOpcionesPregunta(db, "SELECT id,Respuesta FROM admin_main_gwdata.Sanamente_quiz_P where modulo=1 and idPregunta=4 and Version=1;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	grado, err := Controller.GetOpcionesPregunta(db, "SELECT id,Respuesta FROM admin_main_gwdata.Sanamente_quiz_P where modulo=1 and idPregunta=5 and Version=1;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	perfil, err := Controller.GetOpcionesPregunta(db, "SELECT id,Formador FROM F_formadorSanamente where idpais=7  and activo=1 order by Formador")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := FormData{
		Parentesco:   parentesco,
		Sexo:         sexo,
		Departamento: departamento,
		Respuestas:   respuestas,
		Ultgrad:      ultgrad,
		Grado:        grado,
		Perfil:       perfil,
	}
	// Cargar plantilla HTML desde archivo
	tmpl, err := template.ParseFiles("public/HTML/sanamente/sv/pncModuloA.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Renderizar plantilla HTML con opciones de preguntas select
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func ModulobsvHandlerPNC(w http.ResponseWriter, Cfg *Controller.Config) {
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
	parentesco, err := Controller.GetOpcionesPregunta(db, "SELECT id,Respuesta FROM admin_main_gwdata.Sanamente_quiz_P where modulo=2 and idPregunta=1 and Version=1;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	sexo, err := Controller.GetOpcionesPregunta(db, "SELECT id,Respuesta FROM admin_main_gwdata.Sanamente_quiz_P where modulo=2 and idPregunta=2 and Version=1;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	departamento, err := Controller.GetOpcionesPregunta(db, "SELECT Distinct Iddep,Departamento FROM VSedessanamente where categorizacion=4 and IdPais=7 order by Departamento")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respuestas, err := Controller.GetOpcionesPregunta(db, "SELECT id,Respuesta FROM admin_main_gwdata.Sanamente_quiz_P where modulo=2 and idPregunta=3 and Version=1;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ultgrad, err := Controller.GetOpcionesPregunta(db, "SELECT id,Respuesta FROM admin_main_gwdata.Sanamente_quiz_P where modulo=2 and idPregunta=4 and Version=1;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	grado, err := Controller.GetOpcionesPregunta(db, "SELECT id,Respuesta FROM admin_main_gwdata.Sanamente_quiz_P where modulo=2 and idPregunta=5 and Version=1;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	perfil, err := Controller.GetOpcionesPregunta(db, "SELECT id,Formador FROM F_formadorSanamente where idpais=7  and activo=1  order by Formador")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := FormData{
		Parentesco:   parentesco,
		Sexo:         sexo,
		Departamento: departamento,
		Respuestas:   respuestas,
		Ultgrad:      ultgrad,
		Grado:        grado,
		Perfil:       perfil,
	}
	// Cargar plantilla HTML desde archivo
	tmpl, err := template.ParseFiles("public/HTML/sanamente/sv/pncModuloB.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Renderizar plantilla HTML con opciones de preguntas select
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func ModulocsvHandlerPNC(w http.ResponseWriter, Cfg *Controller.Config) {
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
	parentesco, err := Controller.GetOpcionesPregunta(db, "SELECT id,Respuesta FROM admin_main_gwdata.Sanamente_quiz_P where modulo=3 and idPregunta=1 and Version=1;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	sexo, err := Controller.GetOpcionesPregunta(db, "SELECT id,Respuesta FROM admin_main_gwdata.Sanamente_quiz_P where modulo=3 and idPregunta=2 and Version=1;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	departamento, err := Controller.GetOpcionesPregunta(db, "SELECT Distinct Iddep,Departamento FROM VSedessanamente where categorizacion=4 and IdPais=7 order by Departamento")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respuestas, err := Controller.GetOpcionesPregunta(db, "SELECT id,Respuesta FROM admin_main_gwdata.Sanamente_quiz_P where modulo=3 and idPregunta=3 and Version=1;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ultgrad, err := Controller.GetOpcionesPregunta(db, "SELECT id,Respuesta FROM admin_main_gwdata.Sanamente_quiz_P where modulo=3 and idPregunta=4 and Version=1;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	grado, err := Controller.GetOpcionesPregunta(db, "SELECT id,Respuesta FROM admin_main_gwdata.Sanamente_quiz_P where modulo=3 and idPregunta=5 and Version=1;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	perfil, err := Controller.GetOpcionesPregunta(db, "SELECT id,Formador FROM F_formadorSanamente where idpais=7  and activo=1  order by Formador")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := FormData{
		Parentesco:   parentesco,
		Sexo:         sexo,
		Departamento: departamento,
		Respuestas:   respuestas,
		Ultgrad:      ultgrad,
		Grado:        grado,
		Perfil:       perfil,
	}
	// Cargar plantilla HTML desde archivo
	tmpl, err := template.ParseFiles("public/HTML/sanamente/sv/pncModuloC.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Renderizar plantilla HTML con opciones de preguntas select
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func EncuestaRedApoyos(w http.ResponseWriter, r *http.Request, Cfg *Controller.Config) {

	idpais := 0

	pais := r.URL.Query().Get("p")

	switch pais {
	case "sv":
		idpais = 7
	case "gt":
		idpais = 4
	case "hn":
		idpais = 3
	case "mx":
		idpais = 17
	case "col":
		idpais = 16
	case "pn":
		idpais = 1
	}

	data := FormData{
		IdPais: idpais,
	}

	// Cargar plantilla HTML desde archivo
	tmpl, err := template.ParseFiles("public/HTML/sanamente/EncuestaRedApoyos.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Renderizar plantilla HTML con opciones de preguntas select
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func Sanamentehnnew(w http.ResponseWriter, Cfg *Controller.Config) {
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
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	sexo, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM genres")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	departamento, err := Controller.GetOpcionesPregunta(db, "SELECT distinct idd,departamento FROM admin_main_gwdata.geografia where idp=3")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respuestas, err := Controller.GetOpcionesPregunta(db, "SELECT id, label FROM F_respuesta")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respuestas2, err := Controller.GetOpcionesPregunta(db, "SELECT id,name FROM admin_main_gwdata.institutional_people where active_at is not null and id not in(13) order by name")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ultgrad, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM last_grades where active_at is not null")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	grado, err := Controller.GetOpcionesPregunta(db, "SELECT id,name FROM levels where orden is not null order by orden")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	seccion, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM section_schools where activo=1 ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	turno, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM school_beneficiaries_turns")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	depsede, err := Controller.GetOpcionesPregunta(db, "SELECT Distinct Iddep,Departamento FROM VSedessanamente where IdPais=3 order by Departamento")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	perfil, err := Controller.GetOpcionesPregunta(db, "SELECT id, name FROM admin_main_gwdata.type_beneficiarios where id in (10)")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
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
	tmpl, err := template.ParseFiles("public/HTML/sanamente/hn/inscripcionhn copy.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Renderizar plantilla HTML con opciones de preguntas select
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func EncuestaSatisfaccionInter(w http.ResponseWriter, r *http.Request, Cfg *Controller.Config) {

	idpais := 0

	pais := r.URL.Query().Get("p")

	switch pais {
	case "sv":
		idpais = 7
	case "gt":
		idpais = 4
	case "hn":
		idpais = 3
	case "mx":
		idpais = 17
	case "col":
		idpais = 16
	case "pn":
		idpais = 1
	}

	data := FormData{
		IdPais: idpais,
	}

	// Cargar plantilla HTML desde archivo
	tmpl, err := template.ParseFiles("public/HTML/sanamente/EncuestaSatisfaccionInter.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Renderizar plantilla HTML con opciones de preguntas select
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
