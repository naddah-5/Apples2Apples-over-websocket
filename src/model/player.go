package model

type Player struct {
	name string
	bot bool
	hand []Card
	points []Card
}

func NewPlayer(playerName string, bot bool) Player {
	player := Player{
		name: playerName,
		bot: bot,
		hand: *new([]Card),
		points: *new([]Card),
	}
	return player
}

func (p *Player) ID() string {
	return p.name
}

func (p *Player) Bot() bool {
	return p.bot
}