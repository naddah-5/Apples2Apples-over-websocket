package controller

import (
	"bufio"
	"fmt"
	"main/model"
	"main/view"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

func Game() {
	terminal := view.Terminal()
	view.Greeting(terminal)
	for terminal.Scan() {
		input := terminal.Text()
		switch input {
		case "1":
			fmt.Println("Online play is not yet implemented")
		case "2":
			fmt.Println("Online play is not yet implemented")
		case "3":
			board := setupOfflineGame(terminal)
			playGame(terminal, board)
		case "4":
			os.Exit(0)
		default:
			fmt.Println("Please select one of the options")
		}
	}

}

func setupOfflineGame(terminal bufio.Scanner) *model.Board {
	/*
	Prompt the player for a player name
	=======================================================================
	*/
	view.ChooseName()
	var playerName string
	terminal.Scan()
	playerName = terminal.Text()
	fmt.Println("You entered", playerName)
	fmt.Println("1) confirm name\n not 1) back")
	terminal.Scan() 
	var confirmation string = terminal.Text()
	if confirmation != "1" {
		return setupOfflineGame(terminal)
	}

	/*
	Create players and add them to the board.
	=======================================================================
	*/
	board := new(model.Board)
	board.AddPlayer(*model.NewPlayer(playerName, true, false, 7))
	board.AddPlayer(*model.NewPlayer("Bot1", false, true, 7))
	board.AddPlayer(*model.NewPlayer("Bot2", false, true, 7))
	board.AddPlayer(*model.NewPlayer("Bot3", false, true, 7))

	/*
	Load the card decks and add them to the board.
	=======================================================================
	*/
	absRedPath, redPathErr := filepath.Abs("../src/resources/redApples.txt")
	if redPathErr != nil {
		fmt.Println("could not load red apples path ", redPathErr)
		panic(redPathErr)
	}
	redPathErr = board.LoadRedApples(absRedPath)
	if redPathErr != nil {
		fmt.Println("could not load red apples to board ", redPathErr)
		panic(redPathErr)
	}

	absGreenPath, greenPathErr := filepath.Abs("../src/resources/greenApples.txt")
	if greenPathErr != nil {
		fmt.Println("could not load green apples path ", greenPathErr)
		panic(greenPathErr)
	}
	greenPathErr = board.LoadGreenApples(absGreenPath)
	if greenPathErr != nil {
		fmt.Println("could not load green apples to board ", greenPathErr)
		panic(greenPathErr)
	}

	/*
	Shuffle both card decks.
	=======================================================================
	*/
	shuffleGreenErr := board.ShuffleGreenApples()
	if shuffleGreenErr != nil {
		fmt.Println("could not shuffle green apples ", shuffleGreenErr)
		panic(shuffleGreenErr)
	}

	shuffleRedErr := board.ShuffleRedApples()
	if shuffleRedErr != nil {
		fmt.Println("could not shuffle red apples ", shuffleGreenErr)
	}

	/*
	Shuffle the player order.
	=======================================================================
	*/
	shufflePlayerErr := board.ShufflePlayers()
	if shufflePlayerErr != nil {
		fmt.Println("could not shuffle player order ", shufflePlayerErr)
		panic(shufflePlayerErr)
	}

	/*
	Deal out the starting cards to players.
	=======================================================================
	*/
	drawCardErr := board.FillHands()
	if drawCardErr != nil {
		fmt.Println("could not draw initial player hands ", drawCardErr)
		panic(drawCardErr)
	}

	/*
	Randomize the starting judge.
	=======================================================================
	*/
	initJudgeErr := board.InitializeJudge()
	if initJudgeErr != nil {
		fmt.Println("could not initialize judge ", initJudgeErr)
		panic(initJudgeErr)
	}

	/*
	Set the win condition.
	=======================================================================
	*/
	winErr := board.SetWinCondition()
	if winErr != nil {
		fmt.Println("could not set the win condition ", winErr)
		panic(winErr)
	}



	return board
}


func setupOnlineGame(terminal bufio.Scanner) *model.Board {
	/*
	Prompt the player for a player name
	=======================================================================
	*/
	view.ChooseName()
	var playerName string
	terminal.Scan()
	playerName = terminal.Text()
	fmt.Println("You entered", playerName)
	fmt.Println("1) confirm name\n not 1) back")
	terminal.Scan() 
	var confirmation string = terminal.Text()
	if confirmation != "1" {
		return setupOnlineGame(terminal)
	}


	/*
	Establish connections.
	=======================================================================
	*/
	fmt.Println("How many online players?")
	terminal.Scan()
	onlinePlayers, parseErr := strconv.ParseInt(terminal.Text(), 10, 64)
	for parseErr != nil && int(onlinePlayers) < 1 {
		fmt.Println("Please enter an integer larger than zero")		
		fmt.Println("How many online players?")
		terminal.Scan()
		onlinePlayers, parseErr = strconv.ParseInt(terminal.Text(), 10, 64)
	}
	network := new(model.Network)
	go network.Listener()
	fmt.Println("Waiting for players to connect to loaclhost, port 8080...")
	for {
		time.Sleep(1 * time.Second)
		if network.CountOnlinePlayers() == int(onlinePlayers) {
			break
		}
		fmt.Println(network.CountOnlinePlayers(), "players connected...")
	}
	fmt.Println("All players connected!")
	
	

	/*
	Create players and add them to the board.
	=======================================================================
	*/
	board := new(model.Board)
	board.AddPlayer(*model.NewPlayer(playerName, true, false, 7))
	board.AddPlayer(*model.NewPlayer("Bot1", false, true, 7))
	board.AddPlayer(*model.NewPlayer("Bot2", false, true, 7))
	board.AddPlayer(*model.NewPlayer("Bot3", false, true, 7))

	/*
	Load the card decks and add them to the board.
	=======================================================================
	*/
	absRedPath, redPathErr := filepath.Abs("../src/resources/redApples.txt")
	if redPathErr != nil {
		fmt.Println("could not load red apples path ", redPathErr)
		panic(redPathErr)
	}
	redPathErr = board.LoadRedApples(absRedPath)
	if redPathErr != nil {
		fmt.Println("could not load red apples to board ", redPathErr)
		panic(redPathErr)
	}

	absGreenPath, greenPathErr := filepath.Abs("../src/resources/greenApples.txt")
	if greenPathErr != nil {
		fmt.Println("could not load green apples path ", greenPathErr)
		panic(greenPathErr)
	}
	greenPathErr = board.LoadGreenApples(absGreenPath)
	if greenPathErr != nil {
		fmt.Println("could not load green apples to board ", greenPathErr)
		panic(greenPathErr)
	}

	/*
	Shuffle both card decks.
	=======================================================================
	*/
	shuffleGreenErr := board.ShuffleGreenApples()
	if shuffleGreenErr != nil {
		fmt.Println("could not shuffle green apples ", shuffleGreenErr)
		panic(shuffleGreenErr)
	}

	shuffleRedErr := board.ShuffleRedApples()
	if shuffleRedErr != nil {
		fmt.Println("could not shuffle red apples ", shuffleGreenErr)
	}

	/*
	Shuffle the player order.
	=======================================================================
	*/
	shufflePlayerErr := board.ShufflePlayers()
	if shufflePlayerErr != nil {
		fmt.Println("could not shuffle player order ", shufflePlayerErr)
		panic(shufflePlayerErr)
	}

	/*
	Deal out the starting cards to players.
	=======================================================================
	*/
	drawCardErr := board.FillHands()
	if drawCardErr != nil {
		fmt.Println("could not draw initial player hands ", drawCardErr)
		panic(drawCardErr)
	}

	/*
	Randomize the starting judge.
	=======================================================================
	*/
	initJudgeErr := board.InitializeJudge()
	if initJudgeErr != nil {
		fmt.Println("could not initialize judge ", initJudgeErr)
		panic(initJudgeErr)
	}

	/*
	Set the win condition.
	=======================================================================
	*/
	winErr := board.SetWinCondition()
	if winErr != nil {
		fmt.Println("could not set the win condition ", winErr)
		panic(winErr)
	}



	return board
}

func playGame(terminal bufio.Scanner, board *model.Board) {
	for {
		/*
		Check for win condition.
		===============================================================
		*/
		gameOver, winErr := board.GameWinner()
		if winErr != nil {
			resetWinErr := board.SetWinCondition()
			if resetWinErr != nil {
				panic(winErr)
			}
		}
		if gameOver {
			winner, falseWin := board.WhoWonGame()
			if falseWin != nil {
				panic(falseWin)
			}
			view.Winner(winner.PlayerName())
		} else {
			/*
			Start a new round.
			=======================================================
			*/
			scoreBoard := board.DisplayScoreBoard()
			view.ScoreBoard(scoreBoard)
			roundErr := playRound(terminal, board)
			if roundErr != nil {
				panic(roundErr)
			}
		}
	}
}

func playRound(terminal bufio.Scanner, board *model.Board) error {
	/*
	Draw a green apple and put it on the board.
	=======================================================================
	*/
	drawGreenErr := board.DrawGreenApple()
	if drawGreenErr != nil {
		fmt.Println("could not draw green apple ", drawGreenErr)
		return drawGreenErr
	}

	/*
	Prompt all players, except the judge, to play a red apple.
	=======================================================================
	*/
	playErr := board.ChooseCards()
	if playErr != nil {
		fmt.Println("something went wrong during the round, ", playErr)
		return playErr
	}

	/*
	Prompt judge for decision.
	=======================================================================
	*/
	winningCardIndex, judgeErr := board.Judge()
	if judgeErr != nil {
		fmt.Println("could not recieve judge decision ", judgeErr)
		return judgeErr
	}

	winner, indexErr := board.PlayedCards.ShowPlayer(winningCardIndex)
	if indexErr != nil {
		fmt.Println("tried to find player using invalid index ", indexErr)
		return indexErr
	}
	fmt.Println(winner, "won the round")

	greenApple, pickErr := board.PickUpGreenApple()
	if pickErr != nil {
		fmt.Println("did not find the green apple ", pickErr)
		return pickErr
	}
	board.AwardScore(winner, greenApple)

	/*
	Discard played cards.
	=======================================================================
	*/
	disErr := board.DiscardRound()
	if disErr != nil {
		fmt.Println("could not discard the played cards, ", disErr)
		return disErr
	}

	/*
	Players draw new cards.
	=======================================================================
	*/
	drawErr := board.FillHands()
	if drawErr != nil {
		fmt.Println("could not draw new cards, ", drawErr)
		return drawErr
	}

	/*
	Itterate to the next judge.
	=======================================================================
	*/
	board.ItterateJudge()

	return nil
}