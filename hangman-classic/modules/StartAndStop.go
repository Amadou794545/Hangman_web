package modules

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/hajimehoshi/go-mp3"
	"github.com/hajimehoshi/oto/v2"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type HangmanData struct {
	Version    string   `json:"version"`
	Difficulty string   `json:"difficulty"`
	Live       int      `json:"live"`
	Word       string   `json:"word"`
	Result     []string `json:"result"`
	UsedLW     []string `json:"used_lw"`
	Index      []int    `json:"index"`
}

func StartWithSave(FileName string) (HangmanData, error) {
	// Taking the data from the save
	mBytes, errReadFile := os.ReadFile(FileName)

	// Decoding the data and return to main
	dec := json.NewDecoder(strings.NewReader(string(mBytes)))
	var data HangmanData
	if err := dec.Decode(&data); err == io.EOF {
	} else if err != nil {
		fmt.Printf("We failed to take back the data of the save...\nErreur : %d", errReadFile)
	}
	return data, errReadFile
}

func StopAndSave(gameData HangmanData) {
	// Functions for animation and sound or pictures
	StormTrooper()
	S_StormTrooper()
	// Put values in structure variables + enccoding in json.Marshal
	data := HangmanData{gameData.Version, gameData.Difficulty, gameData.Live, gameData.Word, gameData.Result, gameData.UsedLW, gameData.Index}

	// Buff creation to encode
	buf := new(bytes.Buffer)
	encoder := json.NewEncoder(buf)
	encoder.Encode(data)

	FileName, errOpenFile := createFile()
	if errOpenFile == nil {
		createFile()
	}

	// VerifyRecurrenceWord creation
	file, errCreateJson := os.Create(FileName)
	if errCreateJson != nil {
		log.Fatal(errCreateJson)
	}
	defer file.Close()

	// Data copy of the structure encoded in json
	io.Copy(file, buf)
	os.Exit(1)
}

func createFile() (string, error) {
	// Variables
	var FileName string
	var errOpenFile error

	// Condition to choose a possible name file
	name := "Save"
	counter := 0
	extension := ".json"
	CreatedFile := false

	for CreatedFile != true {
		FileName = name + strconv.Itoa(counter) + extension
		_, errOpenFile = os.Open(FileName)

		if errOpenFile != nil {
			CreatedFile = true
		} else {
			counter++
		}
	}
	return FileName, errOpenFile
}

func StormTrooper() {
	var test string
	Clean()
	readFile, err := os.Open("Picture/StormTrooper.txt") // Take the data from the txt file
	if err != nil {
		log.Fatal(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		test = test + fileScanner.Text() + "\n"
	}
	fmt.Printf("\033[198;24H")
	os.Stdout.Write([]byte("\033[1;0H"))
	os.Stdout.Write([]byte(test))

}
func S_StormTrooper() error {
	f, err := os.Open("Picture/ISaidComeBackNow.mp3")
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
} // Sound of stormtrooper
func Pilote() {
	var test string
	Clean()
	readFile, err := os.Open("Picture/Pilote.txt") // Take the data from the txt file
	if err != nil {
		log.Fatal(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		test = test + fileScanner.Text() + "\n"
	}
	fmt.Printf("\033[198;24H")
	os.Stdout.Write([]byte("\033[1;0H"))
	os.Stdout.Write([]byte(test))

} // Print Pilot X-wing
func S_Pilot() error {
	f, err := os.Open("Picture/Pilot.mp3")
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
} // Sounf Pilot voice
