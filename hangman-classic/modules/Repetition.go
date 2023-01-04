package modules

func Repetition(revealW bool, GameState *HangmanData) bool {
	var countLetterAlreadyUsed int
	if revealW == false {
		for _, usedLW := range GameState.UsedLW {
			if GameState.UserLetter == usedLW {
				GameState.AlreadyUsed = "You already used this, check the list of already used input !!"
				return true // Already used
			}
		}
		GameState.UsedLW = append(GameState.UsedLW, GameState.UserLetter)
		GameState.AlreadyUsed = ""
		return false // Not already used
	} else { // Add letter from the reveal
		for _, word := range GameState.Result {
			countLetterAlreadyUsed = 0
			if word != "_" && GameState.UsedLW != nil { // diff from nil because at start it's equal NIL
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
		return false
	}
}
