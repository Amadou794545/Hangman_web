package modules

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
)

func AnimOrPicture() string {
	var userInput string
	var version string
	goodInput := false
	PaintAOP()
	// Ask version and Retry if is not correct
	for goodInput == false {
		fmt.Scan(&userInput)
		if userInput == "exit" { // Command to leave the program instantly
			ExitCommand()
		}
		if userInput != "1" && userInput != "2" {
			fmt.Println("You can only chose the version 1 or 2 !!")
		} else {
			version = userInput
			goodInput = true
		}
	}
	return version
}

// Print Ask Anime Or Picture
func PaintAOP() {
	var lineText string
	Clean()
	readFile, err := os.Open("Picture/AnimOrPicture.txt")
	if err != nil {
		log.Fatal(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		lineText = lineText + fileScanner.Text() + "\n"
	}
	fmt.Printf("\033[198;24H")
	os.Stdout.Write([]byte("\033[1;0H"))
	os.Stdout.Write([]byte(lineText))
}

// Clear Terminal
func Clean() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
