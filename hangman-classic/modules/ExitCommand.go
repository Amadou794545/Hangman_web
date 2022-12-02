package modules

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

func ExitCommand() { // Print Yoda when the user wants to exit
	var green = "\033[32m"
	var reset = "\033[0m"
	var yellow = "\033[33m"
	var brown = "\033[90m"
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()

	fmt.Println(green + "⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀  ⣀⣤⣤⠤⠐⠂⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀")
	time.Sleep(80 * time.Millisecond)
	fmt.Println("⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⡌⡦⠊⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀")
	time.Sleep(80 * time.Millisecond)
	fmt.Println("⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⡀⣼⡊⢀⠔⠀⠀⣄⠤⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣀⣤⣤⣄⣀⠀")
	time.Sleep(80 * time.Millisecond)
	fmt.Println("⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢠⣶⠃⠉⠡⡠⠤⠊⠀⠠⣀⣀⡠⠔⠒⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣠⣾⣿⢟⠿⠛⠛⠁")
	time.Sleep(80 * time.Millisecond)
	fmt.Println("⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣼⡇⠀⠀⠀⠀⠑⠶⠖⠊⠁⠀⠀⠀⡀⠀⠀⠀⢀⣠⣤⣤⡀⠀⠀⠀⠀⠀⢀⣠⣤⣶⣿⣿⠟⡱⠁⠀⠀⠀⠀")
	time.Sleep(80 * time.Millisecond)
	fmt.Println("⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢰⣾⣿⡇⠀⢀⡠⠀⠀⠀⠈⠑⢦⣄⣀⣀⣽⣦⣤⣾⣿⠿⠿⠿⣿⡆⠀⠀⢀⠺⣿⣿⣿⣿⡿⠁⡰⠁⠀⠀⠀⠀⠀")
	time.Sleep(80 * time.Millisecond)
	fmt.Println("⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣾⣿⣿⣧⣠⠊⣠⣶⣾⣿⣿⣶⣶⣿⣿⠿⠛⢿⣿⣫⢕⡠⢥⣈⠀⠙⠀⠰⣷⣿⣿⣿⡿⠋⢀⠜⠁⠀⠀⠀⠀⠀⠀")
	time.Sleep(80 * time.Millisecond)
	fmt.Println("⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠠⢿⣿⣿⣿⣿⣰⣿⣿⠿⣛⡛⢛⣿⣿⣟⢅⠀⠀⢿⣿⠕⢺⣿⡇⠩⠓⠂⢀⠛⠛⠋⢁⣠⠞⠁⠀⠀⠀⠀⠀⠀⠀⠀")
	time.Sleep(80 * time.Millisecond)
	fmt.Println("⠘⢶⡶⢶⣶⣦⣤⣤⣤⣤⣤⣀⣀⣀⣀⡀⠀⠘⣿⣿⣿⠟⠁⡡⣒⣬⢭⢠⠝⢿⡡⠂⠀⠈⠻⣯⣖⣒⣺⡭⠂⢀⠈⣶⣶⣾⠟⠁⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀")
	time.Sleep(80 * time.Millisecond)
	fmt.Println("⠀⠀⠙⠳⣌⡛⢿⣿⣿⣿⣿⣿⣿⣿⣿⣻⣵⣨⣿⣿⡏⢀⠪⠎⠙⠿⣋⠴⡃⢸⣷⣤⣶⡾⠋⠈⠻⣶⣶⣶⣷⣶⣷⣿⣟⠁⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀")
	time.Sleep(80 * time.Millisecond)
	fmt.Println("⠀⠀⠀⠀⠈⠛⢦⣌⡙⠛⠿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡀⠀⠀⠩⠭⡭⠴⠊⢀⠀⠀⠀⠀⠀⠀⠀⠀⠈⣿⣿⣿⣿⣿⡇⠁⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀")
	time.Sleep(80 * time.Millisecond)
	fmt.Println("⠀⠀⠀⠀⠀⠀⠀⠈⠙⠓⠦⣄⡉⠛⠛⠻⢿⣿⣿⣿⣷⡀⠀⠀⠀⠀⢀⣰⠋⠀⠀⠀⠀⠀⣀⣰⠤⣳⣿⣿⣿⣿⣟⠑⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀")
	time.Sleep(80 * time.Millisecond)
	fmt.Println("⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠉⠓⠒⠒⠶⢺⣿⣿⣿⣿⣦⣄⣀⣴⣿⣯⣤⣔⠒⠚⣒⣉⣉⣴⣾⣿⣿⣿⣿⣿⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀")
	time.Sleep(80 * time.Millisecond)
	fmt.Println("⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠛⠹⢿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡇⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀")
	time.Sleep(80 * time.Millisecond)
	fmt.Println("⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠙⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣭⣉⣉⣤⣿⣿⣿⣿⣿⣿⡿" + reset + "⢀⡀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀")
	time.Sleep(80 * time.Millisecond)
	fmt.Println("⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣠" + green + "⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡿⠟⡁" + reset + "⡆⠙⢶⣀⠀⢀⣀⡀⠀⠀⠀⠀⠀⠀⠀⠀")
	time.Sleep(80 * time.Millisecond)
	fmt.Println("⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⣀⣴⣶⣾⣿⣟" + green + "⢿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⠿⢛⣩" + brown + "⣴⣿⠇" + reset + "⡇⠸⡆⠙⢷⣄⠻⣿⣦⡄⠀⠀⠀⠀⠀⠀")
	time.Sleep(80 * time.Millisecond)
	fmt.Println("⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⣼⣿⣿⣿⣿⣿⣿⣿⣎" + green + "⢻⣿⣿⣿⣿⣿⣿⣿⣭⣭⣭⣵" + brown + "⣶⣾⣿⣿⣿⠟" + reset + "⢰⢣⠀⠈⠀⠀⠙⢷⡎⠙⣿⣦⠀⠀⠀⠀⠀")
	time.Sleep(80 * time.Millisecond)
	fmt.Println("⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⣼⣿⣿⣿⣿⣿⣿⣿⣿⡟⣿⡆" + brown + "⢻⣿⣿⣿⣿⣿⣿⣿⣿⣿⠿⠿⠟⠛⠋⠁⢀" + reset + "⠇⢸⡇⠀⠀⠀⠀⠈⠁⠀⢸⣿⡆⠀⠀⠀⠀")
	time.Sleep(80 * time.Millisecond)
	fmt.Println("⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢠⣾⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡜⡿" + brown + "⡘⣿⣿⣿⣿⣿⣶⣶⣤⣤⣤⣤⣤⣤⣤⣴" + reset + "⡎⠖⢹⡇⠀⠀⠀⠀⠀⠀⠀⠀⣿⣷⡄⠀⠀⠀")
	time.Sleep(80 * time.Millisecond)
	fmt.Println("⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣦⡀" + brown + "⠘⢿⣿⣿⣿⣿⣿⣿⣿⣿⠿⠿⠛⠋" + reset + "⡟⠀⠀⣸⣷⣀⣤⣀⣀⣀⣤⣤⣾⣿⣿⣿⠀⠀⠀")
	time.Sleep(80 * time.Millisecond)
	fmt.Println("⠀⠀⠀⠀⠀⠀⠀⠀⠀⣸⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣭⣓⡲⠬" + reset + "⢭⣙⡛⠿⣿⣿⣶⣦⣀⠀" + reset + "⡜⠀⠀⣰⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡇⠀⠀")
	time.Sleep(80 * time.Millisecond)
	fmt.Println("⠀⠀⠀⠀⠀⠀⠀⠀⢀⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣭⣛⣓⠶⠦⠥⣀" + green + "⠙⠋⠉⠉⠻⣄⣀" + reset + "⣸⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡇⠀⠀")
	time.Sleep(80 * time.Millisecond)
	fmt.Println("⠀⠀⠀⠀⠀⠀⠀⠀⣼⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣷⣶⣆" + green + "⠐⣦⣠⣷⠊⠁⠀⠀" + green + "⡭" + reset + "⠙⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡆⠀⠀")
	time.Sleep(80 * time.Millisecond)
	fmt.Println("⠀⠀⠀⠀⠀⠀⠀⢠⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡿⢉" + yellow + "⣛⡛⢻⡗⠂⠀⢀⣷⣄" + green + "⠈⢆" + reset + "⠉⠙⠻⢿⣿⣿⣿⣿⣿⠇⠀⠀")
	time.Sleep(80 * time.Millisecond)
	fmt.Println("⠀⠀⠀⠀⠀⠀⠀⠘⣿⣿⡟⢻⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡟⣉" + green + "⢁⣴⣿⣿⣿⣾⡇" + yellow + "⢀⣀⣼⡿⣿⣷⡌" + green + "⢻⣦⡀" + reset + "⠀⠈⠙⠛⠿⠏⠀⠀⠀")
	time.Sleep(80 * time.Millisecond)
	fmt.Println("⠀⠀⠀⠀⠀⠀⠀⠀⠙⢿⣿⡄⠙⢿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⠿⠛⠛" + green + "⠛⢯⡉⠉⠉⠉⠉⠛⢼⣿⠿⠿⠦" + yellow + "⡙⣿⡆" + green + "⢹⣷⣤⡀⠀⠀⠀⠀⠀⠀⠀")
	time.Sleep(80 * time.Millisecond)
	fmt.Println(reset + "⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠘⠿⠄⠈⠻⠿⠿⠿⠿⠿⠿⠛⠛⠿⠛⠉⠁⠀" + green + "⠀⠀⠀⠀⠀⠻⠿⠿⠿⠿⠟⠉⠀" + yellow + "⠀⠤⠴⠶⠌⠿" + green + "⠘⠿⠿⠿⠿⠶⠤⠀⠀" + reset)
	println()
	time.Sleep(700 * time.Millisecond)
	fmt.Println("If so powerful you are...Why leave ?")
	time.Sleep(3 * time.Second)
	os.Exit(1)
}
