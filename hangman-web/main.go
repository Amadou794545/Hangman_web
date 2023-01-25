package main

import (
	"github.com/gorilla/sessions"
	"gorm.io/driver/postgres"
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
	dsn := "host=localhost user=postgres password=123 dbname=postgres port=5432 sslmode=disable TimeZone=Europe/Paris"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	db.AutoMigrate(&User{}, &Session{}, &HangmanData{}, &ScoreBoard{})
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
