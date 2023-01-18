package main

import (
	"Hangman/modules"
	"fmt"
	"html/template"
	"net/http"
)

type DiffStruc struct {
	easy   string
	medium string
	hard   string
	Reveal []string
}

var Diff DiffStruc

var GameState modules.HangmanData

var tryWord bool
var resultCorrectLetter bool

func Level(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("./Templates/level.html")
	Diff.easy = r.FormValue("Easy")
	Diff.medium = r.FormValue("Medium")
	Diff.hard = r.FormValue("Hard")
	if Diff.easy != "" {
		fmt.Printf("Easy --> Activate")
		Difficulty(Diff.easy, &GameState)
		http.Redirect(w, r, "/game", http.StatusFound)
		GameState.Picture = "/Assets/HANGMAN0.png"
	} else if Diff.medium != "" {
		fmt.Printf("Medium --> Activate")
		Difficulty(Diff.medium, &GameState)
		http.Redirect(w, r, "/game", http.StatusFound)
		GameState.Picture = "/Assets/HANGMAN0.png"
	} else if Diff.hard != "" {
		fmt.Printf("Hard --> Activate")
		Difficulty(Diff.hard, &GameState)
		http.Redirect(w, r, "/game", http.StatusFound)
		GameState.Picture = "/Assets/HANGMAN0.png"
	}
	t.Execute(w, Diff)
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

}

func Game(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./Templates/game.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	GameState.UserLetter = r.FormValue("Enter")
	if GameState.UserLetter != "" {
		if GameState.Live > 0 {
			fmt.Printf("%s --> UserLetter\n", GameState.UserLetter)
			fmt.Printf("%s --> Result\n", GameState.Result)
			GameState.OnlyLowerCase = ""
			GameState.BadInput = ""
			for _, letter := range GameState.UserLetter {
				if string(letter) < "a" || string(letter) > "z" {
					GameState.OnlyLowerCase = "Error : you can only chose lowercase characters for Hangman"
				}
			}
			if GameState.OnlyLowerCase == "" {
				if len(GameState.UserLetter) == 1 || len(GameState.UserLetter) == len(GameState.Word) {
					if modules.Repetition(false, &GameState) == false {
						GameState.Index, tryWord = modules.TryLetter(GameState.UserLetter, GameState.Word)
						GameState.Result, resultCorrectLetter = modules.UpdateResult(GameState.UserLetter, GameState.Index, tryWord, GameState.Result)

						if resultCorrectLetter == false && tryWord == false {
							fmt.Println("Choose: " + GameState.UserLetter)
							GameState.Live -= 1
							GameState.Picture = PrintJose(GameState.Live, GameState.Picture)
							if GameState.Live != 0 {
								println()
								println("Not present in the word, " + string(rune(GameState.Live)+48) + " attempts remaining")
							}
							// Bad word, Print hangman
						} else if resultCorrectLetter == false && tryWord == true {
							fmt.Println("Choose: " + GameState.UserLetter)
							GameState.Live -= 2
							GameState.Picture = PrintJose(GameState.Live, GameState.Picture)

							if GameState.Live != 0 {
								println()
								println("It's not the good word, " + string(rune(GameState.Live)+48) + " attempts remaining")
							}
						} else { // Good letter or word
						}
					}
					if modules.TestFinish(GameState.Result) == true {
						http.Redirect(w, r, "Congratulation", http.StatusFound)
						GameState = modules.HangmanData{}
					}
					fmt.Printf("%d --> Nombre de vie restantes", GameState.Live)
					if GameState.Live == 0 || GameState.Live < 0 {
						http.Redirect(w, r, "/Loser", http.StatusFound)
					}
				} else {
					GameState.BadInput = "Error : you can only put 1 letter or a word with the same number of letter than the word to search!!!!!!"
				}
			}
		}
	}
	t.Execute(w, GameState)
}

func Congratulation(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "Congratulation")
}

func Loser(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "Loser")

}
func Scoreboard(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "Scoreboard")
}
func Inscription(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "Inscription")
}
func Connexion(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "Connexion")
}
