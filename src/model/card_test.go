package model_test

import (
	"main/model"
	"testing"
)

func TestDisplayCard(t *testing.T) {
	card := model.MintCard("green apple", "[A card]", "What the card says.")
	cardView := card.DisplayCard()
	var expected string = "[A card] - What the card says."
	if cardView != expected {
		t.Fail()
	}
}

func TestCardType(t *testing.T) {
	card := model.MintCard("green apple", "[A card]", "What the card says.")
	cardType := card.CardType()
	var expectedType string = "green apple"
	if cardType != expectedType {
		t.Fail()
	}

}