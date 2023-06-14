package model_test

import (
	"main/model"
	"testing"
)

func generateTestPA(testDeck *model.Deck) model.PlayedApples {	
	var roundList []model.PlayerPlayed
	playerOne := *model.NewPlayer("player one", true, false, 7)
	playerTwo := *model.NewPlayer("player two", false, true, 7)
	playerThree := *model.NewPlayer("player three", false, true, 7)
	playerFour := *model.NewPlayer("player four", false, true, 7)

	cardOne, _ := testDeck.DrawCard()
	cardTwo, _ := testDeck.DrawCard()
	cardThree, _ := testDeck.DrawCard()
	cardFour, _ := testDeck.DrawCard()

	roundList = append(roundList, model.PlayerPlays(&playerOne, cardOne))
	roundList = append(roundList, model.PlayerPlays(&playerTwo, cardTwo))
	roundList = append(roundList, model.PlayerPlays(&playerThree, cardThree))
	roundList = append(roundList, model.PlayerPlays(&playerFour, cardFour))

	pa := model.PlayedRound(roundList)
	return pa
}

/*
Statistical test for the player shuffling.
*/
func TestPAShuffle(t *testing.T) {
	testDeck, deckErr := generateTestDeckGA()
	if deckErr != nil {
		t.Log("incorrect test config,", deckErr)
	}
	pa := generateTestPA(&testDeck)
	var coincidence int = 0
	var confidense int = 10_000
	var prevFirst string
	prevFirst, _ = pa.ShowPlayer(0)
	for i := 0; i < confidense; i++ {
		pa.Shuffle()
		first, _ := pa.ShowPlayer(0)
		if prevFirst == first {
			coincidence++ 
		}
		prevFirst = first
	}
	if coincidence > 3*(confidense/4) {
		t.Log("warning: unexpected shuffling statistic; first player match", coincidence, "out of", confidense)
		t.FailNow()
	}
}

func TestDisplayApple(t *testing.T) {
	testDeck, deckErr := generateTestDeckGA()
	if deckErr != nil {
		t.Log("incorrect test config,", deckErr)
		t.FailNow()
	}
	pa := generateTestPA(&testDeck)
	
	printList, showErr := pa.DisplayApples()
	if showErr != nil {
		t.Log(showErr)
		t.FailNow()
	}

	expected := []string{	"[Absurd] - ridiculous, senseless, foolish", 
				"[Abundant] - plentiful, ample, numerous",
				"[Addictive] - obsessive, consuming, captivating",
				"[Adorable] - lovable, charming, delightful"}
	for i := 0; i < len(printList); i++ {
		if printList[i] != expected[i] {
			t.Log("expected,", expected[i], "received,", printList[i])
			t.FailNow()
		}
	}
}

func TestShowPlayers(t *testing.T) {
	testDeck, deckErr := generateTestDeckGA()
	if deckErr != nil {
		t.Log("incorrect test config,", deckErr)
		t.FailNow()
	}
	pa := generateTestPA(&testDeck)
	playerOne, showErrOne := pa.ShowPlayer(0)
	if showErrOne != nil {
		t.Log("unexpected error:", showErrOne)
		t.FailNow()
	}
	if playerOne != "player one"{
		t.Log("expected, player one, received", playerOne)
		t.FailNow()
	}
	playerTwo, showErrTwo := pa.ShowPlayer(1)
	if showErrTwo != nil {
		t.Log("unexpected error:", showErrTwo)
		t.FailNow()
	}
	if playerTwo != "player two" {
		t.Log("expected, player two, received", playerTwo)
		t.FailNow()
	}
	playerThree, showErrThree := pa.ShowPlayer(2)
	if showErrThree != nil {
		t.Log("unexpected error:", showErrThree)
	}
	if playerThree != "player three" {
		t.Log("expected, player three, received", playerThree)
	}
	playerFour, showErrFour := pa.ShowPlayer(3)
	if showErrFour != nil {
		t.Log("unexpected error:", showErrFour)
		t.FailNow()
	}
	if playerFour != "player four" {
		t.Log("expected, player four, received", playerFour)
	}
}

func TestShowPlayerInvalid(t *testing.T) {
	testDeck, deckErr := generateTestDeckGA()
	if deckErr != nil {
		t.Log("incorrect test config,", deckErr)
		t.FailNow()
	}
	pa := generateTestPA(&testDeck)
	_, underErr := pa.ShowPlayer(-1)
	if underErr == nil {
		t.Log("uncaught out of bounds error")
		t.FailNow()
	}
	_, overErr := pa.ShowPlayer(4)
	if overErr == nil {
		t.Log("uncaught out of bounds error")
		t.FailNow()
	}
}

func TestDiscardRound(t *testing.T) {
	testDeck, deckErr := generateTestDeckGA()
	if deckErr != nil {
		t.Log("incorrect test config,", deckErr)
		t.FailNow()
	}
	pa := generateTestPA(&testDeck)
	disErr := pa.DiscardRound(&testDeck)
	if disErr != nil {
		t.Log(disErr)
		t.FailNow()
	}
	if testDeck.CardsInPile() == 0 {
		t.Log("expected to have discarded cards")
		t.FailNow()
	}
	if testDeck.CardsInPile() + testDeck.CardsLeft() != 100 {
		t.Log("expected card sum to equual 100, cardsum is:", testDeck.CardsInPile() + testDeck.CardsLeft())
		t.FailNow()
	}
}