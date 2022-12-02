package modules

import (
	"bufio"
	"fmt"
	"github.com/hajimehoshi/go-mp3"
	"github.com/hajimehoshi/oto/v2"
	"io"
	"log"
	"os"
	"time"
)

func AnimIntro() string {
	var goodInput bool
	var userInputDifficulty string
	var difficulty string
	Hide()
	Clean()
	go MusicIntro()
	time.Sleep(20 * time.Millisecond)
	Intro()
	Histoire()
	Rules()
	Show()

	// Ask difficulty and retry if userInput is not correct
	for goodInput == false {
		fmt.Scan(&userInputDifficulty)
		if userInputDifficulty == "exit" { // Command to leave the program instantly
			ExitCommand()
		}

		if userInputDifficulty != "easy" && userInputDifficulty != "medium" && userInputDifficulty != "hard" {
			fmt.Println("You can only chose the difficulty : easy | medium | hard ")
		} else {
			difficulty = userInputDifficulty
			goodInput = true
		}
	}
	return difficulty
}
func Hide() {
	type Writer interface {
		io.Writer
		Fd() uintptr
	}
	var target Writer = os.Stdout
	fmt.Fprintf(target, "\x1b[?25l")
} // Hide cursor
func Show() {
	type Writer interface {
		io.Writer
		Fd() uintptr
	}
	var target Writer = os.Stdout
	fmt.Fprint(target, "\x1b[?25h")
} // Show cursor

// Voice pilot x-wing
func Voice1() error {
	f, err := os.Open("Animation/Music/Voice1.mp3")
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
}
func Voice2() error {
	f, err := os.Open("Animation/Music/Voice2.mp3")
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
}
func Voice3() error {
	f, err := os.Open("Animation/Music/Voice3.mp3")
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
}
func Voice4() error {
	f, err := os.Open("Animation/Music/Voice4.mp3")
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
}
func Voice5() error {
	f, err := os.Open("Animation/Music/Voice5.mp3")
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
}

func ArrivalXwing() error {
	f, err := os.Open("Animation/Music/ArrivalXwing.mp3")
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
} // Audio x-wing
func XWingSongExt() error {
	f, err := os.Open("Animation/Music/XWingSongExt.mp3")
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
} // Audio x-wing motor
func R2D2XWing() error {
	f, err := os.Open("Animation/Music/R2D2XWing.mp3")
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
} // R2D2 audio
func Hyperspace() error {
	f, err := os.Open("Animation/Music/Hyperspace.mp3")
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
} // Hyperspace audio
func MusicIntro() error {
	f, err := os.Open("Animation/Music/starwarsintro.mp3")
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
} // Music intro

// Anim
func Intro() {

	var Yellow = "\033[33m"
	var Reset = "\u001b[37;1m"
	fmt.Printf("        .                                                                             .              .       .                    .      .                         .                           .                  \n")
	fmt.Printf("                                            .                    .        .                                          .                                                                                            \n")
	fmt.Printf("                      .                                                        .        .                                           .                   .                      .                                  \n")
	fmt.Printf("                                                                                                   .          .              .             .                                                   .         .        \n")
	fmt.Printf("                                    .                .                          " + Yellow + "_________________      ____         __________" + Reset + "                                                                      \n")
	fmt.Printf("         .                                                        .       .    " + Yellow + "/                 |    /    \\    " + Reset + "." + Yellow + "  |          \\" + Reset + "                                    .                                \n")
	fmt.Printf("                                            .                                 " + Yellow + "/    ______   _____| " + Reset + "." + Yellow + " /      \\      |    ___    |" + Reset + "     .                                       .                      \n")
	fmt.Printf("                        .                                                     " + Yellow + "\\    \\    |   |       /   /\\   \\     |   |___>   |" + Reset + "                       .                                             .\n")
	fmt.Printf("                                                                            .  " + Yellow + "\\    \\   |   |      /   /__\\   \\  " + Reset + "." + Yellow + " |         _/" + Reset + "               .                                         .           \n")
	fmt.Printf("         .                                                        .     " + Yellow + "________>    |  |   | " + Reset + "." + Yellow + "   /            \\   |   |\\    \\_______ " + Reset + "   .                                  .                       \n")
	fmt.Printf("                                                      .                " + Yellow + "|            /   |   |    /    ______    \\  |   | \\           |" + Reset + "                           .                                  \n")
	fmt.Printf("                                                                       " + Yellow + "|___________/    |___|   /____/      \\____\\ |___|  \\__________|" + Reset + "                                                       .      \n")
	fmt.Printf("                   .                    .                                " + Yellow + "____    __  " + Reset + "." + Yellow + " _____   ____      " + Reset + "." + Yellow + "  __________   " + Reset + "." + Yellow + "  _________ " + Reset + "               .                      .                       \n")
	fmt.Printf("        .                  .                                            " + Yellow + "\\    \\  /  \\  /    /  /    \\       |          \\    /         |" + Reset + "      .                                                       \n")
	fmt.Printf("                                                                         " + Yellow + "\\    \\/    \\/    /  /      \\      |    ___    |  /    ______|" + Reset + "  .                                                           \n")
	fmt.Printf("                                                     .                    " + Yellow + "\\              /  /   /\\   \\ " + Reset + "." + Yellow + "   |   |___>   |  \\    \\" + Reset + "     |                     .                           .            \n")
	fmt.Printf("                                        .                           .      " + Yellow + "\\            /  /   /__\\   \\    |         _/" + Reset + "." + Yellow + "   \\    \\" + Reset + "    O                                 .                            \n")
	fmt.Printf("              .             .                                               " + Yellow + "\\    /\\    /  /            \\   |   |\\    \\______>    |" + Reset + "  /|\\                                                             \n")
	fmt.Printf("                                                     .                       " + Yellow + "\\  /  \\  /  /    ______    \\  |   | \\              /" + Reset + "   / \\    .             .                                          . \n")
	fmt.Printf("                                     .                            .       .   " + Yellow + "\\/    \\/  /____/      \\____\\ |___|  \\____________/" + Reset + "                                       .                            \n")
	fmt.Printf("        .                                                                                       .                                        .            .                                                           \n")
	fmt.Printf("                                                     .                                                                                                             .         .                                 .  \n")
	fmt.Printf("                     .                .                        .                 .                                   .            .                        .                           .                          \n")
	time.Sleep(3 * time.Second)
} // Print in color star wars
func Histoire() {
	var lineText int
	var charToPrint string
	var counterEveryPicture int
	var timeToPrint int
	var Yellow = "\033[33m"
	var Reset = "\u001b[37;1m"
	readFile, err := os.Open("Animation/Scene/Histoire.txt")
	if err != nil {
		log.Fatal(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		counterEveryPicture -= 1
		if fileScanner.Text() == "1" {
			lineText = 23
			counterEveryPicture = 1
			if timeToPrint != 0 {
				time.Sleep(1600 * time.Millisecond)
			}
			fmt.Printf("\033[198;24H")
			os.Stdout.Write([]byte("\033[1;0H"))
			os.Stdout.Write([]byte(charToPrint))
			charToPrint = ""
			timeToPrint += 1
		}
		if lineText > 0 && counterEveryPicture <= 0 {
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

} // History of the game
func Rules() {
	var linetext int
	var charToPrint string
	var counterEveryPicture int
	var timeToPrint int
	var CounterSong int
	Clean()
	readFile, err := os.Open("Animation/Scene/Regles.txt")
	if err != nil {
		log.Fatal(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		counterEveryPicture -= 1
		if fileScanner.Text() == "1" {
			linetext = 23
			counterEveryPicture = 1
			CounterSong += 1
			if CounterSong == 1 {
				go XWingSongExt()
				go ArrivalXwing()
			}
			if CounterSong == 149 {
				go Voice1()
			}
			if CounterSong == 200 {
				go R2D2XWing()
			}
			if CounterSong == 223 {
				go Voice2()
			}
			if CounterSong == 324 {
				go Voice3()
			}
			if CounterSong == 412 {
				go Voice4()
			}
			if CounterSong == 539 {
				go Voice5()
			}
			if CounterSong == 625 {
				go Hyperspace()
			}
			if timeToPrint != 0 {
				time.Sleep(50 * time.Millisecond) //1400 Mili for rules
			}
			fmt.Printf("\033[198;24H")
			os.Stdout.Write([]byte("\033[1;0H"))
			os.Stdout.Write([]byte(charToPrint))
			charToPrint = ""
			timeToPrint += 1
		}
		if linetext > 0 && counterEveryPicture <= 0 {
			charToPrint = charToPrint + fileScanner.Text() + "\n"
			linetext -= 1
		} // Not color
	}
} // Scene with rules explain
