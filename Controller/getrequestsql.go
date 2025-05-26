package Controller

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
)

func Obtenermunicipiosql(w http.ResponseWriter, r *http.Request, Cfg *Config) {
	departamento := r.FormValue("departamento")
	var opcionesmunsede string
	// Build connection string
	var db *sql.DB
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
	stmt, err := db.Prepare("SELECT [idm],[municipio] FROM [DATAGW].[dbo].[geografia] where [idd]=@p1 order by municipio")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	rows, err := stmt.Query(departamento)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()
	opcionesmunsede += "<option value=''></option>"
	for rows.Next() {
		var id int
		var name string
		err := rows.Scan(&id, &name)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		opcionesmunsede += fmt.Sprintf("<option value='%d'>%s</option>", id, name)
	}
	fmt.Fprint(w, opcionesmunsede)
}
