package modules

func UpdateResult(letter string, index []int, tryWord bool, result []string) ([]string, bool) {
	if tryWord == false {
		if index == nil { // If the index is empty return same result and false (letter)
			return result, false
		} else { // Put the letter on his index position and return true (letter)
			for i := 0; i < len(index); i++ {
				result[index[i]] = letter
			}
		}
		return result, true
	} else { // Test word and return true for good result and false for invalid result
		if len(index) != len(result) {
			return result, false
		} else {
			for i := 0; i < len(index); i++ {
				result[i] = string(letter[i])
			}
			return result, true
		}
	}
}
