package Routes

import (
	"formulario/Controller"
	"formulario/Controller/Salud"
	"formulario/Controller/Salud/CdN"
	"net/http"
)

func ConfigurarRutassalud(Cfg *Controller.Config) {
	//
	http.HandleFunc("/salud/cdn/taller1", func(w http.ResponseWriter, r *http.Request) { CdN.Taller1CdNHandler(w, Cfg) })
	http.HandleFunc("/salud/cdn/taller2", func(w http.ResponseWriter, r *http.Request) { CdN.Taller2CdNHandler(w, Cfg) })
	http.HandleFunc("/salud/cdn/taller3", func(w http.ResponseWriter, r *http.Request) { CdN.Taller3CdNHandler(w, Cfg) })
	http.HandleFunc("/salud/cdn/generoregional", func(w http.ResponseWriter, r *http.Request) { CdN.Genero(w) })
	http.HandleFunc("/salud/cdn/genero/sv", func(w http.ResponseWriter, r *http.Request) { CdN.GeneroSV(w) })
	http.HandleFunc("/salud/hn/new", func(w http.ResponseWriter, r *http.Request) { Salud.Saludhnnew(w, Cfg) })
	// Funciones para JS
	http.HandleFunc("/obtenerdepsedesalud", func(w http.ResponseWriter, r *http.Request) { CdN.ObtenerdepsedeHandlersalud(w, r, Cfg) })
	// POST
	http.HandleFunc("/salud/cdn/taller/post", func(w http.ResponseWriter, r *http.Request) { CdN.PostTallerCdN(w, r, Cfg) })
	http.HandleFunc("/salud/consentimiento/mx", func(w http.ResponseWriter, r *http.Request) { Salud.Consentimientomx(w, Cfg) })
	http.HandleFunc("/salud/consentimiento/virtual/gt", func(w http.ResponseWriter, r *http.Request) { Salud.Consentimientogt(w, Cfg) })
	http.HandleFunc("/salud/consentimiento/mx/post", func(w http.ResponseWriter, r *http.Request) { Salud.Consentimientopost(w, r, Cfg) })
	http.HandleFunc("/salud/consentimiento/gt/post", func(w http.ResponseWriter, r *http.Request) { Salud.Consentimientopostgt(w, r, Cfg) })
	//	http.HandleFunc("/salud/consentimiento/gt/correo", func(w http.ResponseWriter, r *http.Request) { Salud.EmailSalud(w, r, Cfg) })
	http.HandleFunc("/salud/consentimiento/gt/correo", func(w http.ResponseWriter, r *http.Request) { Salud.EmailSalud(w, r) })
	http.HandleFunc("/salud/cdn/genero/post", func(w http.ResponseWriter, r *http.Request) { CdN.PostCuestionarioRegionalGenero(w, r, Cfg) })
	http.HandleFunc("/salud/hn/new/post", func(w http.ResponseWriter, r *http.Request) { Salud.Postesaludhnnew(w, r, Cfg) })
}
