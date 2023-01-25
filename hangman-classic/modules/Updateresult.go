package modules

func UpdateResult(letter string, index string, tryWord bool, result string) (string, bool) {
	if tryWord == false {
		if index == "" { // If the index is empty return same result and false (letter)
			return result, false
		} else { // Put the letter on his index position and return true (letter)
			for i := 0; i < len(index); i++ {
				idx := int(index[i]) - 33
				resultRune := []rune(result)
				resultRune[idx] = []rune(letter)[0]
				result = string(resultRune)
			}
		}
		return result, true
	} else { // Test word and return true for good result and false for invalid result
		if len(index) != len(result) {
			return result, false
		} else {
			for i := 0; i < len(index); i++ {
				idx := int(index[i]) - 33
				resultRune := []rune(result)
				resultRune[idx] = []rune(letter)[i]
				result = string(resultRune)
			}
			return result, true
		}
	}
}
