package modules

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func AsciiArt(word []string) {
	readFile, err := os.Open("standard1.txt") // Take the data from the txt file
	if err != nil {
		log.Fatal(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var textFile []string
	for fileScanner.Scan() {
		textFile = append(textFile, fileScanner.Text()) // Put the data in a table
	}
	var line1 string
	var line2 string
	var line3 string
	var line4 string
	var line5 string
	var line6 string
	var line7 string
	var line8 string
	for i := 0; i < len(word); i++ {
		for j := 0; j < len(textFile); j++ {
			// If the letter is equal to the letter in the data, the variable line take data
			if word[i] == textFile[j] {
				line1 += textFile[j+1]
				line2 += textFile[j+2]
				line3 += textFile[j+3]
				line4 += textFile[j+4]
				line5 += textFile[j+5]
				line6 += textFile[j+6]
				line7 += textFile[j+7]
				line8 += textFile[j+8]
				break
			}
		}
	}
	fmt.Println(line1)
	fmt.Println(line2)
	fmt.Println(line3)
	fmt.Println(line4)
	fmt.Println(line5)
	fmt.Println(line6)
	fmt.Println(line7)
	fmt.Println(line8)
}
