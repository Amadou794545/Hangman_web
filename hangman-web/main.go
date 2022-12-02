package main

import (
	"net/http"
)

const port = ":8080"

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/contact", Contact)
	http.HandleFunc("/jouer", Jouer)

	//http.Handle("/Templates/", http.StripPrefix("/Templates/", http.FileServer(http.Dir("Templates"))))
	http.Handle("/Assets/", http.StripPrefix("/Assets/", http.FileServer(http.Dir("Assets"))))

	http.ListenAndServe(port, nil)

}
