package modules

import (
	"fmt"
	"time"
)

type HangmanData struct {
	Difficulty  string
	Live        int
	Word        string
	Result      []string
	UsedLW      []string
	Index       []int
	UserLetter  string
	Picture     string
	AlreadyUsed string
}

func Hangman(hangData HangmanData) {
	// Variables
	var resultCorrectLetter bool
	var tryWord bool

	// Initialize the game without save

	hangData.Live = 10
	//hangData.Word = InputWord()
	//hangData.Result = RevealLetter(hangData.Word)
	println("Good Luck, you have 10 attempts.")
	println()

	AsciiArt(hangData.Result)

	// Start the game
	for hangData.Live > 0 { // If he has live again
		if TestFinish(hangData.Result) == true {
			//PictureWL("victory")
			return
		}
		var UserLetter string
		//UserLetter, hangData.UsedLW = AskLetter(hangData.Word, &hangData)

		hangData.Index, tryWord = TryLetter(UserLetter, hangData.Word)                                            // Call function TryLetter
		hangData.Result, resultCorrectLetter = UpdateResult(UserLetter, hangData.Index, tryWord, hangData.Result) // Call UpdateResult
		// Bad letter, Print hangman
		if resultCorrectLetter == false && tryWord == false {
			fmt.Println("Choose: " + UserLetter)
			hangData.Live -= 1
			//Picture(hangData.Live)
			if hangData.Live != 0 {
				println()
				println("Not present in the word, " + string(rune(hangData.Live)+48) + " attempts remaining")
				AsciiArt(hangData.Result)
			}
			// Bad word, Print hangman
		} else if resultCorrectLetter == false && tryWord == true {
			fmt.Println("Choose: " + UserLetter)
			for i := 0; i < 2; i++ {
				hangData.Live -= 1
				//Picture(hangData.Live)
				time.Sleep(2 * time.Second)
			}
			if hangData.Live != 0 {
				println()
				println("It's not the good word, " + string(rune(hangData.Live)+48) + " attempts remaining")
				AsciiArt(hangData.Result)
			}
		} else { // Good letter or word
			fmt.Println("Choose: " + UserLetter)
			AsciiArt(hangData.Result)
			println()
		}
	}
	//PictureWL("lose")
}
