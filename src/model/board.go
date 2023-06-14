package model

import (
	"errors"
	"fmt"
	"main/network"
	"math/rand"

	"main/view"
)

type Board struct {
	network network.Network
	players []Player
	judge int
	currentGreenApple Card
	redApples Deck
	greenApples Deck
	PlayedCards PlayedApples
	winCondition int
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

func (b *Board) CountPlayers() int {
	return len(b.players)
}

/*
Returns a pointer to the player matching the player name.

Returns an error if there are no matches.
*/
func (b *Board) findPlayer(playerName string) (*Player, error) {
	for i := 0; i < len(b.players); i++ {
		if b.players[i].PlayerName() == playerName {
			return &b.players[i], nil
		}
	}
	return new(Player), errors.New("player not found")
}


func (b *Board) PlayersHand(playerName string) ([]Card, error) {
	player, findErr := b.findPlayer(playerName)
	if findErr != nil {
		return *new([]Card), findErr
	}
	return player.PlayerHand(), nil
}

func (b *Board) AllHandsFull() bool {
	for i := 0; i < b.CountPlayers(); i++ {
		if len(b.players[i].hand) < b.players[i].handCapacity {
			return false
		}
	}
	return true
}

/*
Returns the string representation of players and their score.
*/
func (b *Board) DisplayScoreBoard() []string {
	var playerScores []string
	for i := 0; i < len(b.players); i++ {
		playerName := b.players[i].PlayerName()
		playerScore := b.players[i].Score()
		playerScores = append(playerScores, playerName + ": \t" + fmt.Sprint(playerScore))
	}
	return playerScores
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
func (b *Board) currentJudgeIndex() int {
	return b.judge
}

func (b *Board) CurrentJudgeName() string {
	return b.players[b.currentJudgeIndex()].PlayerName()
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
Awards the scoreCard to the player whos name matches playerName.

Returns an error if there is no player with that name.
*/
func (b *Board) AwardScore(playerName string, scoreCard Card) error {
	player, err := b.findPlayer(playerName)
	if err != nil {
		return err
	}
	player.IncreaseScore(scoreCard)
	return nil
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

/*
Draws a green apple and places it on the board.

Returns an error if a green apple can not be drawn.
*/
func (b *Board) DrawGreenApple() error {
	card, err := b.greenApples.DrawCard()
	if err != nil {
		return err
	}
	b.currentGreenApple = card
	return nil
}

/*
Returns the string representation of the current green apple on the board.
*/
func (b *Board) CurrentGreenApple() string {
	return b.currentGreenApple.DisplayCard()
}

/*
Retrieve the current green apple from the board.

Returns an error if there is no green apple on the board.
*/
func (b *Board) PickUpGreenApple() (Card, error) {
	card := b.currentGreenApple
	if card == (Card{}) {
		return card, errors.New("no green apple on board")
	}
	b.currentGreenApple = *new(Card)
	return card, nil
}

/*
Define the win condition 
*/
func (b *Board) SetWinCondition() error {
	if len(b.players) < 4 {
		return errors.New("not enough players")
	}
	b.winCondition = max(4, 12 - len(b.players))
	return nil
}

func max(a int, b int) int {
	if a >= b {
		return a
	}
	return b
}


/*
Returns the win condition, primarily usefull for testing.
*/
func (b *Board) GetWinCondition() int {
	return b.winCondition
}


/*
Check if any player satisfy the win condition.

Returns an error if the win condition is not set correctly.
*/
func (b *Board) GameWinner() (bool, error) {
	if b.winCondition <= 0 {
		return false, errors.New("invalid win condition")
	}
	for i := 0; i < len(b.players); i++ {
		if b.players[i].Score() >= b.winCondition {
			return true, nil
		}
	}
	return false, nil
}

/*
Returns a player that satisfy the win condition.

Returns an error if there is no player satisfying the win condition.
*/
func (b *Board) WhoWonGame() (Player, error) {
	for i := 0; i < len(b.players); i++ {
		if b.players[i].Score() >= b.winCondition {
			return b.players[i], nil
		}
	}
	return *new(Player), errors.New("there is no winner")
}

/*
Goes through all players and retrieves which card they want to play. Returns the 
PlayedApples struct containing all played cards.

Returns an error if a player plays an invalid card index, this should cause a panic.
Meaning that this method should not be used to validate user input.
*/
func (b *Board) ChooseCards() error {
	pa := new(PlayedApples)
	currentJudge := b.CurrentJudgeName()
	greenApple := b.CurrentGreenApple()
	for i := 0; i < len(b.players); i++ {
		if b.players[i].PlayerName() == currentJudge {
			// skips the judge
			continue
		}
		if b.players[i].Bot() {
			var randomCardIndex int = rand.Intn(b.players[i].CardsInHand())
			card, cardErr := b.players[i].PlayCard(randomCardIndex)
			if cardErr != nil {
				return errors.New("bot tried to play invalid card, " + cardErr.Error())
			}
			pa.SubmitCard(&b.players[i], card)
		}
		if b.players[i].Host() && !b.players[i].Bot() {
			hand := b.players[i].ShowHand()
			cardIndex := view.ChooseCard(greenApple, hand)
			card, playerCardErr := b.players[i].PlayCard(cardIndex)
			if playerCardErr != nil {
				return errors.New("host tried to play invalid card, " + playerCardErr.Error())
			}
			pa.SubmitCard(&b.players[i], card)
		}
		if !b.players[i].Host() && !b.players[i].Bot() {
			/*
			TODO: Add online call here!
			*/
		}
	}
	b.PlayedCards = *pa
	return nil
}

/*
Goes through all players and perferms the DrawCard(redApples) method on them, 
this will cause them to fill their hands to capacity with red apples.

Returns an error if a player is unable to draw cards from the deck.
*/
func (b *Board) FillHands() error {
	for i := 0; i < len(b.players); i++ {
		drawErr := b.players[i].DrawCard(&b.redApples)
		if drawErr != nil {
			return drawErr
		}
	}
	return nil
}

/*
Calls for the judge to decide the round winner.
The round winner is given by their index in the 
PlayersPlayed.pp struct.

Returns an error if no apples have been played.
*/
func (b *Board) Judge() (int, error) {	
	var greenApple string = b.CurrentGreenApple()
	currentJudge := b.players[b.currentJudgeIndex()]
	redApples, err := b.PlayedCards.DisplayApples()
	if err != nil {
		return -1, errors.New("no apples played")
	}
	
	if currentJudge.Bot() {
		return rand.Intn(len(b.PlayedCards.pp)), nil
	}
	if currentJudge.Host() && !currentJudge.Bot() {
		return view.JudgeCards(greenApple, redApples), nil
	}
	if !currentJudge.Host() && !currentJudge.Bot() {
		/*
		TODO: implement network call.
		*/
		return -1, errors.New("online play not implemented")
	}
	
	return -1, errors.New("unexpected judge status")
}