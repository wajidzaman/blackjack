package main

import (
	cards "blackjack/cards"
	"fmt"
	"testing"
)

var ranks = map[string]bool{
	"A":  true,
	"2":  true,
	"3":  true,
	"4":  true,
	"5":  true,
	"6":  true,
	"7":  true,
	"8":  true,
	"9":  true,
	"10": true,
	"J":  true,
	"Q":  true,
	"K":  true,
}
var suits = map[string]bool{
	"heart":   true,
	"spade":   true,
	"club":    true,
	"diamond": true,
}

func TestDeck(t *testing.T) {
	deck := cards.NewDeck()
	var hand [52]*cards.Card
	fmt.Println("cards", len(deck.Cards))
	for i := 0; i < 52; i++ {

		hand[i] = deck.DealCard()

		if !ranks[hand[i].Rank] {
			t.Errorf("Unexpected rank %s", hand[i].Rank)
		}

		if !suits[hand[i].Suit] {
			t.Errorf("Unexpected suit %s", hand[i].Suit)
		}
	}

	if deck.DealCard() != nil {
		t.Errorf("Expected nil after dealing all 52 cards")
	}

}
