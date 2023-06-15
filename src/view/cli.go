package view

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
)

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

func Terminal() bufio.Scanner {
	terminal := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(terminal)

	return *scanner
}

/*
Displays a greeting message that is set locally.
The message should display all available options
and take the corresponding action on valid input.
*/
func Greeting(scanner bufio.Scanner)  {
	var GREETING string = "Hello, do you want to play a game. \n 1) Host game\n 2) Join game\n 3) Play bots\n 4) Exit"
	err := clear()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(GREETING)
}

/*
Prompt user to choose a player name.
*/
func ChooseName() {
	clear()
	fmt.Println("Please choose a player name: ")
}

/*
Print out the players hand and current green apple, take which card to play from terminal in the form of an index int.
*/
func ChooseCard(greenApple string, hand []string) int {
	clear()
	fmt.Println("The current green apple is ", greenApple)
	fmt.Println("Please select a card to play:")
	for i := 0; i < len(hand); i++ {
		fmt.Println("[", i, "]: ", hand[i])
	}
	terminal := Terminal()
	for terminal.Scan() {
		input := terminal.Text()
		choice, convErr := strconv.ParseInt(input, 10, 64)
		if convErr != nil {
			fmt.Println("Please only enter the integer representation of your choice.")
			continue
		}
		return int(choice)
	}
	WaitPlayerCards()
	return 0
}

func WaitPlayerCards() {
	clear()
	fmt.Println("Waiting for players to submit cards...")
}

/*
Print out the submitted red apples and current green apple, take the winning red apples index from terminal in the 
form of an index int.
*/
func JudgeCards(greenApple string, redApples []string) int {
	clear()
	fmt.Println("Current green apple ", greenApple)
	fmt.Println("Submitted red apples")
	for i := 0; i < len(redApples); i++ {
		fmt.Println("[", i, "]", redApples[i])
	}

	terminal := Terminal()
	var choice int
	for terminal.Scan() {
		input := terminal.Text()
		parseChoice, convErr := strconv.ParseInt(input, 10, 64)
		if convErr != nil {
			fmt.Println("Please only enter the integer representation of your choice.")
			continue
		}
		choice = int(parseChoice)
		if choice < 0 || choice >= len(redApples) {
			fmt.Println("Please choose a valid option")
		} else {
			break
		}
	}
	return choice
}


func DisplaySubmissions(greenApple string, redApples []string) {
	clear()
	fmt.Println("Current green apple ", greenApple)
	fmt.Println("Submitted red apples")
	for i := 0; i < len(redApples); i++ {
		fmt.Println(redApples[i])
	}
	fmt.Println("Waiting for judgement...")
}

func Winner(name string) {
	fmt.Println(name, " has won the game, congratulation!")
	os.Exit(0)
}