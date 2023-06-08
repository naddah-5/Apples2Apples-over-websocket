package model_test

import (
	"fmt"
	"main/model"
	"path/filepath"
	"testing"
)

func TestGeneratePile(t *testing.T) {
	absPath, pathErr := filepath.Abs("../resources/greenApples.txt")
	if pathErr != nil {
		t.FailNow()
		t.Log("test incorrectly configured")
	}
	testDeck, err := model.GenerateDeck(absPath, "green apple")
	if err != nil {
		t.FailNow()
	}
	var expectedCard string = "[Absurd] - ridiculous, senseless, foolish"
	card := testDeck.DrawCard()
	if card.DisplayCard() != expectedCard {
		fmt.Println("received ", card.DisplayCard())
		fmt.Println("expected ", expectedCard)
		t.Fail()
	}
}