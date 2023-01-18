package main

import (
	"github.com/gorilla/sessions"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"html/template"
	"net/http"
)

const port = ":443"

var tpl *template.Template

var db *gorm.DB
var store = sessions.NewCookieStore([]byte("super-secret-password"))

func main() {
	tpl, _ = template.ParseGlob("Templates/*.html")
	var err error
	db, err = gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	db.AutoMigrate(&User{}, &Session{})
	http.HandleFunc("/", Start)
	http.HandleFunc("/connexion", Connexion)
	http.HandleFunc("/home", Home)
	http.HandleFunc("/level", Level)
	http.HandleFunc("/game", Game)
	http.HandleFunc("/Congratulation", Congratulation)
	http.HandleFunc("/Loser", Loser)
	http.HandleFunc("/scoreboard", Scoreboard)
	http.HandleFunc("/inscription", Inscription)

	http.Handle("/Templates/", http.StripPrefix("/Templates/", http.FileServer(http.Dir("Templates"))))
	http.Handle("/Assets/", http.StripPrefix("/Assets/", http.FileServer(http.Dir("Assets"))))

	http.ListenAndServe(port, nil)
}
