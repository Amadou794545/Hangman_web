package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func Level(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "Level")

}

func Start(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "Start")
	if r.URL.Path != "/" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}
}

func RenderTemplate(w http.ResponseWriter, tmpl string) {
	t, err := template.ParseFiles("./Templates/" + tmpl + ".html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	t.Execute(w, nil)
}

func Home(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "Home")

	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Address = %s\n", address)

}

func Game(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "Game")
}
