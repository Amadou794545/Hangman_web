package main

import (
	"html/template"
	"net/http"
)

//func Home(w http.ResponseWriter, r *http.Request) {
//	RenderTemplate(w, "home")
//}

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

	//switch r.Method {
	//case "POST":
	//	// Call ParseForm() to parse the raw query and update r.PostForm and r.Form.
	//	if err := r.ParseForm(); err != nil {
	//	fmt.Fprintf(w, "ParseForm() err: %v", err)
	//	return
	//	}
	//	fmt.Fprintf(w, "Post from website! r.PostFrom = %v\n", r.PostForm)
	//	name := r.FormValue("name")
	//	address := r.FormValue("address")
	//	fmt.Fprintf(w, "Name = %s\n", name)
	//	fmt.Fprintf(w, "Address = %s\n", address)
	//default:
	//		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	//	}
}

func Game(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "Game")
}
