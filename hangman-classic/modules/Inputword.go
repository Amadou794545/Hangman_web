package modules

import (
	"bufio"
	"math/rand"
	"os"
	"time"
)

func InputWord() string { // Take randomly a word in file given .

	var words []string
	var word string

	file := Difficulty()

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
