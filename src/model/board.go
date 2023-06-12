package model

import (
	"errors"
	"main/network"
	"math/rand"
)

type Board struct {
	network network.Network
	players []Player
	judge int
	redApples Deck
	greenApples Deck
	playedCards []PlayedApples
}

/*
Adds a player to the board unless the player name is 
already in use.

Returns an error for name collisions.
*/
func (b *Board) AddPlayer(player Player) error {
	if b.validateName(player.PlayerName()) {
		b.players = append(b.players, player)
		return nil
	}
	return errors.New("name is unavailable")
}

/*
Checks if a name is available, if it is available returns 
true. Otherwise returns false.
*/
func (b *Board) validateName(name string) bool {
	for i := 0; i < len(b.players); i++ {
		if name == b.players[i].PlayerName() {
			return false
		}
	}
	return true
}

/*
Shuffle the player order. 

Returns an error if there are no players.

This method is usefull for randomizing the order of players before 
the round robin starts.
*/
func (b *Board) ShufflePlayers() error {
	if len(b.players) == 0 {
		return errors.New("can not shuffle zero players")
	}
	for i := 0; i < 7; i++ {
		for j := 0; j < len(b.players); j++ {
			k := rand.Intn(len(b.players))
			b.players[j], b.players[k] = b.players[k], b.players[j]
		}
	}
	return nil
}

/*
Returns current judge index.
*/
func (b *Board) CurrentJudge() int {
	return b.judge
}

/*
Itterates to a new judge and return the new judge index.
*/
func (b *Board) ItterateJudge() int {
	b.judge++
	if b.judge >= len(b.players) {
		b.judge = 0
	}
	return b.judge
}

/*
Loads a deck of red apples from a resource file.

Returns an error if the file path is incorrect.
*/
func (b *Board) LoadRedApples(source string) error {
	redAppleDeck, deckErr := GenerateDeck(source, "red apple")
	if deckErr != nil {
		return deckErr
	}
	b.redApples = redAppleDeck
	return nil
}

/*
Loads a deck of green apples from a resource file.

Returns an error if the file path is incorrect.
*/
func (b *Board) LoadGreenApples(source string) error {
	greenAppleDeck, deckErr := GenerateDeck(source, "green apple")
	if deckErr != nil {
		return deckErr
	}
	b.greenApples = greenAppleDeck
	return nil
}


