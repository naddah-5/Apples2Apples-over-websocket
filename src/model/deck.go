package model

import (
	"bufio"
	"errors"
	"math/rand"
	"os"
	"strings"
)

type Deck struct {
	allowedCardType string
	deck []Card
	discard []Card
}

/*
Creates a Deck from a text file with the given type "deckType". 
A Deck contains both the deck and a discard pile for the deck.

Errors on incorrect file path.
*/
func GenerateDeck(source string, deckType string) (Deck, error) {
	f, fileErr := os.Open(source)
	if fileErr != nil {
		return Deck{}, fileErr
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)

	deck := createDeck(deckType)
	
	for scanner.Scan() { 
		cardData := strings.Split(scanner.Text(), " - ")
		formatDescription((&cardData[1])) // Remove parenthesis from description using a in-place function.

		newCard := MintCard(deckType, cardData[0], cardData[1])
		deck.deck = append(deck.deck, newCard)
	}
	return deck, nil
}

/*
Creates a new Deck with given type.
*/
func createDeck(deckType string) Deck {
	newDeck := *new(Deck)
	newDeck.setDeckType(deckType)
	return newDeck
}

/*
Removes the surrounding parenthesis from the description, note that there is a trailing white space at the end.
*/
func formatDescription(card *string) {
	*card = strings.Trim(*card, "(")
	*card = strings.Trim(*card, ") ")
}

/*
Returns how many cards are left in the shuffled deck as an int.
*/
func (d *Deck) CardsLeft() int {
	return len(d.deck)
}

/*
Returns how many cards there are in the discard pile as an int.
*/
func (d *Deck) CardsInPile() int {
	return len(d.discard)
}

/*
Draw the first card from the Deck.

Returns error if there are no cards in the deck
*/
func (d *Deck) DrawCard() (Card, error) {
	if len(d.deck) < 1 {
		return *new(Card), errors.New("deck is empty")
	}
	var card Card
	card, d.deck = d.deck[0], d.deck[1:]
	return card, nil
}

/*
Add a card to the discard pile, returns error if the card type 
does not match the pile type.
*/
func (d *Deck) DiscardCard(card Card) error {
	if card.cardType != d.allowedCardType {
		return errors.New("card type must match deck type")
	}
	d.discard = append(d.discard, card)
	return nil
}

/*
In-place shuffling of a deck, the discard pile will not be shuffled.
*/
func (d *Deck) ShuffleDeck() error {
	if len(d.deck) < 1 {
		return errors.New("can not shuffle a empty deck")
	}
	for j := 0; j < 7; j++ {
		for i := 0; i < len(d.deck); i++ {
			k := rand.Intn(len(d.deck))
			d.deck[i], d.deck[k] = d.deck[k], d.deck[i]
		}
	}
	return nil
}

/*
Append  the discard pile to the deck then shuffle it. 
*/
func (d *Deck) CombineShuffle() {
	d.deck = append(d.deck, d.discard...)
	d.discard = *new([]Card)
	d.ShuffleDeck()
}

/*
Set the type of cards that are accepted in the Deck.
*/
func (d *Deck) setDeckType(cardType string) error {
	// If a deck already contain cards the allowed types can not be changed.
	if (len(d.deck) > 0 || len(d.discard) > 0) {
		return errors.New("can not change type of non-empty deck")
	}
	d.allowedCardType = cardType
	return nil
}