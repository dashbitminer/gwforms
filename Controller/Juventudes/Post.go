package Juventudes

import (
	"context"
	"database/sql"
	"fmt"
	"formulario/Controller"
	"log"
	"net/http"
)

func PostPre(w http.ResponseWriter, r *http.Request, Cfg *Controller.Config) {
	nombre := r.FormValue("nombreCompleto")
	fechaN := r.FormValue("fechaNacimiento")
	depto := r.FormValue("departamentoResidencia")
	mun := r.FormValue("municipioResidencia")
	comunidad := r.FormValue("coloniaComunidad")
	telefono := r.FormValue("telefonoContacto")
	disponible := r.FormValue("disponibilidad")
	grado := r.FormValue("grado")
	diplomado := r.FormValue("diplomado")
	nombreRep := r.FormValue("nombreRep")
	telefonoRep := r.FormValue("telefonoRep")

	var db *sql.DB
	var stmt *sql.Stmt
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
	if depto == "123" {
		stmt, err = db.Prepare("insert into [Chemonics].[dbo].[FormChemonicsInteres] ([nombre],[fechaN],[depto],[mun],[comunidad],[telefono],[disponible],[grado],[diplomado],[nombreRep],[telefonoRep]) values (@p1,@p2,@p3,@p4,@p5,@p6,@p7,@p8,@p9,@p10,@p11)")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else {
		stmt, err = db.Prepare("insert into [Nido].[dbo].[FormNidoInteres] ([nombre],[fechaN],[depto],[mun],[comunidad],[telefono],[disponible],[grado],[diplomado],[nombreRep],[telefonoRep]) values (@p1,@p2,@p3,@p4,@p5,@p6,@p7,@p8,@p9,@p10,@p11)")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
	rows, err := stmt.Query(nombre, fechaN, depto, mun, comunidad, telefono, disponible, grado, diplomado, nombreRep, telefonoRep)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	fmt.Fprint(w, "1")

}
