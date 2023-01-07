package main

func PrintJose(Live int, Picture string) string {
	switch Live {
	case 10:
		Picture = "/Assets/HANGMAN0.png"
	case 9:
		Picture = "/Assets/HANGMAN1.png"
	case 8:
		Picture = "/Assets/HANGMAN2.png"
	case 7:
		Picture = "/Assets/HANGMAN3.png"
	case 6:
		Picture = "/Assets/HANGMAN4.png"
	case 5:
		Picture = "/Assets/HANGMAN5.png"
	case 4:
		Picture = "/Assets/HANGMAN6.png"
	case 3:
		Picture = "/Assets/HANGMAN7.png"
	case 2:
		Picture = "/Assets/HANGMAN8.png"
	case 1:
		Picture = "/Assets/HANGMAN9.png"
	case 0:
		Picture = "/Assets/HANGMAN10.png"

	}
	return Picture
}
