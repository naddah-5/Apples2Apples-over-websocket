package model_test

import (
	"main/model"
	"testing"
)

func generateTestPlayer() model.Player {
	player := model.NewPlayer("test player", true, false, 7)
	return *player
}

func TestGeneratePlayerAndGetters(t *testing.T) {
	player := generateTestPlayer()
	if player.PlayerName() != "test player" {
		t.Log("unexpected player name")
		t.FailNow()
	}
	if !player.Host() {
		t.Log("unexpected host status")
		t.FailNow()
	}
	if player.Bot() {
		t.Log("unexpected bot status")
		t.FailNow()
	}
	if player.Score() != 0 {
		t.Log("unexpected player score")
		t.FailNow()
	}
	if len(player.PlayerHand()) != 0 {
		t.Log("unexpected player hand")
		t.FailNow()
	}
	if player.Score() != 0 {
		t.Log("unexpected player score")
		t.FailNow()
	}
}

func TestPlayerDrawCard(t *testing.T) {
	player := generateTestPlayer()
	testDeck, deckErr := generateTestDeckGA()
	if deckErr != nil {
		t.Log(deckErr)
		t.FailNow()
	}
	drawErr := player.DrawCard(&testDeck)
	if drawErr != nil {
		t.Log(drawErr)
		t.FailNow()
	}
	if len(player.PlayerHand()) != player.HandCapacity() {
		t.Log("expected to draw until hand was full")
		t.FailNow()
	}
	if testDeck.CardsLeft() > (100 - player.HandCapacity()) {
		t.Log("expected cards to be removed from deck")
		t.FailNow()
	}
}

func TestDrawCardError(t *testing.T) {
	player := generateTestPlayer()
	testDeck := *new(model.Deck)
	drawErr := player.DrawCard(&testDeck)
	if drawErr == nil {
		t.Log("uncaught error, drawing cards from empty deck")
		t.FailNow()
	}
}

func TestPlayerDrawCardCombineShuffle(t *testing.T) {
	player := generateTestPlayer()
	testDeck, deckErr := generateTestDeckGA()
	if deckErr != nil {
		t.Log(deckErr)
		t.FailNow()
	}
	
	// Move 95 cards from the testDecks deck to its pile.
	for i := 0; i < 95; i++ {
		rotateCard, rotateErr := testDeck.DrawCard()
		if rotateErr != nil {
			t.Log(rotateErr)
			t.FailNow()
		}
		discardErr := testDeck.DiscardCard(rotateCard)
		if discardErr != nil {
			t.Log(discardErr)
			t.FailNow()
		}
	}
	
	playerDrawErr := player.DrawCard(&testDeck)
	if playerDrawErr != nil {
		t.Log(playerDrawErr)
		t.FailNow()
	}
}

func TestPlayCardZero(t *testing.T) {
	player := generateTestPlayer()
	testDeck, deckErr := generateTestDeckGA()
	if deckErr != nil {
		t.Log(deckErr)
		t.FailNow()
	}
	player.DrawCard(&testDeck)
	playedCard, playErr := player.PlayCard(0)
	if playErr != nil {
		t.Log(playErr)
		t.FailNow()
	}
	if playedCard.DisplayCard() != "[Absurd] - (ridiculous, senseless, foolish)" {
		t.Log("unexpected card:", playedCard.DisplayCard())
		t.FailNow()
	}
	if player.CardsInHand() != player.HandCapacity() - 1 {
		t.Log("expected card to be removed:", player.CardsInHand(), "cards in hand,", player.HandCapacity(), "hand capacity")
		t.FailNow()
	}
}

func TestPlayCardLast(t *testing.T) {
	player := generateTestPlayer()
	testDeck, deckErr := generateTestDeckGA()
	if deckErr != nil {
		t.Log(deckErr)
		t.FailNow()
	}
	player.DrawCard(&testDeck)
	playedCard, playErr := player.PlayCard(player.HandCapacity()-1)
	if playErr != nil {
		t.Log(playErr)
		t.FailNow()
	}
	if playedCard.DisplayCard() != "[Amazing] - (astonishing, surprising, wonderful)" {
		t.Log("unexpected card:", playedCard.DisplayCard())
		t.FailNow()
	}
	if player.CardsInHand() != player.HandCapacity() - 1 {
		t.Log("expected card to be removed:", player.CardsInHand(), "cards in hand,", player.HandCapacity(), "hand capacity")
		t.FailNow()
	}
}

func TestPlayCardCenter(t *testing.T) {
	player := generateTestPlayer()
	testDeck, deckErr := generateTestDeckGA()
	if deckErr != nil {
		t.Log(deckErr)
		t.FailNow()
	}
	player.DrawCard(&testDeck)
	playedCard, playErr := player.PlayCard(player.HandCapacity()/2)
	if playErr != nil {
		t.Log(playErr)
		t.FailNow()
	}
	if playedCard.DisplayCard() != "[Adorable] - (lovable, charming, delightful)" {
		t.Log("unexpected card:", playedCard.DisplayCard())
		t.FailNow()
	}
	if player.CardsInHand() != player.HandCapacity() - 1 {
		t.Log("expected card to be removed:", player.CardsInHand(), "cards in hand,", player.HandCapacity(), "hand capacity")
		t.FailNow()
	}
}

func TestPlayCardInvalid(t *testing.T) {
	player := generateTestPlayer()
	testDeck, deckErr := generateTestDeckGA()
	if deckErr != nil {
		t.Log(deckErr)
		t.FailNow()
	}
	player.DrawCard(&testDeck)
	_, playErrOver := player.PlayCard(player.HandCapacity())
	if playErrOver == nil {
		t.Log("uncaught error, played invalid card index")
		t.FailNow()
	}
	_, playErrUnder := player.PlayCard(-1)
	if playErrUnder == nil {
		t.Log("uncaught error, played invalid card index")
		t.FailNow()
	}
}

func TestIncreaseScore(t *testing.T) {
	player := generateTestPlayer()
	valueCard := model.MintCard("valuable card", "header", "description")
	initScore := player.Score()
	updatedScore := player.IncreaseScore(valueCard)
	if updatedScore <= initScore {
		t.Log("expected score to increase")
		t.FailNow()
	}
}

func TestShowHand(t *testing.T) {
	testDeck, deckErr := generateTestDeckRA()
	if deckErr != nil {
		t.Log(deckErr)
		t.FailNow()
	}
	player := model.NewPlayer("player one", true, false, 7)
	player.DrawCard(&testDeck)
	hand := player.ShowHand()
	var expectedHand []string = []string{
		"[A Bad Haircut] - The perfect start to a bad hair day.",
		"[A Bakery] - Some bakers start work at 3:00 in the morning, so breads and donuts are fresh for  breakfast.",
		"[A Broken Leg] - I was riding my bike when I hit this big rock . . .",
		"[A Bull Fight] - Also known as \"la fiesta brava\" (the brave festival).  A whole lot of bull..",
		"[A Cabin In The Woods] - Henry David Thoreau went to Walden Pond for two years.  All we're asking for is one lousy weekend!",
		"[A Can Of Worms] - Now you've opened it.",
		"[A Car Crash] - \"Hey, it was an accident!\"",
	}
	for i := 0; i < len(hand); i++ {
		if hand[i] != expectedHand[i] {
			t.Log("expected: ", expectedHand[i], "\nreceived: ", hand[i])
			t.FailNow()
		}
	}
}