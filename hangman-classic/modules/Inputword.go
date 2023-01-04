package modules

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func InputWord(file string) string { // Take randomly a word in file given .
	var words []string
	var word string
	switch file {
	case "words.txt":
		fmt.Printf("BOOM")
		file = "../hangman-web/Templates/words.txt"
	case "words2.txt":
		fmt.Printf("BOOM")
		file = "../hangman-web/Templates/words2.txt"
	case "words3.txt":
		fmt.Printf("BOOM")
		file = "../hangman-web/Templates/words3.txt"
	}
	// Open file
	openFile, err := os.Open(file)
	if err != nil {
	}
	defer openFile.Close()
	scanner := bufio.NewScanner(openFile)
	scanner.Split(bufio.ScanWords)

	// Put words in array
	for scanner.Scan() {
		words = append(words, []string{scanner.Text()}...)
	}

	// Select randomly a word in array
	rand.Seed(time.Now().Unix())
	b := rand.Perm(len(words))
	word = words[b[0]]

	// Return word selected
	return word
}
