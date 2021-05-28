package games

import (
	"ddz/cards"
	"ddz/players"
)

type Game struct {
	cardsBoot *cards.CardsBoot
	players   [3]*players.Player
	lordCards []*cards.Card
}

func NewGame() *Game {
	return &Game{
		cardsBoot: cards.NewCardsBoot(),
		players:   [3]*players.Player{},
		lordCards: []*cards.Card{},
	}
}
