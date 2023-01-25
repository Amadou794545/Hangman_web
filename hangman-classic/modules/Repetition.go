package modules

import (
	"strings"
)

func Repetition(revealW bool, UsedLW string, UserLetter string, AlreadyUsed string, Result string) (bool, string, string, string, string) {
	var countLetterAlreadyUsed int
	if revealW == false {
		usedWords := strings.Split(UsedLW, ",")
		found := false
		for _, word := range usedWords {
			if word == UserLetter {
				found = true
				break
			}
		}
		if found {
			AlreadyUsed = "You already used this, check the list of already used input !!"
			return true, UsedLW, UserLetter, AlreadyUsed, Result // Already used
		}
		UsedLW = UsedLW + UserLetter
		UsedLW = UsedLW + ","
		AlreadyUsed = ""
		return false, UsedLW, UserLetter, AlreadyUsed, Result // Not already used
	} else { // Add letter from the reveal
		for _, word := range Result {
			countLetterAlreadyUsed = 0
			if string(word) != "_" && UsedLW != "" { // diff from nil because at start it's equal NIL
				for _, usedLetter := range UsedLW {
					if word == usedLetter {
						countLetterAlreadyUsed += 1
					}
				}
				if countLetterAlreadyUsed == 0 {
					UsedLW = UsedLW + string(word)
					UsedLW = UsedLW + ","

				}
			}
			if string(word) != "_" && UsedLW == "" {
				UsedLW = UsedLW + string(word)
				UsedLW = UsedLW + ","
			}
		}
		return false, UsedLW, UserLetter, AlreadyUsed, Result
	}
}
