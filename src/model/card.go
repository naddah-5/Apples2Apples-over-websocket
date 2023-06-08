package model

type Card struct {
	cardType string
	header string
	description string
}


/*
Creates and returns a new card struct.
*/
func MintCard(cardType string, header string, description string) Card {
	var newCard Card = Card{
		cardType: cardType,
		header: header,
		description: description,
	}
	return newCard
}

/*
Returns a string representation of a card.
*/
func (c *Card) DisplayCard() string {
	var text string = c.header + " - " + c.description
	return text
}

/*
Returns the cards type.
*/
func (c *Card) CardType() string {
	return c.cardType
}