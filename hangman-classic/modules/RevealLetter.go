package modules

import (
	"math/rand"
	"time"
)

func RevealLetter(word string, difficulty string) []string { // Function to reveal random letters depending on the difficulty
	var finalPrint []string
	for i := 0; i < len(word); i++ { // The final print is full of underscore
		finalPrint = append(finalPrint, "_")
	}
	nbLetterToReveal := Difficulties(word, difficulty) // Number of letter that we need to reveal
	var possibleLetter []string
	var letter string
	var recurrenceLetter int
	for i := 0; i < len(word); i++ { // If recurrence of a letter is less or equal to nbLetterToReveal, letter add to possibleLetter
		letter = string(word[i])
		for j := 0; j < len(word); j++ {
			if word[i] == word[j] {
				recurrenceLetter++
			}
		}
		for k := 0; k < len(possibleLetter); k++ { // If the letter is already in possible letter don't add the letter, for equal luck on every possible letter
			if letter == possibleLetter[k] {
				recurrenceLetter += 1000
			}
		}
		if recurrenceLetter <= nbLetterToReveal {
			possibleLetter = append(possibleLetter, letter)
		}
		recurrenceLetter = 0
	}
	for nbLetterToReveal != 0 { // Chose a letter in the possibleLetter and subtract to nbLetterToReveal
		nextTest := nbLetterToReveal
		rand.Seed(time.Now().UnixNano())
		charset := possibleLetter                      // possibleLetter is a string with the letter who can choose to reveal
		randomChar := charset[rand.Intn(len(charset))] // Variable randomChar take a random character from possibleLetter
		var indexTab []int
		for i := 0; i < len(word); i++ {
			if randomChar[0] == word[i] {
				indexTab = append(indexTab, i)
			}
		}
		nextTest -= len(indexTab)
		if randomChar != "0" && nextTest >= 0 { // Swap letter with the underscore in the final print
			nbLetterToReveal -= len(indexTab)
			for i := 0; i < len(indexTab); i++ {
				finalPrint[indexTab[i]] = randomChar
			}
			for i := 0; i < len(possibleLetter); i++ { // If the letter is swapped, the possibleLetter become a zero
				if randomChar == possibleLetter[i] {
					possibleLetter[i] = "0"
				}
			}
		}
	}
	repetitionLW("nil", true, finalPrint, true)
	return finalPrint
}

// Function to know what the user choose for the difficulty
func Difficulties(word string, userChoice string) int {
	var difficulty int
	if userChoice == "easy" {
		difficulty = len(word) / 2
	} else if userChoice == "medium" {
		difficulty = len(word)/2 - 1
	} else if userChoice == "hard" {
		difficulty = 0
	}
	return difficulty
}
