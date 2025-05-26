package Routes

import (
	"formulario/Controller"
	"net/http"
)

func ConfigurarRutassql(Cfg *Controller.Config) {
	http.HandleFunc("/obtenermunicipiosql", func(w http.ResponseWriter, r *http.Request) { Controller.Obtenermunicipiosql(w, r, Cfg) })
}
