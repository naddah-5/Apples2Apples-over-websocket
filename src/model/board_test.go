package model_test

import (
	"main/model"
	"testing"
)

func TestRandomJudge(t *testing.T) {
	playerOne := *model.NewPlayer("player one", true, false, 7)
	playerTwo := *model.NewPlayer("player two", true, false, 7)
	playerThree := *model.NewPlayer("player three", true, false, 7)
	playerFour := *model.NewPlayer("player four", true, false, 7)
	
	var board model.Board = *new(model.Board)

	pOneErr := board.AddPlayer(playerOne)
	pTwoErr := board.AddPlayer(playerTwo)
	pThreeErr := board.AddPlayer(playerThree)
	pFourErr := board.AddPlayer(playerFour)

	if pOneErr != nil || pTwoErr != nil || pThreeErr != nil || pFourErr != nil {
		t.Log("incorrect testing configuration")
		t.FailNow()
	}

	shufErr := board.ShufflePlayers()
	if shufErr != nil {
		t.Log("unexpected shuffling error:", shufErr)
		t.FailNow()
	}
	var coincidence int = 0
	var confidense int = 10_000
	var delim int = 7_500
	var prevJudge string = board.CurrentJudgeName()

	// verify that judge itteration works
	for i := 0; i < 10_000; i++ {
		board.ItterateJudge()
		if board.CurrentJudgeName() == prevJudge {
			t.Log("judge did not itterate")
			t.FailNow()
		}
		prevJudge = board.CurrentJudgeName()
	}

	for i:= 0; i < confidense; i++ {
		shufErr := board.ShufflePlayers()
		if shufErr != nil {
			t.Log("unexpected shuffling error:", shufErr)
			t.FailNow()
		}
		if board.CurrentJudgeName() == prevJudge {
			coincidence++
		}
		prevJudge = board.CurrentJudgeName()
	}
	if coincidence >= delim {
		t.Log("warning: unexpected shuffling distribution")
		t.FailNow()
	}
}

func TestSetWinCondition(t *testing.T) {
	playerOne := *model.NewPlayer("player one", true, false, 7)
	playerTwo := *model.NewPlayer("player two", false, false, 7)
	playerThree := *model.NewPlayer("player three", false, false, 7)
	playerFour := *model.NewPlayer("player four", false, false, 7)
	
	var board model.Board = *new(model.Board)

	wErr := board.SetWinCondition()
	if wErr == nil {
		t.Log("did not catch, not enough players error")
		t.FailNow()
	}

	pOneErr := board.AddPlayer(playerOne)
	pTwoErr := board.AddPlayer(playerTwo)
	pThreeErr := board.AddPlayer(playerThree)
	
	if pOneErr != nil || pTwoErr != nil || pThreeErr != nil {
		t.Log("incorrect testing configuration")
		t.FailNow()
	}
	
	wErr2 := board.SetWinCondition()
	if wErr2 == nil {
		t.Log("did not catch, not enough players error")
		t.FailNow()
	}
	
	pFourErr := board.AddPlayer(playerFour)
	if pFourErr != nil {
		t.Log("incorrect testing configuration")
		t.FailNow()
	}

	fourErr := board.SetWinCondition()
	if fourErr != nil {
		t.Log("unexpected error:", fourErr)
		t.FailNow()
	}
	if board.GetWinCondition() != 8 {
		t.Log("incorrect win condition, expected win condition to be 8")
		t.FailNow()
	}

	playerFive := *model.NewPlayer("player five", false, false, 7)
	pFiveErr := board.AddPlayer(playerFive)
	if pFiveErr != nil {
		t.Log("incorrect testing configuration")
		t.FailNow()
	}
	fiveErr := board.SetWinCondition()
	if fiveErr != nil {
		t.Log("unexpected error:", fiveErr)
		t.FailNow()
	}
	if board.GetWinCondition() != 7 {
		t.Log("incorrect win condition, expected win condition to be 7")
		t.FailNow()
	}

	playerSix := *model.NewPlayer("player six", false, false, 7)
	psixErr := board.AddPlayer(playerSix)
	if psixErr != nil {
		t.Log("incorrect testing configuration")
		t.FailNow()
	}
	sixErr := board.SetWinCondition()
	if sixErr != nil {
		t.Log("unexpected error:", sixErr)
	}
	if board.GetWinCondition() != 6 {
		t.Log("incorrect win condition, expected win condition to be 6")
		t.FailNow()
	}

	playerSeven := *model.NewPlayer("player seven", false, false, 7)
	psevenErr := board.AddPlayer(playerSeven)
	if psevenErr != nil {
		t.Log("incorrect testing configuration")
		t.FailNow()
	}
	sevenErr := board.SetWinCondition()
	if sevenErr != nil {
		t.Log("unexpected error:", sevenErr)
		t.FailNow()
	}
	if board.GetWinCondition() != 5 {
		t.Log("incorrect win condtion, expected win condition to be 5")
		t.FailNow()
	}

	playerEight := *model.NewPlayer("player eight", false, false, 7)
	peightErr := board.AddPlayer(playerEight)
	if peightErr != nil {
		t.Log("incorrect testing configuration")
	}
	eightErr := board.SetWinCondition()
	if eightErr != nil {
		t.Log("unexpected error:", eightErr)
		t.FailNow()
	}
	if board.GetWinCondition() != 4 {
		t.Log("incorrect win condition, expected win condition to be 4")
		t.FailNow()
	}
	
	playerNine := *model.NewPlayer("player nine", false, false, 7)
	pnineErr := board.AddPlayer(playerNine)
	if pnineErr != nil {
		t.Log("incorrect testing configuration")
		t.FailNow()
	}
	nineErr := board.SetWinCondition()
	if nineErr != nil {
		t.Log("unexpected error:", nineErr)
		t.FailNow()
	}
	if board.GetWinCondition() != 4 {
		t.Log("incorrect win condition, expected win condition to be 4")
		t.FailNow()
	}	
}

func TestWinCheck(t *testing.T) {
	playerOne := *model.NewPlayer("player one", true, false, 7)
	playerTwo := *model.NewPlayer("player two", false, false, 7)
	playerThree := *model.NewPlayer("player three", false, false, 7)
	playerFour := *model.NewPlayer("player four", false, false, 7)
	
	var board model.Board = *new(model.Board)

	board.AddPlayer(playerOne)
	board.AddPlayer(playerTwo)
	board.AddPlayer(playerThree)
	board.AddPlayer(playerFour)
	board.SetWinCondition()

	testDeck, testErr := generateTestDeck()
	if testErr != nil {
		t.Log("incorrectly configured test, generate deck error detected", testErr)
		t.FailNow()
	}
	for i := 0; i < 4; i++ {
		scoreCard, drawErr := testDeck.DrawCard()
		if drawErr != nil {
			t.Log("incorrectly configured test, draw error detected", drawErr)
		}
		board.AwardScore(playerOne.PlayerName(), scoreCard)
	}

	winConditionMet, winErr := board.Winner()
	if winErr != nil {
		t.Log("unexpected error: ", winErr)
		t.FailNow()
	}
	if winConditionMet {
		t.Log("unexpected win condition")
		t.FailNow()
	}

	for i := 0; i < 8; i++ {
		scoreCard, drawErr := testDeck.DrawCard()
		if drawErr != nil {
			t.Log("incorrectly configured test, draw error detected", drawErr)
		}
		board.AwardScore(playerTwo.PlayerName(), scoreCard)
	}

	winConditionMet2, _ := board.Winner()
	if !winConditionMet2 {
		t.Log("expected to have a winner")
		t.FailNow()
	}

	winner, winConErr := board.WhoWon()
	if winConErr != nil {
		t.Log("unexpected error:", winConErr)
		t.FailNow()
	}
	winnerName := winner.PlayerName()
	if winnerName != "player two" {
		t.Log("expected player two to win")
		t.FailNow()
	}
}

func TestWinConditionMorePlayers(t *testing.T) {
	playerOne := *model.NewPlayer("player one", true, false, 7)
	playerTwo := *model.NewPlayer("player two", false, false, 7)
	playerThree := *model.NewPlayer("player three", false, false, 7)
	playerFour := *model.NewPlayer("player four", false, false, 7)
	playerFive := *model.NewPlayer("player five", false, false, 7)
	playerSix := *model.NewPlayer("player six", false, false, 7)
	playerSeven := *model.NewPlayer("player seven", false, false, 7)
	


	var board model.Board = *new(model.Board)

	board.AddPlayer(playerOne)
	board.AddPlayer(playerTwo)
	board.AddPlayer(playerThree)
	board.AddPlayer(playerFour)
	board.AddPlayer(playerFive)
	board.AddPlayer(playerSix)
	board.AddPlayer(playerSeven)
	board.SetWinCondition()

	testDeck, testErr := generateTestDeck()
	if testErr != nil {
		t.Log("incorrectly configured test, generate deck error detected", testErr)
		t.FailNow()
	}

	for i := 0; i < 5; i++ {
		scoreCard, drawErr := testDeck.DrawCard()
		if drawErr != nil {
			t.Log("incorrectly configured test, draw error detected", drawErr)
		}
		board.AwardScore(playerThree.PlayerName(), scoreCard)
	}
	winConditionMet, winErr := board.Winner()
	if winErr != nil {
		t.Log("unexpected error: ", winErr)
		t.FailNow()
	}
	if !winConditionMet {
		t.Log("expected win condition")
		t.FailNow()
	}
	winner, winPlayErr := board.WhoWon()
	if winPlayErr != nil {
		t.Log("expected a winner, received error:", winPlayErr)
		t.FailNow()
	}
	winnerName := winner.PlayerName()
	if winnerName != "player three" {
		t.Log("unexpected winner, expected player three to win, actually ", winnerName)
		t.FailNow()
	}
}