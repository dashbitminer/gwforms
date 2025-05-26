package Routes

import (
	"formulario/Controller"
	"formulario/Controller/Raices"
	"net/http"
)

func ConfigurarRutasraices(Cfg *Controller.Config) {
	http.HandleFunc("/raices/preinscripcion", func(w http.ResponseWriter, r *http.Request) { Raices.Preraices(w) })
	http.HandleFunc("/raices/preinscripcion/post", func(w http.ResponseWriter, r *http.Request) { Raices.PostPre(w, r, Cfg) })
}
