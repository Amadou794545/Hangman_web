package modules

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func PictureWL(result string) {
	var textPicture string
	Clean()
	if result == "lose" { // Print Lose Picture
		readFile, err := os.Open("Picture/SceneFinalDefaitePicture.txt") // Take the data from the txt file
		if err != nil {
			log.Fatal(err)
		}
		fileScanner := bufio.NewScanner(readFile)
		fileScanner.Split(bufio.ScanLines)
		for fileScanner.Scan() {
			textPicture = textPicture + fileScanner.Text() + "\n"
		}
		fmt.Printf("\033[198;24H")
		os.Stdout.Write([]byte("\033[1;0H"))
		os.Stdout.Write([]byte(textPicture))
	} else { // Print Win Picture
		readFile, err := os.Open("Picture/SceneFinalVictoirePicture.txt") // Take the data from the txt file
		if err != nil {
			log.Fatal(err)
		}
		fileScanner := bufio.NewScanner(readFile)
		fileScanner.Split(bufio.ScanLines)
		for fileScanner.Scan() {
			textPicture = textPicture + fileScanner.Text() + "\n"
		}
		fmt.Printf("\033[198;24H")
		os.Stdout.Write([]byte("\033[1;0H"))
		os.Stdout.Write([]byte(textPicture))
	}

}
