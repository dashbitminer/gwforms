package Routes

import (
	"formulario/Controller"
	"formulario/Controller/Juventudes"
	"net/http"
)

func ConfigurarRutasjuventudes(Cfg *Controller.Config) {
	http.HandleFunc("/juventud/cp/preinscripcion", func(w http.ResponseWriter, r *http.Request) { Juventudes.Prechemonics(w) })
	http.HandleFunc("/juventud/cp/preinscripcion/post", func(w http.ResponseWriter, r *http.Request) { Juventudes.PostPre(w, r, Cfg) })
}
