package CdN

import (
	"database/sql"
	"fmt"
	"formulario/Controller"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func ObtenerdepsedeHandlersalud(w http.ResponseWriter, r *http.Request, Cfg *Controller.Config) {
	paissede := r.FormValue("paissede")
	var opcionesdepsede string
	var db *sql.DB

	// Build connection string
	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", Cfg.User, Cfg.Password, Cfg.Server, Cfg.Port, Cfg.Database)
	var err error
	// Create connection pool
	db, err = sql.Open("mysql", connString)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	rows, err := db.Query("SELECT Distinct Iddep,Departamento FROM VSedesSalud where IdPais=? order by Departamento", paissede)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	opcionesdepsede += "<option value=''></option>"
	for rows.Next() {
		var id int
		var name string
		err := rows.Scan(&id, &name)
		if err != nil {
			log.Fatal(err)
		}
		opcionesdepsede += fmt.Sprintf("<option value='%d'>%s</option>", id, name)
	}
	fmt.Fprint(w, opcionesdepsede)
}
