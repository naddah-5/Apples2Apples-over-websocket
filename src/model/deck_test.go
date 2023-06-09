package model_test

import (
	"errors"
	"fmt"
	"main/model"
	"path/filepath"
	"testing"
)

func generateTestDeck() (model.Deck, error) {
	absPath, pathErr := filepath.Abs("../resources/greenApples.txt")
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
		t.FailNow()
	}
	var expectedCard string = "[Absurd] - ridiculous, senseless, foolish"
	card, emptyDeckErr := testDeck.DrawCard()
	if emptyDeckErr != nil {
		t.Log(emptyDeckErr)
		t.FailNow()
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
		t.FailNow()
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
}
func TestDiscardCardInvalid(t *testing.T) {
	testDeck, deckErr := generateTestDeck()
	if deckErr != nil {
		t.Log(deckErr)
		t.FailNow()
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
	var coincidense int = 0
	var confidence int = 10_000
	for n := 0; n < confidence; n++ {
		testDeck, deckGenErr := generateTestDeck()
		initialDeck, _ := generateTestDeck()
		if deckGenErr != nil {
			t.Log(deckGenErr)
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
			if initialCard == actualCard {
				coincidense++
			}
		}
	}
	if coincidense >= confidence/2 {
		t.Log("warning, shuffling overlap exceeding expectations")
		t.Fail()
	}
}