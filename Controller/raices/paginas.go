package Raices

import (
	"net/http"
	"text/template"
)

func Preraices(w http.ResponseWriter) {

	// Cargar plantilla HTML desde archivo
	tmpl, err := template.ParseFiles("public/HTML/Participantes/Raices/pre.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Renderizar plantilla HTML con opciones de preguntas select
	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
