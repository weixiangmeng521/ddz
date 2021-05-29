package cards

import (
	"fmt"
	"testing"
)

func TestCardsIsEqual(t *testing.T) {
	card := NewCard("3", Spade)
	res := card.IsEqual(NewCard("3", Spade))
	fmt.Println(res)
}
