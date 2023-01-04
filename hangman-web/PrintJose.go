package main

func PrintJose(Live int, Picture string) string {
	switch Live {
	case 10:
		Picture = "/Assets/chiffre10.jpeg"
	case 9:
		Picture = "/Assets/chiffre9.jpeg"
	case 8:
		Picture = "/Assets/chiffre8.jpeg"
	case 7:
		Picture = "/Assets/chiffre7.png"
	case 6:
		Picture = "/Assets/chiffre6.png"
	case 5:
		Picture = "/Assets/chiffre5.jpg"
	case 4:
		Picture = "/Assets/chiffre4.jpeg"
	case 3:
		Picture = "/Assets/chiffre3.jpeg"
	case 2:
		Picture = "/Assets/chiffre2.jpeg"
	case 1:
		Picture = "/Assets/chiffre1.jpeg"
	case 0:
		Picture = "/Assets/chiffre0.jpeg"
	}
	return Picture
}
