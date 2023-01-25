package modules

func TryLetter(letter, word string) (string, bool) { // Function to try the letter chose by the user
	var indexTab string
	var tryWord bool // true == it's word ; false it's a letter
	if len(letter) == len(word) {
		tryWord = true
	} else {
		tryWord = false
	}
	if len(letter) == 1 {
		for i := 0; i < len(word); i++ {
			if letter[0] == word[i] {
				indexTab = indexTab + string(rune(i+33))
			}
		}
	} else {
		for i := 0; i < len(word); i++ {
			if letter == word {
				indexTab = indexTab + string(rune(i+33))
			}
		}
	}
	return indexTab, tryWord // Return the position/index of the letter in the word
}
