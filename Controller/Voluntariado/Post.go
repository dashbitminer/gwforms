package Voluntariado

import (
	"context"
	"database/sql"
	"fmt"
	"formulario/Controller"
	"log"
	"net/http"
)

func PostSatisfaccionEvento(w http.ResponseWriter, r *http.Request, Cfg *Controller.Config) {
	pais := r.FormValue("p1")
	asiste := r.FormValue("p45")
	razonnoasiste := r.FormValue("noasiste")
	otroprograma := r.FormValue("otro2")
	edad := r.FormValue("p4")
	sexo := r.FormValue("p5")
	tipovoluntario := r.FormValue("p6")
	opinion := r.FormValue("opinion")
	sugerencias := r.FormValue("sugerencias")
	tienecomentario := r.FormValue("p51")
	comentarios := r.FormValue("comentarios")

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
	// Insertar los valores en la tabla correspondiente

	tx, err := db.Begin()
	if err != nil {
		tx.Rollback()
		http.Error(w, "Error conexion", http.StatusInternalServerError)

	}

	stmt, err = tx.Prepare("INSERT INTO [Voluntariado].[dbo].[F_encuestaSatisfaccionEventoVol]([pais],[asiste],[razonnoasiste],[otroprograma],[edad],[sexo],[tipovoluntario],[opinion],[sugerencias],[tienecomentario],[comentarios]) OUTPUT INSERTED.ID VALUES(@P1,@P2,@P3,@P4,@P5,@P6,@P7,@P8,@P9,@P10,@P11)")
	if err != nil {
		tx.Rollback()
		http.Error(w, "Error al preparar la consulta", http.StatusInternalServerError)

	}

	var lastInsertId int
	err = stmt.QueryRow(pais, asiste, razonnoasiste, otroprograma, edad, sexo, tipovoluntario, opinion, sugerencias, tienecomentario, comentarios).Scan(&lastInsertId)
	if err != nil {
		tx.Rollback()
		http.Error(w, "Error al ejecutar la consulta", http.StatusInternalServerError)
	}

	areaprograma := r.Form["p3"]
	if len(areaprograma) != 0 {
		stmt2, err := db.Prepare("INSERT INTO [Voluntariado].[dbo].[F_multipleVol]([idPregunta],[idFR],[idRespuesta],[idFormulario]) VALUES (@P1, @P2, @P3, @P4)")
		if err != nil {
			tx.Rollback()
			http.Error(w, "Error al ejecutar la consulta area", http.StatusInternalServerError)
		}
		defer stmt.Close()

		for _, valor := range areaprograma {
			_, err = stmt2.Exec(4, lastInsertId, valor, 2)
			if err != nil {
				tx.Rollback()
				http.Error(w, "Error al ejecutar la consulta area", http.StatusInternalServerError)
			}
		}
	}

	evalua := r.Form["p47"]
	if len(evalua) != 0 {
		stmt2, err := db.Prepare("INSERT INTO [Voluntariado].[dbo].[F_multipleVol]([idPregunta],[idFR],[idRespuesta],[idFormulario]) VALUES (@P1, @P2, @P3, @P4)")
		if err != nil {
			tx.Rollback()
			http.Error(w, "Error al ejecutar la consulta evaluacion", http.StatusInternalServerError)
		}
		defer stmt.Close()

		for _, valor := range evalua {
			_, err = stmt2.Exec(9, lastInsertId, valor, 2)
			if err != nil {
				tx.Rollback()
				http.Error(w, "Error al ejecutar la consulta evaluacion", http.StatusInternalServerError)
			}
		}
	}

	aspectos := r.Form["p48"]
	if len(aspectos) != 0 {
		stmt2, err := db.Prepare("INSERT INTO [Voluntariado].[dbo].[F_multipleVol]([idPregunta],[idFR],[idRespuesta],[idFormulario]) VALUES (@P1, @P2, @P3, @P4)")
		if err != nil {
			tx.Rollback()
			http.Error(w, "Error al ejecutar la consulta aspectos", http.StatusInternalServerError)
		}
		defer stmt.Close()

		for _, valor := range aspectos {
			_, err = stmt2.Exec(10, lastInsertId, valor, 2)
			if err != nil {
				tx.Rollback()
				http.Error(w, "Error al ejecutar la consulta aspectos", http.StatusInternalServerError)
			}
		}
	}

	defer stmt.Close()

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
	}

	fmt.Fprint(w, "1")

}
