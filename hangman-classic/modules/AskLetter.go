package modules

import (
	"fmt"
)

var goodLW string
var UsedLW []string

func AskLetter(word string, hangData HangmanData) (string, []string) { // Ask userInput and check if it's between "a" and "z"
	var userInput string
	_, hangData.UsedLW = repetitionLW(goodLW, true, nil, false) // Print letter/word already used
	fmt.Println("Chose letter or word to try to find the word")
	fmt.Scan(&userInput)
	if userInput == "exit" { // Command to leave the program instantly
		ExitCommand()
	}
	if userInput == "STOP" {
		goodLW = userInput
		return goodLW, hangData.UsedLW
	}
	if len(userInput) == 1 && userInput < "a" || len(userInput) == 1 && userInput > "z" { // Only input lowercase, another thing excludes
		userInput = ""
		fmt.Println("Error : you can only chose a lowercase for Hangman")
		AskLetter(word, hangData) // User put other things than lowercase, so we call AskLetter and put an error message
	} else if len(userInput) == len(word) {
		for i := 0; i < len(userInput); i++ {
			if string(userInput[i]) < "a" || string(userInput[i]) > "z" {
				userInput = ""
				fmt.Println("Error : you can only chose lowercase characters for Hangman")
				AskLetter(word, hangData) // User put other things than lowercase, so we call AskLetter and put an error message
			}
		}
		goodLW = userInput
	} else if len(userInput) > 1 && len(userInput) < len(word) || len(userInput) > len(word) { // Check if the length of userInput is equal to the length of the word
		fmt.Println("Error : you can only put 1 letter or a word with the same number of letter than the word to search!!!!!!")
		userInput = ""
		AskLetter(word, hangData)
	} else {
		goodLW = userInput
	}
	if userInput != "" {
		var goodLetter bool
		goodLetter, hangData.UsedLW = repetitionLW(goodLW, false, nil, false)
		if goodLetter == false {
			AskLetter(word, hangData)
		}
	} else {
		return goodLW, hangData.UsedLW
	}
	return goodLW, hangData.UsedLW
}

// Function to know if the letter userInput is already used OR Print letter/word already used
func repetitionLW(goodLW string, printUsed bool, reveal []string, revealW bool) (bool, []string) {
	var letterNotUsed int
	var countLetterAlreadyUsed int
	println()
	if revealW == false {
		if printUsed == false {
			if UsedLW == nil {
				UsedLW = append(UsedLW, goodLW)
			} else {
				for _, lw := range UsedLW {
					if goodLW == lw {
						letterNotUsed += 1
					}
				}
				if letterNotUsed == 0 {
					UsedLW = append(UsedLW, goodLW)
				} else {
					fmt.Println("You already used this, check the list of already used input !!")
					return false, UsedLW
				}
			}
			return true, UsedLW
		} else { // Print Letter already used
			for i := 0; i < len(UsedLW); i++ {
				if UsedLW != nil {
					if i == 0 {
						fmt.Printf("Letter or Word already used :")
						fmt.Printf(" " + UsedLW[i] + ",")
					} else {
						fmt.Printf(" " + UsedLW[i] + ",")
					}
				}
			}
			if UsedLW != nil {
				println()
			}
			return true, UsedLW
		}
	} else { // Add letter from the reveal
		for _, word := range reveal {
			countLetterAlreadyUsed = 0
			if word != "_" && UsedLW != nil {
				for _, usedLetter := range UsedLW {
					if word == usedLetter {
						countLetterAlreadyUsed += 1
					}
				}
				if countLetterAlreadyUsed == 0 {
					UsedLW = append(UsedLW, word)
				}
			}
			if word != "_" && UsedLW == nil {
				UsedLW = append(UsedLW, word)
			}
		}
		return true, UsedLW
	}
}
