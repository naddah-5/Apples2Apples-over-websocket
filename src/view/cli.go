package view

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

/*
Attempt to clear the terminal screen, if the OS is unsupported returns an error.
*/
func clear() error {
	for i := 0; i < 1; i++ {
		fmt.Println()
	}
	return nil
	
	// if runtime.GOOS == "linux" {
	// 	cmd := exec.Command("clear")
	// 	cmd.Stdout = os.Stdout
	// 	cmd.Run()
	// 	return nil
	// } else if runtime.GOOS == "windows" {
	// 	cmd := exec.Command("cmd", "/c", "cls")
	// 	cmd.Stdout = os.Stdout
	// 	cmd.Run()
	// 	return nil
	// }
	// return errors.New("Unsupported OS")
}

/*
Return an input scanner.
*/
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
	var GREETING string = "Hello, do you want to play a game. \n 1) Play bots\n 2) Host game\n 3) Join game\n 4) Exit"
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
func ChooseName(terminal bufio.Scanner) (string, error) {
	clear()
	fmt.Println("Please enter player name:")
	terminal.Scan()
	playerName := terminal.Text()
	fmt.Println("You entered", playerName)
	fmt.Println("1) confirm name\n 2) select name\n 3) exit game")
	for terminal.Scan() {
		var input string = terminal.Text()
		if input == "1" {
			return playerName, nil
		} else if input == "2" {
			return ChooseName(terminal)
		} else if input == "3" {
			return "", errors.New("exit game")
		}

	}
	return ChooseName(terminal)
}

/*
Prompt the user for the number of online players.
*/
func OnlinePlayers(terminal bufio.Scanner) int {
	fmt.Println("How many online players?")
	terminal.Scan()
	onlinePlayers, parseErr := strconv.ParseInt(terminal.Text(), 10, 64)
	for parseErr != nil && int(onlinePlayers) < 1 {
		fmt.Println("Please enter an integer larger than zero")		
		fmt.Println("How many online players?")
		terminal.Scan()
		onlinePlayers, parseErr = strconv.ParseInt(terminal.Text(), 10, 64)
	}
	return int(onlinePlayers)
}

/*
Print out the players hand and current green apple, take which card to play from terminal in the form of an index int.
*/
func ChooseCard(greenApple string, hand []string) int {
	clear()
	fmt.Println("The current green apple is", greenApple)
	fmt.Println("Please select a card to play:")
	for i := 0; i < len(hand); i++ {
		fmt.Println("[", i, "]: ", hand[i])
	}
	fmt.Println("Select card by submitting its index:")
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
	fmt.Println("Current green apple", greenApple)
	fmt.Println("Submitted red apples")
	for i := 0; i < len(redApples); i++ {
		fmt.Println("[", i, "]", redApples[i])
	}
	fmt.Println("Select winning card by submitting its index:")

	terminal := Terminal()
	var choice int
	for terminal.Scan() {
		input := terminal.Text()
		parseChoice, convErr := strconv.ParseInt(input, 10, 64)
		if convErr != nil {
			fmt.Println("Please enter a valid option.")
			continue
		}
		choice = int(parseChoice)
		if choice < 0 || choice >= len(redApples) {
			fmt.Println("Please select a valid option")
		} else {
			break
		}
	}
	return choice
}

/*
Displays the submitted apples to the player.
*/
func DisplaySubmissions(greenApple string, redApples []string) {
	clear()
	fmt.Println("Current green apple", greenApple)
	fmt.Println("Submitted red apples")
	for i := 0; i < len(redApples); i++ {
		fmt.Println(redApples[i])
	}
	fmt.Println("Waiting for judgement...")
}

/*
Display the game winner to the user.
*/
func Winner(name string) {
	fmt.Println(name, "has won the game, congratulation!")
}

/*
Displays all players and their current score to the user.
*/
func ScoreBoard(score []string) {
	clear()
	for i := 0; i < len(score); i++ {
		fmt.Println(score[i])
	}
}

/*
Displays the incomming message and returns the users input.

If the scan loop is somehow broken, returns default value 0.
*/
func OnlinePlay(validInputLimit int, display []string) string {
	for i := 0; i < len(display); i++ {
		fmt.Println(display[i])
	}
	terminal := Terminal()
	for terminal.Scan() {
		inputStr := terminal.Text()
		input64, _ := strconv.ParseInt(inputStr, 10, 64)
		var input int = int(input64)
		if input >= 0 && input <= validInputLimit {
			return inputStr
		}
	}
	return "0"
}

/*
Display the received message.
*/
func OnlineDisplay(display []string) {
	for i := 0; i < len(display); i++ {
		fmt.Println(display[i])
	}
}