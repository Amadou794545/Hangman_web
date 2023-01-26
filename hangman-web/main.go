package main

import (
	"archive/zip"
	"fmt"
	"github.com/gorilla/sessions"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"html/template"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

const port = ":443"

var tpl *template.Template

var db *gorm.DB
var store = sessions.NewCookieStore([]byte("super-secret-password"))

func main() {
	Unzip()
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
	http.Handle("/Music/", http.StripPrefix("/Music/", http.FileServer(http.Dir("Music"))))

	http.ListenAndServe(port, nil)
}

func Unzip() {
	DownloadAnim()
	dst := "Music"
	archive, err := zip.OpenReader("Music.zip")
	if err != nil {
		panic(err)
	}
	defer archive.Close()

	for _, f := range archive.File {
		filePath := filepath.Join(dst, f.Name)

		if !strings.HasPrefix(filePath, filepath.Clean(dst)+string(os.PathSeparator)) {
			fmt.Println("Invalid file path")
			return
		}
		if f.FileInfo().IsDir() {
			os.MkdirAll(filePath, os.ModePerm)
			continue
		}

		if err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
			panic(err)
		}

		dstFile, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			panic(err)
		}

		fileInArchive, err := f.Open()
		if err != nil {
			panic(err)
		}

		if _, err := io.Copy(dstFile, fileInArchive); err != nil {
			panic(err)
		}

		dstFile.Close()
		fileInArchive.Close()
	}
} // Unzip and create file
func DownloadAnim() {
	url := "https://drive.google.com/uc?export=download&id=1EHuyaZUIGI445WEnlsn2VWuHqD1FYqEb"
	fileName := "Music.zip"

	output, err := os.Create(fileName)
	defer output.Close()

	response, err := http.Get(url)
	if err != nil {
		return
	}
	defer response.Body.Close()

	_, err = io.Copy(output, response.Body)

} // Download all files of animation
