package CdN

import (
	"context"
	"database/sql"
	"fmt"
	"formulario/Controller"
	"html/template"
	"log"
	"net/http"
)

func PostTallerCdN(w http.ResponseWriter, r *http.Request, Cfg *Controller.Config) {
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

	// Obtener los valores de los campos del formulario

	idF := 10
	sexo := r.FormValue("sexo")
	pais := r.FormValue("pais")
	sede := r.FormValue("sede2")
	p1 := r.FormValue("p1")
	p2 := r.FormValue("p2")
	p3 := r.FormValue("p3")
	p4 := r.FormValue("p4")
	p5 := r.FormValue("p5")
	modulo := r.FormValue("modulo")

	type Datos struct {
		Id int
	}

	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
		tx.Rollback()
		http.Error(w, "Error conexion", http.StatusInternalServerError)

	}

	stmt, err := tx.Prepare("INSERT INTO FormCDNAQuizRegistro (idF,sexo, pais, sede, modulo, p1,p2,p3,p4,p5) OUTPUT INSERTED.ID VALUES(@P1,@P2,@P3,@P4,@P5,@P6,@P7,@P8,@P9,@P10)")
	if err != nil {
		log.Fatal(err)
		tx.Rollback()
		http.Error(w, "Error al preparar la consulta", http.StatusInternalServerError)

	}

	var lastInsertId int
	err = stmt.QueryRow(idF, sexo, pais, sede, modulo, p1, p2, p3, p4, p5).Scan(&lastInsertId)
	if err != nil {
		log.Fatal(err)
		tx.Rollback()
		http.Error(w, "Error al ejecutar la consulta", http.StatusInternalServerError)
	}

	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}

	datos := Datos{
		Id: lastInsertId,
	}

	t, err := template.ParseFiles("Public/Html/Respuesta.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Renderizar la plantilla con los datos
	err = t.Execute(w, datos)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func PostCuestionarioRegionalGenero(w http.ResponseWriter, r *http.Request, Cfg *Controller.Config) {
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

	// Obtener los valores de los campos del formulario

	idF := 14
	edad := r.FormValue("edad")
	pais := r.FormValue("pais")
	sede := r.FormValue("sede2")

	violenciaidentificada := r.FormValue("violencia")
	respcaso := r.FormValue("caso")
	respsituacion1 := r.FormValue("sr1")
	respsituacion2 := r.FormValue("sr2")
	respsituacion3 := r.FormValue("sr3")
	respsituacion4 := r.FormValue("sr4")
	respsituacion5 := r.FormValue("sr5")
	respcaso2 := r.FormValue("caso2")
	motivorespcaso2 := r.FormValue("motivocaso2")

	opinionfrase1 := r.FormValue("psb1")
	opinionfrase2 := r.FormValue("psb2")
	opinionfrase3 := r.FormValue("psb3")
	opinionfrase4 := r.FormValue("psb4")
	opinionfrase5 := r.FormValue("psb5")

	respcaso3 := r.FormValue("caso3")
	motivorespcaso3 := r.FormValue("motivocaso3")
	otroconsejo := r.FormValue("otroconsejo")

	confianzahabla := r.FormValue("seleccionp8")
	confianzadecide := r.FormValue("seleccionp8-1")
	familiarentaller := r.FormValue("talleres")

	aporteclub := r.FormValue("seleccionp9")
	metas := r.FormValue("metas")
	percepcion := r.FormValue("percepcion")
	sugerencia := r.FormValue("mejoras")
	practica := r.FormValue("practica")

	type Datos struct {
		Id int
	}

	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
		tx.Rollback()
		http.Error(w, "Error conexion", http.StatusInternalServerError)

	}

	stmt, err := tx.Prepare("INSERT INTO FormGeneroRegional ([idF],[pais],[edad],[sede],[violenciaidentificada],[respcaso],[respsituacion1],[respsituacion2],[respsituacion3],[respsituacion4],[respsituacion5],[respcaso2],[motivorespcaso2],[opinionfrase1],[opinionfrase2],[opinionfrase3],[opinionfrase4],[opinionfrase5],[respcaso3],[motivorespcaso3],[otroconsejo],[confianzahabla],[confianzadecide],[familiarentaller],[aporteclub],[metas],[percepcion],[sugerencia],[practica]) OUTPUT INSERTED.ID VALUES(@P1,@P2,@P3,@P4,@P5,@P6,@P7,@P8,@P9,@P10,@P11,@P12,@P13,@P14,@P15,@P16,@P17,@P18,@P19,@P20,@P21,@P22,@P23,@P24,@P25,@P26,@P27,@P28,@p29)")
	if err != nil {
		log.Fatal(err)
		tx.Rollback()
		http.Error(w, "Error al preparar la consulta", http.StatusInternalServerError)

	}

	var lastInsertId int
	err = stmt.QueryRow(idF, pais, edad, sede, violenciaidentificada, respcaso, respsituacion1, respsituacion2, respsituacion3, respsituacion4, respsituacion5, respcaso2, motivorespcaso2, opinionfrase1, opinionfrase2, opinionfrase3, opinionfrase4, opinionfrase5, respcaso3, motivorespcaso3, otroconsejo, confianzahabla, confianzadecide, familiarentaller, aporteclub, metas, percepcion, sugerencia, practica).Scan(&lastInsertId)
	if err != nil {
		log.Fatal(err)
		tx.Rollback()
		http.Error(w, "Error al ejecutar la consulta", http.StatusInternalServerError)
	}

	situacionesviolentasmultiple := r.Form["situacion"]
	if len(situacionesviolentasmultiple) == 0 {
	} else {
		stmt, err := tx.Prepare("INSERT INTO MultiplesSalud (idparticipante,idrespuesta,idpregunta,idF) VALUES (@P1,@P2,1,14)")
		if err != nil {
			log.Fatal(err)
			tx.Rollback()
			http.Error(w, "Error al preparar la consulta ocupacion", http.StatusInternalServerError)
			return
		}
		defer stmt.Close()

		for _, valor := range situacionesviolentasmultiple {
			if valor != "" {
				_, err = stmt.Exec(lastInsertId, valor)
				if err != nil {
					log.Fatal(err)
					tx.Rollback()
					http.Error(w, "Error al ejecutar la consulta situaciones", http.StatusInternalServerError)
					return
				}
			}
		}
	}

	tipoviolenciasituacion1multiple := r.Form["sr1-opc"]
	if len(tipoviolenciasituacion1multiple) == 0 {
	} else {
		stmt, err := tx.Prepare("INSERT INTO MultiplesSalud (idparticipante,idrespuesta,idpregunta,idF) VALUES (@P1,@P2,2,14)")
		if err != nil {
			log.Fatal(err)
			tx.Rollback()
			http.Error(w, "Error al preparar la consulta ocupacion", http.StatusInternalServerError)
			return
		}
		defer stmt.Close()

		for _, valor := range tipoviolenciasituacion1multiple {
			if valor != "" {
				_, err = stmt.Exec(lastInsertId, valor)
				if err != nil {
					log.Fatal(err)
					tx.Rollback()
					http.Error(w, "Error al ejecutar la consulta situaciones", http.StatusInternalServerError)
					return
				}
			}
		}
	}

	tipoviolenciasituacion2multiple := r.Form["sr2-opc"]
	if len(tipoviolenciasituacion2multiple) == 0 {
	} else {
		stmt, err := tx.Prepare("INSERT INTO MultiplesSalud (idparticipante,idrespuesta,idpregunta,idF) VALUES (@P1,@P2,3,14)")
		if err != nil {
			log.Fatal(err)
			tx.Rollback()
			http.Error(w, "Error al preparar la consulta ocupacion", http.StatusInternalServerError)
			return
		}
		defer stmt.Close()

		for _, valor := range tipoviolenciasituacion2multiple {
			if valor != "" {
				_, err = stmt.Exec(lastInsertId, valor)
				if err != nil {
					log.Fatal(err)
					tx.Rollback()
					http.Error(w, "Error al ejecutar la consulta situaciones", http.StatusInternalServerError)
					return
				}
			}
		}
	}

	tipoviolenciasituacion3multiple := r.Form["sr3-opc"]
	if len(tipoviolenciasituacion3multiple) == 0 {
	} else {
		stmt, err := tx.Prepare("INSERT INTO MultiplesSalud (idparticipante,idrespuesta,idpregunta,idF) VALUES (@P1,@P2,4,14)")
		if err != nil {
			log.Fatal(err)
			tx.Rollback()
			http.Error(w, "Error al preparar la consulta ocupacion", http.StatusInternalServerError)
			return
		}
		defer stmt.Close()

		for _, valor := range tipoviolenciasituacion3multiple {
			if valor != "" {
				_, err = stmt.Exec(lastInsertId, valor)
				if err != nil {
					log.Fatal(err)
					tx.Rollback()
					http.Error(w, "Error al ejecutar la consulta situaciones", http.StatusInternalServerError)
					return
				}
			}
		}
	}

	tipoviolenciasituacion4multiple := r.Form["sr4-opc"]
	if len(tipoviolenciasituacion4multiple) == 0 {
	} else {
		stmt, err := tx.Prepare("INSERT INTO MultiplesSalud (idparticipante,idrespuesta,idpregunta,idF) VALUES (@P1,@P2,5,14)")
		if err != nil {
			log.Fatal(err)
			tx.Rollback()
			http.Error(w, "Error al preparar la consulta ocupacion", http.StatusInternalServerError)
			return
		}
		defer stmt.Close()

		for _, valor := range tipoviolenciasituacion4multiple {
			if valor != "" {
				_, err = stmt.Exec(lastInsertId, valor)
				if err != nil {
					log.Fatal(err)
					tx.Rollback()
					http.Error(w, "Error al ejecutar la consulta situaciones", http.StatusInternalServerError)
					return
				}
			}
		}
	}

	tipoviolenciasituacion5multiple := r.Form["sr5-opc"]
	if len(tipoviolenciasituacion5multiple) == 0 {
	} else {
		stmt, err := tx.Prepare("INSERT INTO MultiplesSalud (idparticipante,idrespuesta,idpregunta,idF) VALUES (@P1,@P2,6,14)")
		if err != nil {
			log.Fatal(err)
			tx.Rollback()
			http.Error(w, "Error al preparar la consulta ocupacion", http.StatusInternalServerError)
			return
		}
		defer stmt.Close()

		for _, valor := range tipoviolenciasituacion5multiple {
			if valor != "" {
				_, err = stmt.Exec(lastInsertId, valor)
				if err != nil {
					log.Fatal(err)
					tx.Rollback()
					http.Error(w, "Error al ejecutar la consulta situaciones", http.StatusInternalServerError)
					return
				}
			}
		}
	}

	pedirconsejomultiple := r.Form["consejo"]
	if len(pedirconsejomultiple) == 0 {
	} else {
		stmt, err := tx.Prepare("INSERT INTO MultiplesSalud (idparticipante,idrespuesta,idpregunta,idF) VALUES (@P1,@P2,7,14)")
		if err != nil {
			log.Fatal(err)
			tx.Rollback()
			http.Error(w, "Error al preparar la consulta ocupacion", http.StatusInternalServerError)
			return
		}
		defer stmt.Close()

		for _, valor := range pedirconsejomultiple {
			if valor != "" {
				_, err = stmt.Exec(lastInsertId, valor)
				if err != nil {
					log.Fatal(err)
					tx.Rollback()
					http.Error(w, "Error al ejecutar la consulta situaciones", http.StatusInternalServerError)
					return
				}
			}
		}
	}

	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}

	datos := Datos{
		Id: lastInsertId,
	}

	t, err := template.ParseFiles("Public/Html/Respuesta.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Renderizar la plantilla con los datos
	err = t.Execute(w, datos)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
