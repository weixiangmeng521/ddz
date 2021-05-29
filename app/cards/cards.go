package cards

import (
	"fmt"
	"strings"
)

type CardsList struct {
	list []*Card
}

func NewCardsList(cards ...*Card) *CardsList {
	return &CardsList{cards}
}

func (t *CardsList) ToString() string {
	arr := []string{}
	for _, card := range t.list {
		arr = append(arr, card.ToString())
	}
	return strings.Join(arr, ",")
}

func (t *CardsList) Display() {
	fmt.Println(t.ToString())
}
