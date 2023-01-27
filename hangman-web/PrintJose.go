package main

func PrintJose(Live int, Picture string) string {

	//Switch the png to display with the live in parameters
	switch Live {
	case 10:
		Picture = "/Assets/zombie.png"
	case 9:
		Picture = "/Assets/zombie1.png"
	case 8:
		Picture = "/Assets/zombie2.png"
	case 7:
		Picture = "/Assets/zombie3.png"
	case 6:
		Picture = "/Assets/zombie4.png"
	case 5:
		Picture = "/Assets/zombie5.png"
	case 4:
		Picture = "/Assets/zombie6.png"
	case 3:
		Picture = "/Assets/zombie7.png"
	case 2:
		Picture = "/Assets/zombie8.png"
	case 1:
		Picture = "/Assets/zombie9.png"
	case 0:
		Picture = "/Assets/zombie.png"

	}
	return Picture
}
