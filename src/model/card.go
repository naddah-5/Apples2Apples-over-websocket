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
	return Card{
		cardType: cardType,
		header: header,
		description: description,
	}
}

/*
Returns a string representation of a card.
*/
func (c *Card) DisplayCard() string {
	return c.header + " - " + c.description
	
}

/*
Returns the cards type.
*/
func (c *Card) CardType() string {
	return c.cardType
}