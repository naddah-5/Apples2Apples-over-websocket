package model

type GameModel struct {
	online string
	players []Player
	judge int
	redApples Deck
	greenApples Deck
	playedCards []Card
}