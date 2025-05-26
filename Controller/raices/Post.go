package Raices

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
	stmt, err := db.Prepare("insert into [GWFORMS].[dbo].[FormRaicesInteres] ([nombre],[fechaN],[depto],[mun],[comunidad],[telefono],[disponible]) values (@p1,@p2,@p3,@p4,@p5,@p6,@p7)")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	rows, err := stmt.Query(nombre, fechaN, depto, mun, comunidad, telefono, disponible)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	fmt.Fprint(w, "1")

}
