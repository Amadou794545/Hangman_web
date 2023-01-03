package modules

import (
	"fmt"
)

var goodLW string
var UsedLW []string

func AskLetter(word string, GameState *HangmanData) (string, []string) { // Ask GameState.UserLetter and check if it's between "a" and "z"

	_, GameState.UsedLW = repetitionLW(goodLW, true, nil, false) // Print letter/word already used

	if len(GameState.UserLetter) == len(word) {
		goodLW = GameState.UserLetter
	} else {
		goodLW = GameState.UserLetter
	}
	if GameState.UserLetter != "" {
		var goodLetter bool
		goodLetter, GameState.UsedLW = repetitionLW(goodLW, false, nil, false)
		if goodLetter == false {
			//faire quelque chose pour les répétitions de lettres
		}
	} else {
		return goodLW, GameState.UsedLW
	}
	return goodLW, GameState.UsedLW
}

// Function to know if the letter GameState.UserLetter is already used OR Print letter/word already used
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
