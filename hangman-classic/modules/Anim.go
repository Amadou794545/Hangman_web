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

func Anim(Live int) {
	Hide()
	// Every Scene to play for every Live
	if Live == 9 {
		var lineText int
		var charToPrint string
		var counterEveryPicture int
		var timeToPrint int
		Clean()
		go S_Scene1()
		readFile, err := os.Open("Animation/Scene/Scene1.txt")
		if err != nil {
			log.Fatal(err)
		}
		fileScanner := bufio.NewScanner(readFile)
		fileScanner.Split(bufio.ScanLines)
		for fileScanner.Scan() {
			counterEveryPicture -= 1
			if fileScanner.Text() == "1" {
				lineText = 24 //
				counterEveryPicture = 1
				if timeToPrint != 0 {
					time.Sleep(50 * time.Millisecond)
				}
				fmt.Printf("\033[198;24H")
				os.Stdout.Write([]byte("\033[1;0H"))
				os.Stdout.Write([]byte(charToPrint))
				charToPrint = ""
				timeToPrint += 1
			}
			if lineText > 0 && counterEveryPicture <= 0 {
				charToPrint = charToPrint + fileScanner.Text() + "\n"
				lineText -= 1
			} // Not color
		}
	}
	if Live == 8 {
		var lineText int
		var charToPrint string
		var counterEveryPicture int
		var timeToPrint int
		Clean()
		go S_Scene2()
		readFile, err := os.Open("Animation/Scene/Scene2.txt")
		if err != nil {
			log.Fatal(err)
		}
		fileScanner := bufio.NewScanner(readFile)
		fileScanner.Split(bufio.ScanLines)
		for fileScanner.Scan() {
			counterEveryPicture -= 1
			if fileScanner.Text() == "1" {
				lineText = 24
				counterEveryPicture = 1
				if timeToPrint != 0 {
					time.Sleep(50 * time.Millisecond)
				}
				fmt.Printf("\033[198;24H")
				os.Stdout.Write([]byte("\033[1;0H"))
				os.Stdout.Write([]byte(charToPrint))
				charToPrint = ""
				timeToPrint += 1
			}
			if lineText > 0 && counterEveryPicture <= 0 {
				charToPrint = charToPrint + fileScanner.Text() + "\n"
				lineText -= 1
			} // Not color
		}
	}
	if Live == 7 {
		var lineText int
		var charToPrint string
		var counterEveryPicture int
		var timeToPrint int
		Clean()
		go S_Scene3()
		readFile, err := os.Open("Animation/Scene/Scene3.txt")
		if err != nil {
			log.Fatal(err)
		}
		fileScanner := bufio.NewScanner(readFile)
		fileScanner.Split(bufio.ScanLines)
		for fileScanner.Scan() {
			counterEveryPicture -= 1
			if fileScanner.Text() == "1" {
				lineText = 24
				counterEveryPicture = 1
				if timeToPrint != 0 {
					time.Sleep(50 * time.Millisecond)
				}
				fmt.Printf("\033[198;24H")
				os.Stdout.Write([]byte("\033[1;0H"))
				os.Stdout.Write([]byte(charToPrint))
				charToPrint = ""
				timeToPrint += 1
			}
			if lineText > 0 && counterEveryPicture <= 0 {
				charToPrint = charToPrint + fileScanner.Text() + "\n"
				lineText -= 1
			} // Not color
		}
	}
	if Live == 6 {
		var lineText int
		var charToPrint string
		var counterEveryPicture int
		var timeToPrint int
		Clean()
		go S_Scene4()
		readFile, err := os.Open("Animation/Scene/Scene4.txt")
		if err != nil {
			log.Fatal(err)
		}
		fileScanner := bufio.NewScanner(readFile)
		fileScanner.Split(bufio.ScanLines)
		for fileScanner.Scan() {
			counterEveryPicture -= 1
			if fileScanner.Text() == "1" {
				lineText = 24
				counterEveryPicture = 1
				if timeToPrint != 0 {
					time.Sleep(50 * time.Millisecond)
				}
				fmt.Printf("\033[198;24H")
				os.Stdout.Write([]byte("\033[1;0H"))
				os.Stdout.Write([]byte(charToPrint))
				charToPrint = ""
				timeToPrint += 1
			}
			if lineText > 0 && counterEveryPicture <= 0 {
				charToPrint = charToPrint + fileScanner.Text() + "\n"
				lineText -= 1
			} // Not color
		}
	}
	if Live == 5 {
		var lineText int
		var charToPrint string
		var counterEveryPicture int
		var timeToPrint int
		Clean()
		go S_Scene5()
		readFile, err := os.Open("Animation/Scene/Scene5.txt")
		if err != nil {
			log.Fatal(err)
		}
		fileScanner := bufio.NewScanner(readFile)
		fileScanner.Split(bufio.ScanLines)
		for fileScanner.Scan() {
			counterEveryPicture -= 1
			if fileScanner.Text() == "1" {
				lineText = 24
				counterEveryPicture = 1
				if timeToPrint != 0 {
					time.Sleep(50 * time.Millisecond)
				}
				fmt.Printf("\033[198;24H")
				os.Stdout.Write([]byte("\033[1;0H"))
				os.Stdout.Write([]byte(charToPrint))
				charToPrint = ""
				timeToPrint += 1
			}
			if lineText > 0 && counterEveryPicture <= 0 {
				charToPrint = charToPrint + fileScanner.Text() + "\n"
				lineText -= 1
			} // Not color
		}
	}
	if Live == 4 {
		var lineText int
		var charToPrint string
		var counterEveryPicture int
		var timeToPrint int
		Clean()
		go S_Scene6()
		readFile, err := os.Open("Animation/Scene/Scene6.txt")
		if err != nil {
			log.Fatal(err)
		}
		fileScanner := bufio.NewScanner(readFile)
		fileScanner.Split(bufio.ScanLines)
		for fileScanner.Scan() {
			counterEveryPicture -= 1
			if fileScanner.Text() == "1" {
				lineText = 24
				counterEveryPicture = 1
				if timeToPrint != 0 {
					time.Sleep(50 * time.Millisecond)
				}
				fmt.Printf("\033[198;24H")
				os.Stdout.Write([]byte("\033[1;0H"))
				os.Stdout.Write([]byte(charToPrint))
				charToPrint = ""
				timeToPrint += 1
			}
			if lineText > 0 && counterEveryPicture <= 0 {
				charToPrint = charToPrint + fileScanner.Text() + "\n"
				lineText -= 1
			} // Not color
		}
	}
	if Live == 3 {
		var lineText int
		var charToPrint string
		var counterEveryPicture int
		var timeToPrint int
		Clean()
		go S_Scene7()
		readFile, err := os.Open("Animation/Scene/Scene7.txt")
		if err != nil {
			log.Fatal(err)
		}
		fileScanner := bufio.NewScanner(readFile)
		fileScanner.Split(bufio.ScanLines)
		for fileScanner.Scan() {
			counterEveryPicture -= 1
			if fileScanner.Text() == "1" {
				lineText = 24
				counterEveryPicture = 1
				if timeToPrint != 0 {
					time.Sleep(50 * time.Millisecond)
				}
				fmt.Printf("\033[198;24H")
				os.Stdout.Write([]byte("\033[1;0H"))
				os.Stdout.Write([]byte(charToPrint))
				charToPrint = ""
				timeToPrint += 1
			}
			if lineText > 0 && counterEveryPicture <= 0 {
				charToPrint = charToPrint + fileScanner.Text() + "\n"
				lineText -= 1
			} // Not color
		}
	}
	if Live == 2 {
		var lineText int
		var charToPrint string
		var counterEveryPicture int
		var timeToPrint int
		Clean()
		go S_Scene8()
		readFile, err := os.Open("Animation/Scene/Scene8.txt")
		if err != nil {
			log.Fatal(err)
		}
		fileScanner := bufio.NewScanner(readFile)
		fileScanner.Split(bufio.ScanLines)
		for fileScanner.Scan() {
			counterEveryPicture -= 1
			if fileScanner.Text() == "1" {
				lineText = 24
				counterEveryPicture = 1
				if timeToPrint != 0 {
					time.Sleep(50 * time.Millisecond)
				}
				fmt.Printf("\033[198;24H")
				os.Stdout.Write([]byte("\033[1;0H"))
				os.Stdout.Write([]byte(charToPrint))
				charToPrint = ""
				timeToPrint += 1
			}
			if lineText > 0 && counterEveryPicture <= 0 {
				charToPrint = charToPrint + fileScanner.Text() + "\n"
				lineText -= 1
			} // Not color
		}
	}
	if Live == 1 {
		var lineText int
		var charToPrint string
		var counterEveryPicture int
		var timeToPrint int
		Clean()
		go S_Scene9()
		readFile, err := os.Open("Animation/Scene/Scene9.txt")
		if err != nil {
			log.Fatal(err)
		}
		fileScanner := bufio.NewScanner(readFile)
		fileScanner.Split(bufio.ScanLines)
		for fileScanner.Scan() {
			counterEveryPicture -= 1
			if fileScanner.Text() == "1" {
				lineText = 24
				counterEveryPicture = 1
				if timeToPrint != 0 {
					time.Sleep(150 * time.Millisecond)
				}
				fmt.Printf("\033[198;24H")
				os.Stdout.Write([]byte("\033[1;0H"))
				os.Stdout.Write([]byte(charToPrint))
				charToPrint = ""
				timeToPrint += 1
			}
			if lineText > 0 && counterEveryPicture <= 0 {
				charToPrint = charToPrint + fileScanner.Text() + "\n"
				lineText -= 1
			} // Not color
		}
	}
	if Live == 0 {
		var lineText int
		var charToPrint string
		var counterEveryPicture int
		var timeToPrint int
		Clean()
		go S_Scene10()
		readFile, err := os.Open("Animation/Scene/Scene10.txt")
		if err != nil {
			log.Fatal(err)
		}
		fileScanner := bufio.NewScanner(readFile)
		fileScanner.Split(bufio.ScanLines)
		for fileScanner.Scan() {
			counterEveryPicture -= 1
			if fileScanner.Text() == "1" {
				lineText = 24
				counterEveryPicture = 1
				if timeToPrint != 0 {
					time.Sleep(50 * time.Millisecond)
				}
				fmt.Printf("\033[198;24H")
				os.Stdout.Write([]byte("\033[1;0H"))
				os.Stdout.Write([]byte(charToPrint))
				charToPrint = ""
				timeToPrint += 1
			}
			if lineText > 0 && counterEveryPicture <= 0 {
				charToPrint = charToPrint + fileScanner.Text() + "\n"
				lineText -= 1
			} // Not color
		}
	}
	Show()
}

// Musique for every Scene
func S_Scene1() error {
	f, err := os.Open("Animation/Music/MusicScene1.mp3")
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
func S_Scene2() error {
	f, err := os.Open("Animation/Music/MusicScene2.mp3")
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
func S_Scene3() error {
	f, err := os.Open("Animation/Music/MusicScene3.mp3")
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
func S_Scene4() error {
	f, err := os.Open("Animation/Music/MusicScene4.mp3")
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
func S_Scene5() error {
	f, err := os.Open("Animation/Music/MusicScene5.mp3")
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
func S_Scene6() error {
	f, err := os.Open("Animation/Music/MusicScene6.mp3")
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
func S_Scene7() error {
	f, err := os.Open("Animation/Music/MusicScene7.mp3")
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
func S_Scene8() error {
	f, err := os.Open("Animation/Music/MusicScene8.mp3")
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
func S_Scene9() error {
	f, err := os.Open("Animation/Music/Scene9Music.mp3")
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
func S_Scene10() error {
	f, err := os.Open("Animation/Music/MusicScene10.mp3")
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
