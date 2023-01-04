package modules

import (
	"fmt"
)

var goodLW string

func AskLetter(word string, GameState *HangmanData) (string, []string) { // Ask GameState.UserLetter and check if it's between "a" and "z"

	_, GameState.UsedLW = repetitionLW(goodLW, true, nil, false, GameState) // Print letter/word already used

	if len(GameState.UserLetter) == len(word) {
		goodLW = GameState.UserLetter
	} else {
		goodLW = GameState.UserLetter
	}
	if GameState.UserLetter != "" {
		var goodLetter bool
		goodLetter, GameState.UsedLW = repetitionLW(goodLW, false, nil, false, GameState)
		if goodLetter == false {
			GameState.AlreadyUsed = "You already used this, check the list of already used input !!"
		}
	} else {
		return goodLW, GameState.UsedLW
	}
	return goodLW, GameState.UsedLW
}

// Function to know if the letter GameState.UserLetter is already used OR Print letter/word already used
func repetitionLW(goodLW string, printUsed bool, reveal []string, revealW bool, GameState *HangmanData) (bool, []string) {
	var letterNotUsed int
	var countLetterAlreadyUsed int
	println()
	if revealW == false {
		if printUsed == false {
			if GameState.UsedLW == nil {
				GameState.UsedLW = append(GameState.UsedLW, goodLW)
			} else {
				for _, lw := range GameState.UsedLW {
					if goodLW == lw {
						letterNotUsed += 1
					}
				}
				if letterNotUsed == 0 {
					GameState.UsedLW = append(GameState.UsedLW, goodLW)
				} else {
					fmt.Println("You already used this, check the list of already used input !!")
					return false, GameState.UsedLW
				}
			}
			return true, GameState.UsedLW
		} else { // Print Letter already used
			for i := 0; i < len(GameState.UsedLW); i++ {
				if GameState.UsedLW != nil {
					if i == 0 {
						fmt.Printf("Letter or Word already used :")
						fmt.Printf(" " + GameState.UsedLW[i] + ",")
					} else {
						fmt.Printf(" " + GameState.UsedLW[i] + ",")
					}
				}
			}
			if GameState.UsedLW != nil {
				println()
			}
			return true, GameState.UsedLW
		}
	} else { // Add letter from the reveal
		for _, word := range reveal {
			countLetterAlreadyUsed = 0
			if word != "_" && GameState.UsedLW != nil {
				for _, usedLetter := range GameState.UsedLW {
					if word == usedLetter {
						countLetterAlreadyUsed += 1
					}
				}
				if countLetterAlreadyUsed == 0 {
					GameState.UsedLW = append(GameState.UsedLW, word)
				}
			}
			if word != "_" && GameState.UsedLW == nil {
				GameState.UsedLW = append(GameState.UsedLW, word)
			}
		}
		return true, GameState.UsedLW
	}
}
