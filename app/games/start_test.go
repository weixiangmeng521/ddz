package games

import (
	c "ddz/app/cards"
	"ddz/app/players"
	"fmt"
	"testing"
)

func TestNewGame(t *testing.T) {
	g := NewGame("test")
	g.JoinPlayer(players.NewPlayer("Kenny"))
	g.JoinPlayer(players.NewPlayer("Kyle"))
	g.JoinPlayer(players.NewPlayer("Cartman"))

	if !g.CanStart() {
		t.Error()
	}
	g.Shuffle()
	g.Licensing()
	g.Display()
}

func TestGetCardsPattern(t *testing.T) {
	cards := []*c.Card{
		c.NewCard("J", c.Club),
		c.NewCard("J", c.Club),
		c.NewCard("J", c.Club),
		c.NewCard("Q", c.Club),
		c.NewCard("Q", c.Club),
		c.NewCard("Q", c.Club),
	}
	res := GetCardsPattern(cards...)
	fmt.Println(res.ToString())
}

func TestConvertCards(t *testing.T) {
	cards := []*c.Card{
		c.NewCard("J", c.Club),
		c.NewCard("Q", c.Club),
		c.NewCard("Q", c.Club),
		c.NewCard("J", c.Club),
		c.NewCard("J", c.Club),
		c.NewCard("Q", c.Club),
	}
	pattern := GetCardsPattern(cards...)
	obj := ConvertCards(pattern, cards)
	c.NewCardsList(obj.GetCards()...).Display()
}

func TestStart(t *testing.T) {
	Start()
}
