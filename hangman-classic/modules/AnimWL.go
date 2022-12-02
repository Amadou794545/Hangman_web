package modules

import (
	"bufio"
	"fmt"
	"github.com/hajimehoshi/go-mp3"
	"github.com/hajimehoshi/oto/v2"
	"log"
	"os"
	"time"
)

func AnimWL(result string) {
	var lineTxt int
	var charToPrint string
	var counterEveryPicture int
	var timeToPrint int
	var counterForSound int
	if result == "lose" {
		Clean()
		Hide()
		go MusicLose()
		readFile, err := os.Open("Animation/Scene/SceneFinalLose.txt")
		if err != nil {
			log.Fatal(err)
		}
		fileScanner := bufio.NewScanner(readFile)
		fileScanner.Split(bufio.ScanLines)
		for fileScanner.Scan() {
			counterEveryPicture -= 1
			if fileScanner.Text() == "1" {
				lineTxt = 23
				counterEveryPicture = 1
				if timeToPrint != 0 {
					time.Sleep(200000 * time.Microsecond)
				}
				fmt.Printf("\033[198;24H")
				os.Stdout.Write([]byte("\033[1;0H"))
				os.Stdout.Write([]byte(charToPrint))
				charToPrint = ""
				timeToPrint += 1
			}
			if lineTxt > 0 && counterEveryPicture <= 0 {
				charToPrint = charToPrint + fileScanner.Text() + "\n"
				lineTxt -= 1
			} // Not color
		}
		Show()
		readFile.Close()
	} else {
		Clean()
		Hide()
		readFile, err := os.Open("Animation/Scene/SceneFinalVictoire.txt")
		if err != nil {
			log.Fatal(err)
		}
		fileScanner := bufio.NewScanner(readFile)
		fileScanner.Split(bufio.ScanLines)
		for fileScanner.Scan() {
			counterEveryPicture -= 1
			if fileScanner.Text() == "1" {
				lineTxt = 24
				counterEveryPicture = 1
				counterForSound += 1
				if counterForSound == 30 {
					go Thanks()
				}
				if counterForSound == 57 {
					go MusicWin()
				}
				if timeToPrint != 0 {
					time.Sleep(150 * time.Millisecond)
				}
				fmt.Printf("\033[198;24H")
				os.Stdout.Write([]byte("\033[1;0H"))
				os.Stdout.Write([]byte(charToPrint))
				charToPrint = ""
				timeToPrint += 1
			}
			if lineTxt > 0 && counterEveryPicture <= 0 {
				charToPrint = charToPrint + fileScanner.Text() + "\n"
				lineTxt -= 1
			}
		}
		Show()
		readFile.Close()
	}
}

func MusicWin() error {
	f, err := os.Open("Animation/Music/Cantina_Winner_Soong.mp3")
	if err != nil {
		return err
	}
	defer f.Close()

	d, err := mp3.NewDecoder(f)
	if err != nil {
		return err
	}

	c, ready, err := oto.NewContext(d.SampleRate(), 2, 2)
	if err != nil {
		return err
	}
	<-ready

	p := c.NewPlayer(d)
	defer p.Close()
	p.Play()
	for {
		time.Sleep(time.Second)
		if !p.IsPlaying() {
			break
		}
	}

	return nil
} // Music Winner
func MusicLose() error {
	f, err := os.Open("Animation/Music/generique.mp3")
	if err != nil {
		return err
	}
	defer f.Close()

	d, err := mp3.NewDecoder(f)
	if err != nil {
		return err
	}

	c, ready, err := oto.NewContext(d.SampleRate(), 2, 2)
	if err != nil {
		return err
	}
	<-ready

	p := c.NewPlayer(d)
	defer p.Close()
	p.Play()
	for {
		time.Sleep(time.Second)
		if !p.IsPlaying() {
			break
		}
	}

	return nil
} // Music Loser
func Thanks() error {
	f, err := os.Open("Animation/Music/ThanksALotToRescueMeFromTheJawasGame.mp3")
	if err != nil {
		return err
	}
	defer f.Close()

	d, err := mp3.NewDecoder(f)
	if err != nil {
		return err
	}

	c, ready, err := oto.NewContext(d.SampleRate(), 2, 2)
	if err != nil {
		return err
	}
	<-ready

	p := c.NewPlayer(d)
	defer p.Close()
	p.Play()
	for {
		time.Sleep(time.Second)
		if !p.IsPlaying() {
			break
		}
	}

	return nil
} // Music Thanks C3PO
