package model

import "errors"

type Player struct {
	name string
	bot bool
	hand []Card
	handCapacity int
	points []Card
}

/*
Creates and returns a new player.
*/
func NewPlayer(playerName string, bot bool) Player {
	player := Player{
		name: playerName,
		bot: bot,
		hand: *new([]Card),
		handCapacity: 7,
		points: *new([]Card),
	}
	return player
}

/*
Returns the players name.
*/
func (p *Player) PlayerName() string {
	return p.name
}

/*
Returns true if the player is a bot.
*/
func (p *Player) Bot() bool {
	return p.bot
}

/*
Player draws card until hand is full. If deck runs out before that, 
attempts to shuffle the discar pile and continue drawing cards.

Returns error if there is not enough cards to draw in the deck.
*/
func (p *Player) DrawCard(deck *Deck) error {
	for len(p.hand) < p.handCapacity {
		newCard, drawErr := deck.DrawCard()
		if drawErr != nil {
			if deck.CardsInPile() == 0 {
				return errors.New("not enough cards in deck")
			}
			deck.CombineShuffle()
			newCard, drawErr = deck.DrawCard()
		}
		p.hand = append(p.hand, newCard)
	}
	return nil
}

/*
Returns a card from the player hands, and removes it from that players hand.

Returns error if a invalid index is given.
*/
func (p *Player) PlayCard(index int) (Card, error) {
	if index < 0 || len(p.hand) < index {
		return *new(Card), errors.New("invalid card index")
	}
	card := p.hand[index]
	var boundedCut int = max((index-1), 0)
	p.hand = append(p.hand[:boundedCut], p.hand[index:]...)
	return card, nil
}

/*
Returns the larger of two integers, if they are the same size 
the first argument is returned.
*/
func max(a int, b int) int {
	if a >= b {
		return a
	}
	return b
}

/*
Adds the given card to player points and returns the players 
score using the Score() function.
*/
func (p *Player) IncreaseScore(card Card) int {
	p.hand = append(p.hand, card)
	return p.Score()
}

/*
Returns the length of the players points, which 
is the current scoring system.
*/
func (p *Player) Score() int {
	return len(p.points)
}

/*
Returns a COPY of the players hand, do not discard these cards into a pile.
*/
func (p *Player) PlayerHand() []Card {
	return p.hand
}