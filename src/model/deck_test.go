package model_test

import (
	"errors"
	"fmt"
	"main/model"
	"path/filepath"
	"testing"
)

func generateTestDeck() (model.Deck, error) {
	absPath, pathErr := filepath.Abs("../resources/testSet.txt")
	if pathErr != nil {
		return *new(model.Deck), errors.New("test incorrectly configured, invalid resource path")
	}
	testDeck, deckErr := model.GenerateDeck(absPath, "green apple")
	if deckErr != nil {
		return *new(model.Deck), deckErr
	}
	return testDeck, nil
}

func TestGeneratePile(t *testing.T) {
	testDeck, deckGenErr := generateTestDeck()
	if deckGenErr != nil {
		t.Log(deckGenErr)
		t.Fail()
	}
	var expectedCard string = "[Absurd] - ridiculous, senseless, foolish"
	card, emptyDeckErr := testDeck.DrawCard()
	if emptyDeckErr != nil {
		t.Log(emptyDeckErr)
		t.Fail()
	}
	if card.DisplayCard() != expectedCard {
		fmt.Println("received ", card.DisplayCard())
		fmt.Println("expected ", expectedCard)
		t.Fail()
	}
}

func TestEmptyDeckDraw(t *testing.T) {
	emptyDeck := new(model.Deck)
	_, err := emptyDeck.DrawCard()
	if err == nil {
		t.Log("draw from empty deck not caught")
		t.Fail()
	}
}

func TestDiscardCardValid(t *testing.T) {
	testDeck, deckErr := generateTestDeck()
	if deckErr != nil {
		t.Log(deckErr)
		t.Fail()
	}
	testCard, err := testDeck.DrawCard()
	if err != nil {
		t.Log("unexpected error, ", err)
		t.Fail()
	}
	discarErr := testDeck.DiscardCard(testCard)
	if discarErr != nil {
		t.Log("unexpected error ", discarErr)
		t.Fail()
	}
	if testDeck.CardsInPile() + testDeck.CardsLeft() != 100 {
		t.Log("expected 100 cards in total, found", testDeck.CardsInPile() + testDeck.CardsLeft())
		t.FailNow()
	}
}
func TestDiscardCardInvalid(t *testing.T) {
	testDeck, deckErr := generateTestDeck()
	if deckErr != nil {
		t.Log(deckErr)
		t.Fail()
	}
	testCard := model.MintCard("broken type", "default header", "default description")
	discarErr := testDeck.DiscardCard(testCard)
	if discarErr == nil {
		t.Log("failed to catch invalid discard")
		t.Fail()
	}
}


/*
This is a statistical unit test, it aims to verify that the shuffle method
does in fact shuffle. If one of the first 100 cards have the same index 
after a deck had been shuffled as it did before, the coincidence counter is
increased. If the number of coincidences is more than half of the confidence
the shuffling is deemed not random.
*/
func TestShuffle(t *testing.T) {
	var cardComparisons int = 0
	var coincidense int = 0
	var repeatFor int = 10_000
	for n := 0; n < repeatFor; n++ {
		testDeck, deckGenErr := generateTestDeck()
		initialDeck, _ := generateTestDeck()
		if deckGenErr != nil {
			t.Log(deckGenErr)
			t.Fail()
		}
		shuffleErr := testDeck.ShuffleDeck()
		if shuffleErr != nil {
			t.Log(shuffleErr)
			t.Fail()
		}
		for i := 0; i < 100; i++ {
			initialCard, _ := initialDeck.DrawCard()
			actualCard, drawErr := testDeck.DrawCard()
			if drawErr != nil {
				t.Log(drawErr)
				t.Fail()
			}
			cardComparisons++
			if initialCard.DisplayCard() == actualCard.DisplayCard() {
				coincidense++
			}
		}
	}
	if coincidense >= cardComparisons/2 {
		t.Log("warning, shuffling overlap exceeding expectations", coincidense, "/", repeatFor)
		t.Fail()
	}
}

func TestCombineShuffle(t *testing.T) {
	testDeck, _ := generateTestDeck()
	for testDeck.CardsLeft() > 0  {
		card, err := testDeck.DrawCard()
		if err != nil {
			t.Log("test overdrawing from test deck")
			t.Fail()
		}
		typeErr := testDeck.DiscardCard(card)
		if typeErr != nil {
			t.Log("type missmatch, cards are not same type as deck")
			t.Fail()
		}
	}
	
	testDeck.CombineShuffle()
	if testDeck.CardsLeft() != 100 {
		t.Log("deck size missmatch: expected 100 cards in deck, have", testDeck.CardsLeft())
		t.Fail()
	}
	if testDeck.CardsInPile() != 0 {
		t.Log("deck size missmatch: expected 0 cards in pile, have", testDeck.CardsInPile())
		t.Fail()
	}
}