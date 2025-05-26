package Routes

import (
	"formulario/Controller"
	"formulario/Controller/Voluntariado"
	"net/http"
)

func ConfigurarRutasvoluntariado(Cfg *Controller.Config) {
	//Form encuesta de satisfacción post evento
	http.HandleFunc("/voluntariado/encuesta/evento", func(w http.ResponseWriter, r *http.Request) { Voluntariado.SatisfaccionEvento(w, Cfg) })
	http.HandleFunc("/voluntariado/encuesta/evento/post", func(w http.ResponseWriter, r *http.Request) { Voluntariado.PostSatisfaccionEvento(w, r, Cfg) })
	http.HandleFunc("/voluntariado/encuesta/respuesta", func(w http.ResponseWriter, r *http.Request) { Voluntariado.RespuestaEncuestaVol(w) })
}
