package handlers

import (
	"museum/data"
	"net/http"
	"text/template"
)

func HandleHello(w http.ResponseWriter, r *http.Request) {
	w.Write(([]byte)("Hello from GO"))
}

func HandleTemplate(w http.ResponseWriter, r *http.Request) {
	html, err := template.ParseFiles("templates/index.template")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	} else {
		html.Execute(w, data.GetAll())
	}
}