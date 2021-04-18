package cards

type Card struct {
	Rank string
	Suit string
}

func NewCard(rank, suit string) *Card {
	return &Card{rank, suit}
}

func (card *Card) ToStr() string {
	return card.Rank + "/" + card.Suit
}
