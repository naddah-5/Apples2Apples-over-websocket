package model

import (
	"errors"
	"math/rand"
)

type PlayedApples struct {
	pp []PlayerPlayed
}

type PlayerPlayed struct{
	players *Player
	card Card
}


/*
Shuffles the order of the submitted cards.
Returns an error if there are no submissions.
*/
func (pa *PlayedApples) Shuffle() error {
	if len(pa.pp) == 0 {
		return errors.New("can not shuffle zero players")
	}
	for i := 0; i < 7; i++ {
		for j := 0; j < len(pa.pp); j++ {
			k := rand.Intn(len(pa.pp))
			pa.pp[j], pa.pp[k] = pa.pp[k], pa.pp[j]
		}
	}
	return nil
}

/*
Returns the display card results for all played cards, in order.
Returns error if there are no cards in the struct.
*/
func (pa *PlayedApples) DisplayApples() ([]string, error) {
	if len(pa.pp) == 0 {
		return *new([]string), errors.New("no cards played")
	}
	var apples []string
	for i := 0; i < len(pa.pp); i++ {
		apples = append(apples, pa.pp[i].card.DisplayCard())
	}
	return apples, nil
}

/*
Returns the player name of the chosen index, usefull for showing who 
won the round when the judge chooses a winning card.
*/
func (pa *PlayedApples) ShowPlayer(index int) (string, error) {
	if index < 0 || index >= len(pa.pp) {
		return "", errors.New("index out of bounds")
	}
	return pa.pp[index].players.PlayerName(), nil
}

/*
Discards all cards in PlayedApples and replace them with empty cards 
so that they are not accidentally used again.
Returns an error if the injected deck does not match the card type.
*/
func (pa *PlayedApples) DiscardRound(deck *Deck) error {
	for i := 0; i < len(pa.pp); i++ {
		disErr := deck.DiscardCard(pa.pp[i].card)
		if disErr != nil {
			return disErr
		}
		pa.pp[i].card = *new(Card)
	}
	return nil
}