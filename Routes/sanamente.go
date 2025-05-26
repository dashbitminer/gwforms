package Routes

import (
	"formulario/Controller"
	"formulario/Controller/Sanamente"
	"net/http"
)

func ConfigurarRutasSanamente(Cfg *Controller.Config) {
	//Form de Inscripcion
	http.HandleFunc("/sanamente/inscripcion/sv/pnc", func(w http.ResponseWriter, r *http.Request) { Sanamente.SanamenteInscripcionPNCsv(w, r, Cfg) })
	http.HandleFunc("/sanamente/inscripcion/sv/pnc/post", func(w http.ResponseWriter, r *http.Request) { Sanamente.Postsanamenteinscripcionpnc(w, r, Cfg) })
	http.HandleFunc("/sanamente/inscripcion/hn/new", func(w http.ResponseWriter, r *http.Request) { Sanamente.Sanamentehnnew(w, Cfg) })
	http.HandleFunc("/sanamente/inscripcion/hn/new/post", func(w http.ResponseWriter, r *http.Request) { Sanamente.Postsanamenteinscripcionnew(w, r, Cfg) })
	//Actividades de Cierre
	http.HandleFunc("/sanamente/moduloa/sv/pnc", func(w http.ResponseWriter, r *http.Request) { Sanamente.ModuloasvHandlerPNC(w, Cfg) })
	http.HandleFunc("/sanamente/modulob/sv/pnc", func(w http.ResponseWriter, r *http.Request) { Sanamente.ModulobsvHandlerPNC(w, Cfg) })
	http.HandleFunc("/sanamente/moduloc/sv/pnc", func(w http.ResponseWriter, r *http.Request) { Sanamente.ModulocsvHandlerPNC(w, Cfg) })
	http.HandleFunc("/sanamente/modulo/pnc/post", func(w http.ResponseWriter, r *http.Request) { Sanamente.Postmodulo(w, r, Cfg) })
	//Encuesta Red y Apoyos
	http.HandleFunc("/sanamente/Encuesta/RedyApoyos", func(w http.ResponseWriter, r *http.Request) { Sanamente.EncuestaRedApoyos(w, r, Cfg) })
	http.HandleFunc("/sanamente/Encuesta/RedyApoyos/post", func(w http.ResponseWriter, r *http.Request) { Sanamente.PostEncuestaRedApoyos(w, r, Cfg) })
	//Encuesta satisfacción | Intervencionista
	http.HandleFunc("/sanamente/Encuesta/Intervencionista", func(w http.ResponseWriter, r *http.Request) { Sanamente.EncuestaSatisfaccionInter(w, r, Cfg) })
	http.HandleFunc("/sanamente/Encuesta/Intervencionista/post", func(w http.ResponseWriter, r *http.Request) { Sanamente.PostEncuestaSatisfaccionInter(w, r, Cfg) })

}
