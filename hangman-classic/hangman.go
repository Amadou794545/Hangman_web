package main

import (
	"Hangman/modules"
	"fmt"
	"os"
	"time"
)

func main() {
	// Variables
	var hangData modules.HangmanData
	var resultCorrectLetter bool
	var tryWord bool
	var playWithSave bool
	var userInput string
	saveWithArg := true

	// Start
	if len(os.Args) > 1 {
		// Check if it's a start from save and verify the file
		if os.Args[1] == "--startWith" {
			file := os.Args[2]
			Check, err := os.Open(file)
			playWithSave = true
			if err != nil { // If file doesn't exist playWithSave is false
				playWithSave = false
				println("Your save is not correct , thanks to put an other name file !")
				saveWithArg = false
			}
			Check.Close()
			if playWithSave == true {
				jsonFile := ".json"
				indexJson := 4
				Count := 0
				for i := len(file) - 1; i > len(file)-5; i-- {
					if jsonFile[indexJson] != file[i] {
						Count += 1
					}
					if Count > 0 {
						playWithSave = false
					}
					indexJson -= 1
				}
			}
			// Now check with userInput
			for playWithSave == false {
				fmt.Scan(&userInput)
				try, err := os.Open(userInput)
				playWithSave = true
				if err != nil {
					playWithSave = false
				}
				try.Close()
				if playWithSave == false {
					println("Your save is not correct, thanks to put an other name file !")
				}
				if playWithSave == true {
					jsonFile := ".json"
					indexJson := 4
					Count := 0
					for i := len(userInput) - 1; i > len(userInput)-5; i-- {
						if jsonFile[indexJson] != userInput[i] {
							Count += 1
						}
						if Count > 0 {
							playWithSave = false
						}
						indexJson -= 1
					}
				}
			}
			if saveWithArg == true {
				hangData, _ = modules.StartWithSave(os.Args[2])
			} else {
				hangData, _ = modules.StartWithSave(userInput)
			}
			modules.Clean()
			modules.Pilote()
			modules.S_Pilot()
			modules.Clean()
		}
	}

	// Condition if the game is starting without save
	if playWithSave == false {
		// Hangman Version with animations or pictures, and choose the difficulty
		hangData.Version = modules.AnimOrPicture()
		if hangData.Version == "1" {
			modules.Unzip() // Download All files needed for animations and unzip
			hangData.Difficulty = modules.AnimIntro()
		} else {
			hangData.Difficulty = modules.PictureIntro()
		}
	}

	// Initialize the game without save
	if playWithSave != true {
		hangData.Live = 10
		hangData.Word = modules.InputWord()
		hangData.Result = modules.RevealLetter(hangData.Word, hangData.Difficulty)
		println("Good Luck, you have 10 attempts.")
		println()
	}
	if playWithSave == true {
		println()
		if hangData.Live != 10 {
			println(string(rune(hangData.Live)+48) + " attempts remaining")
		} else {
			println("10 attempts remaining")
		}

	}
	modules.AsciiArt(hangData.Result)
	if playWithSave != false {
		modules.UsedLW = hangData.UsedLW
	}

	// Start the game
	for hangData.Live > 0 { // If he has live again
		if modules.TestFinish(hangData.Result) == true {
			if hangData.Version == "1" {
				modules.AnimWL("victory")
			} else {
				modules.PictureWL("victory")
			}
			return
		}
		var UserLetter string
		UserLetter, hangData.UsedLW = modules.AskLetter(hangData.Word, hangData)
		if UserLetter == "STOP" {
			modules.StopAndSave(hangData)
		}
		hangData.Index, tryWord = modules.TryLetter(UserLetter, hangData.Word)                                            // Call function TryLetter
		hangData.Result, resultCorrectLetter = modules.UpdateResult(UserLetter, hangData.Index, tryWord, hangData.Result) // Call UpdateResult
		// Bad letter, Print hangman
		if resultCorrectLetter == false && tryWord == false {
			fmt.Println("Choose: " + UserLetter)
			hangData.Live -= 1
			if hangData.Version == "1" {
				modules.Anim(hangData.Live)
			} else {
				modules.Picture(hangData.Live)
			}
			if hangData.Live != 0 {
				println()
				println("Not present in the word, " + string(rune(hangData.Live)+48) + " attempts remaining")
				modules.AsciiArt(hangData.Result)
			}
			// Bad word, Print hangman
		} else if resultCorrectLetter == false && tryWord == true {
			fmt.Println("Choose: " + UserLetter)
			for i := 0; i < 2; i++ {
				hangData.Live -= 1
				if hangData.Version == "1" {
					modules.Anim(hangData.Live)
				} else {
					modules.Picture(hangData.Live)
				}
				time.Sleep(2 * time.Second)
			}
			if hangData.Live != 0 {
				println()
				println("It's not the good word, " + string(rune(hangData.Live)+48) + " attempts remaining")
				modules.AsciiArt(hangData.Result)
			}
		} else { // Good letter or word
			fmt.Println("Choose: " + UserLetter)
			modules.AsciiArt(hangData.Result)
			println()
		}
	}

	// If live == 0, Print Lose Scene
	if hangData.Version == "1" {
		modules.AnimWL("lose")
	} else {
		modules.PictureWL("lose")
	}
}
