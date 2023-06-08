package model

import (
	"bufio"
	"os"
	"strings"
)

type CardPile struct {
	deck []Card
	discard []Card
}

/*
Creates a new CardPile with initial capacity of 200 cards.
*/
func CreatePile() CardPile {
	newPile := *new(CardPile)
	return newPile
}

/*
Draw the first card from the CardPile.
*/
func (p *CardPile) DrawCard() Card {
	var card Card
	card, p.deck = p.deck[0], p.deck[1:]
	return card
}

func GenerateDeck(source string, deckType string) (CardPile, error) {
	f, fileErr := os.Open(source)
	if fileErr != nil {
		return CardPile{}, fileErr
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	pile := CreatePile()
	
	for scanner.Scan() { 
		card := strings.Split(scanner.Text(), " - ")
		formatDescription((&card[1])) // Remove parenthesis from description using a pointer function.

		newCard := MintCard(deckType, card[0], card[1])
		pile.deck = append(pile.deck, newCard)
	}
	return pile, nil
}

/*
Removes the surrounding parenthesis from the description, note that there is a trailing white space at the end.
*/
func formatDescription(card *string) {
	*card = strings.Trim(*card, "(")
	*card = strings.Trim(*card, ") ")
}