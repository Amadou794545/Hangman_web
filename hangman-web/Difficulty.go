package main

import (
	"Hangman/modules"
)

func Difficulty(difficulty string, GameState *HangmanData) {
	if difficulty == "easy" {
		file := "words.txt"
		GameState.Word = modules.InputWord(file)
		GameState.Result = modules.RevealLetter(GameState.Word, GameState.Result)
		_, GameState.UsedLW, GameState.UserLetter, GameState.AlreadyUsed, GameState.Result = modules.Repetition(true, GameState.UsedLW, GameState.UserLetter, GameState.AlreadyUsed, GameState.Result)
		GameState.Live = 10
		GameState.GCompleted = "false"
		GameState.GameDifficulty = difficulty
	} else if difficulty == "medium" {
		file := "words2.txt"
		GameState.Word = modules.InputWord(file)
		GameState.Result = modules.RevealLetter(GameState.Word, GameState.Result)
		_, GameState.UsedLW, GameState.UserLetter, GameState.AlreadyUsed, GameState.Result = modules.Repetition(true, GameState.UsedLW, GameState.UserLetter, GameState.AlreadyUsed, GameState.Result)
		GameState.Live = 10
		GameState.GCompleted = "false"
		GameState.GameDifficulty = difficulty
	} else {
		file := "words3.txt"
		GameState.Word = modules.InputWord(file)
		GameState.Result = modules.RevealLetter(GameState.Word, GameState.Result)
		_, GameState.UsedLW, GameState.UserLetter, GameState.AlreadyUsed, GameState.Result = modules.Repetition(true, GameState.UsedLW, GameState.UserLetter, GameState.AlreadyUsed, GameState.Result)
		GameState.Live = 10
		GameState.GCompleted = "false"
		GameState.GameDifficulty = difficulty
	}
	return
}
