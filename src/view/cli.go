package view

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

/*
Displays a greeting message that is set locally.
The message should display all available options
and take the corresponding action on valid input.
*/
func Greeting() {
	var GREETING string = "Hello, do you want to play a game. \n 1) Host game\n 2) Join game\n 3) Play bots\n 4) Exit"
	err := clear()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(GREETING)
	terminal := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(terminal)

	for scanner.Scan() {
		input := scanner.Text()
		parseGreeting(input)
	}
}

/*
Decides what to do based on the input.
*/
func parseGreeting(input string) {
	switch input {
	case "1":
		fmt.Println("you choose wisely", input)
	case "2":
		fmt.Println("you choose poorly", input)
	case "3":
		fmt.Println("are you confused?", input)
	case "4":
		os.Exit(0)
	default:
		fmt.Println("Please select one of the options, input the number corresponding to your choice.")
	}
}

/*
Attempt to clear the terminal screen, if the OS is unsupported returns an error.
*/
func clear() error {
	if runtime.GOOS == "linux" {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
		return nil
	} else if runtime.GOOS == "windows" {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
		return nil
	}
	return errors.New("Unsupported OS")
}


func ChooseCard() int {
	return 0
}