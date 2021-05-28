package players

import (
	"ddz/cards"
	"fmt"
	"testing"
)

func TestDealCards(t *testing.T) {
	p := NewPlayer()
	cardsList := []*cards.Card{
		cards.NewCard("3", cards.Spade),
		cards.NewCard("4", cards.Spade),
		cards.NewCard("5", cards.Spade),
		cards.NewCard("6", cards.Spade),
		cards.NewCard("7", cards.Spade),
		cards.NewCard("8", cards.Spade),
		cards.NewCard("9", cards.Spade),
		cards.NewCard("10", cards.Spade),
		cards.NewCard("J", cards.Spade),
		cards.NewCard("Q", cards.Spade),
		cards.NewCard("K", cards.Spade),
	}
	p.AcceptCards(cardsList...)
	fmt.Println(p.CheckCards())

	p.DealCard(cards.NewCard("3", cards.Spade))
	fmt.Println(p.CheckCards())

	dealList := []*cards.Card{
		cards.NewCard("7", cards.Spade),
		cards.NewCard("8", cards.Spade),
		cards.NewCard("9", cards.Spade),
	}
	fmt.Println(p.DealCards(dealList...))
	fmt.Println(p.CheckCards())
}
