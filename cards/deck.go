package cards

import (
	"math/rand"
)

var ranks = []string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}
var suits = []string{"spade", "diamond", "heart", "club"}

type Deck struct {
	Pile
}

func NewDeck() *Deck {
	var deck *Deck = new(Deck)
	for _, rank := range ranks {
		for _, suit := range suits {
			deck.PutDown(NewCard(rank, suit))
		}
	}
	deck.shuffle()
	return deck

}
func (deck *Deck) shuffle() {
	for i := 0; i < len(deck.Cards); i++ {

		j := rand.Intn(len(deck.Cards) - 1)
		deck.Cards[i], deck.Cards[j] = deck.Cards[j], deck.Cards[i]
	}
}
func (deck *Deck) DealCard() *Card {
	if len(deck.Cards) == 0 {
		return nil
	}
	card := deck.Cards[len(deck.Cards)-1]
	deck.Cards = deck.Cards[:len(deck.Cards)-1]
	return card
}
