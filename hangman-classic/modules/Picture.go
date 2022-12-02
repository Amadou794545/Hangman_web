package modules

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

func Picture(live int) {
	var textPicture string
	Clean()
	// Every Picture to print for every live

	if live == 9 {
		readFile, err := os.Open("Picture/Scene1Picture.txt") // Take the data from the txt file
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
	if live == 8 {
		readFile, err := os.Open("Picture/Scene2Picture.txt") // Take the data from the txt file
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
	if live == 7 {
		readFile, err := os.Open("Picture/Scene3Picture.txt") // Take the data from the txt file
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
	if live == 6 {
		readFile, err := os.Open("Picture/Scene4Picture.txt") // Take the data from the txt file
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
	if live == 5 {
		readFile, err := os.Open("Picture/Scene5Picture.txt") // Take the data from the txt file
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
	if live == 4 {
		readFile, err := os.Open("Picture/Scene6Picture.txt") // Take the data from the txt file
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
	if live == 3 {
		readFile, err := os.Open("Picture/Scene7Picture.txt") // Take the data from the txt file
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
	if live == 2 {
		readFile, err := os.Open("Picture/Scene8Picture.txt") // Take the data from the txt file
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
	if live == 1 {
		readFile, err := os.Open("Picture/Scene9Picture.txt") // Take the data from the txt file
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
	if live == 0 {
		readFile, err := os.Open("Picture/Scene10Picture.txt") // Take the data from the txt file
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
		time.Sleep(2 * time.Second)
		Clean()
		readFile, err = os.Open("Picture/Scene10Picture2.txt") // Take the data from the txt file
		if err != nil {
			log.Fatal(err)
		}
		fileScanner = bufio.NewScanner(readFile)
		fileScanner.Split(bufio.ScanLines)
		for fileScanner.Scan() {
			textPicture = textPicture + fileScanner.Text() + "\n"
		}
		fmt.Printf("\033[198;24H")
		os.Stdout.Write([]byte("\033[1;0H"))
		os.Stdout.Write([]byte(textPicture))
		time.Sleep(2 * time.Second)
	}
}
