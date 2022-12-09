package main

import (
	"net/http"
)

const port = ":8080"

func main() {
	http.HandleFunc("/", Start)
	http.HandleFunc("/home", Home)
	http.HandleFunc("/level", Level)
	http.HandleFunc("/game", Game)

	http.Handle("/Templates/", http.StripPrefix("/Templates/", http.FileServer(http.Dir("Templates"))))
	http.Handle("/Assets/", http.StripPrefix("/Assets/", http.FileServer(http.Dir("Assets"))))

	http.ListenAndServe(port, nil)
}
