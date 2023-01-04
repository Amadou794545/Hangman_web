package main

import (
	"Hangman/modules"
	"fmt"
)

func Difficulty(difficulty string, GameState *modules.HangmanData) {
	if difficulty == "easy" {
		fmt.Printf("Easy --> Diffuclty Launch")
		file := "word.txt"
		GameState.Word = modules.InputWord(file)
		modules.RevealLetter(GameState.Word, GameState)
		GameState.Live = 10
	} else if difficulty == "medium" {
		fmt.Printf("Medium --> Diffuclty Launch")
		file := "word1.txt"
		GameState.Word = modules.InputWord(file)
		modules.RevealLetter(GameState.Word, GameState)
		GameState.Live = 10
	} else {
		fmt.Printf("Hard --> Diffuclty Launch")
		file := "word2.txt"
		GameState.Word = modules.InputWord(file)
		modules.RevealLetter(GameState.Word, GameState)
		GameState.Live = 10
	}
	return
}
