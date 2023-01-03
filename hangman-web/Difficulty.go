package main

import (
	"Hangman/modules"
)

func Difficulty(difficulty string, GameState *modules.HangmanData) {
	if difficulty == "easy" {
		file := "word.txt"
		GameState.Word = modules.InputWord(file)
		modules.RevealLetter(GameState.Word, GameState)
		GameState.Live = 10
	} else if difficulty == "medium" {
		file := "word1.txt"
		GameState.Word = modules.InputWord(file)
		modules.RevealLetter(GameState.Word, GameState)
		GameState.Live = 10
	} else {
		file := "word2.txt"
		GameState.Word = modules.InputWord(file)
		modules.RevealLetter(GameState.Word, GameState)
		GameState.Live = 10
	}
	return
}
