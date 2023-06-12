package model

import (
	"errors"
	"main/network"
	"math/rand"
)

type Game struct {
	network network.Network
	players []Player
	judge int
	redApples Deck
	greenApples Deck
	playedCards []PlayedApples
}

/*
Adds a player to the game unless the player name is 
already in use, for which it returns an error.
*/
func (g *Game) AddPlayer(player Player) error {
	if g.validateName(player.PlayerName()) {
		g.players = append(g.players, player)
		return nil
	}
	return errors.New("name is unavailable")
}

/*
Checks if a name is available, if it is available returns 
true. Otherwise returns false.
*/
func (g *Game) validateName(name string) bool {
	for i := 0; i < len(g.players); i++ {
		if name == g.players[i].PlayerName() {
			return false
		}
	}
	return true
}

/*
Shuffle the player order, returns an error if there are no players.
This method is usefull for randomizing the order of players before 
the round robin starts.
*/
func (g *Game) ShufflePlayers() error {
	if len(g.players) == 0 {
		return errors.New("can not shuffle zero players")
	}
	for i := 0; i < 7; i++ {
		for j := 0; j < len(g.players); j++ {
			k := rand.Intn(len(g.players))
			g.players[j], g.players[k] = g.players[k], g.players[j]
		}
	}
	return nil
}

/*
Returns current judge index.
*/
func (g *Game) CurrentJudge() int {
	return g.judge
}

/*
Itterates to a new judge and return the new judge index.
*/
func (g *Game) ItterateJudge() int {
	g.judge++
	if g.judge >= len(g.players) {
		g.judge = 0
	}
	return g.judge
}

/*
Loads a deck of red apples from a resource file, returns an 
error if the file path is incorrect.
*/
func (g *Game) LoadRedApples(source string) error {
	redAppleDeck, deckErr := GenerateDeck(source, "red apple")
	if deckErr != nil {
		return deckErr
	}
	g.redApples = redAppleDeck
	return nil
}

/*
Loads a deck of green apples from a resource file, returns an 
error if the file path is incorrect.
*/
func (g *Game) LoadGreenApples(source string) error {
	greenAppleDeck, deckErr := GenerateDeck(source, "green apple")
	if deckErr != nil {
		return deckErr
	}
	g.greenApples = greenAppleDeck
	return nil
}


