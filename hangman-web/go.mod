module main

go 1.19

replace Hangman => ../hangman-classic

require Hangman v0.0.0-00010101000000-000000000000

require (
	github.com/hajimehoshi/go-mp3 v0.3.3 // indirect
	github.com/hajimehoshi/oto/v2 v2.3.1 // indirect
	golang.org/x/sys v0.0.0-20220712014510-0a85c31ab51e // indirect
)
