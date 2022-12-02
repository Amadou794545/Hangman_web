package modules

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func InputWord() string { // Test the file and chose a valid word inside
	arguments := os.Args
	var file []string
	if len(arguments) < 2 || len(arguments) > 2 { // If he puts less or more than one argument
		println("There is no file name in the arguments or too much arguments")
		badFile := AskFile(0)          // Ask file and keep only word from the file
		for _, word := range badFile { // Keep only valid word
			if word != "" {
				file = append(file, word)
			}
		}
		file = VerifyRecurrenceWord(file)
		rand.Seed(time.Now().Unix())
		b := rand.Perm(len(file))
		word := file[b[0]]
		return word
	} else {
		fileBad := AskFile(1)          // Verify arguments file and select one word inside or ask for another file
		for _, word := range fileBad { // Keep only valid word
			if word != "" {
				file = append(file, word)
			}
		}
		file = VerifyRecurrenceWord(file)
		rand.Seed(time.Now().Unix())
		b := rand.Perm(len(file))
		word := file[b[0]]
		return word
	}
}

var availableWord []string

func AskFile(GoodArg int) []string { // Ask fileInput to the user and check if it's between "a" and "z"
	var fileInput string
	var possLetter []string
	var goodFile bool
	if GoodArg == 0 { // if the fileInput is not given, we ask to give
		fmt.Println("Give a fileInput name with word in to start the game !")
		fmt.Scan(&fileInput)
		if fileInput == "exit" { // Command to leave the program instantly
			ExitCommand()
		}
	} else {
		fileInput = os.Args[1]
		GoodArg = 0
	}
	possLetter, goodFile = VerifyWordFromFile(fileInput) // Test the fileInput and take valid word
	availableWord = append(availableWord, possLetter...)
	if goodFile == false {
		fileInput = ""
		fmt.Println("Error : your fileInput is no correct to play the game !")
		AskFile(GoodArg) // User put other things than lowercase so we call AskLetter and put an error message
	}
	return availableWord
}

func VerifyWordFromFile(file string) ([]string, bool) { // Function to chose randomly a word in the file
	var possibleLetter []string
	var badChar int // If a char is not include between a and z bad take 1 more
	f, err := os.Open(file)
	if err != nil {
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() { // Scan every word and select only valid word in lowercase
		badChar = 0
		for _, letter := range scanner.Text() {
			if letter < 'a' || letter > 'z' {
				badChar += 1
			}
		}
		if badChar == 0 { // If not bad char , add word to Final possibleLetter
			possibleLetter = append(possibleLetter, []string{scanner.Text()}...)
		}
	}
	if err := scanner.Err(); err != nil {
	}
	if len(possibleLetter) < 2 { // If possibleLetter have less than 2 word return false/possibleLetter or more than 2 return true/possibleLetter
		if len(possibleLetter) == 1 {
			possibleLetter[0] = ""
		} else {
			possibleLetter = append(possibleLetter, "")
		}
		return possibleLetter, false
	}
	return possibleLetter, true
}

func VerifyRecurrenceWord(TextFileWithRecurrence []string) []string {
	var counter int
	var perfectTextFile []string
	for i := 0; i < len(TextFileWithRecurrence); i++ {
		counter = 0
		for j := 0; j < len(TextFileWithRecurrence); j++ { // Replace recurrence by a space
			if TextFileWithRecurrence[i] == TextFileWithRecurrence[j] {
				counter++
			}
			if counter >= 2 {
				TextFileWithRecurrence[j] = " "
			}
		}
	}
	for _, verified := range TextFileWithRecurrence { // Select only word and not a space
		if verified != " " || len(verified) > 1 {
			perfectTextFile = append(perfectTextFile, verified)
		}
	}
	return perfectTextFile
}
