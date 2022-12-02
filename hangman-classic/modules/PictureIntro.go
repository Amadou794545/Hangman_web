package modules

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

func PictureIntro() string {
	var goodInput bool
	var UserInputDifficulty string
	var difficulty string
	Clean()
	PictureHistoire()
	time.Sleep(15 * time.Second)
	PaintRules()
	goodInput = false

	// Ask Difficulty to play the game
	for goodInput == false {
		fmt.Scan(&UserInputDifficulty)
		if UserInputDifficulty == "exit" { // Command to leave the program instantly
			ExitCommand()
		}
		if UserInputDifficulty != "easy" && UserInputDifficulty != "medium" && UserInputDifficulty != "hard" {
			fmt.Println("You can only chose the difficulty : easy | medium | hard ")
		} else {
			difficulty = UserInputDifficulty
			goodInput = true
		}
	}
	return difficulty
}
func PictureHistoire() {
	var lineText int
	var charToPrint string
	var timeToPrint int
	var CounterSong int
	var Yellow = "\033[33m"
	var Reset = "\u001b[37;1m"
	readFile, err := os.Open("Picture/HistoirePicture.txt") // Take the data from the txt file
	if err != nil {
		log.Fatal(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		timeToPrint -= 1
		if fileScanner.Text() == "1" {
			lineText = 23 // Line number to print
			timeToPrint = 1
			fmt.Printf("\033[198;24H")
			os.Stdout.Write([]byte("\033[1;0H"))
			os.Stdout.Write([]byte(charToPrint))
			charToPrint = ""
			CounterSong += 1
		}
		if lineText > 0 && timeToPrint <= 0 {
			for i := 0; i < len(fileScanner.Text())-1; i++ {
				charToPrint = charToPrint + string(fileScanner.Text()[i])
				if fileScanner.Text()[i+1] >= 'a' && fileScanner.Text()[i+1] <= 'z' || fileScanner.Text()[i+1] >= 'A' && fileScanner.Text()[i+1] <= 'Z' {
					charToPrint = charToPrint + Yellow
				}
				if fileScanner.Text()[i+1] == '.' {
					if fileScanner.Text()[i] >= 'a' && fileScanner.Text()[i] <= 'z' || fileScanner.Text()[i] >= 'A' && fileScanner.Text()[i] <= 'Z' || fileScanner.Text()[i] >= '0' && fileScanner.Text()[i] <= '9' {
					} else {
						charToPrint = charToPrint + Reset
					}
				}
			}
			charToPrint = charToPrint + "\n"
			lineText -= 1
		}
	}

}
func PaintRules() {
	var textPicture string
	Clean()
	readFile, err := os.Open("Picture/ReglesPicture.txt") // Take the data from the txt file
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
