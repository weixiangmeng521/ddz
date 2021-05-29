package cards

import (
	"fmt"
	"testing"
)

func TestInitCards(t *testing.T) {
	boots := NewCardsBoot()
	boots.Init()
	boots.Display()
}

func TestShuffleCards(t *testing.T) {
	boots := NewCardsBoot()
	boots.Init()
	boots.Shuffle()
	boots.Display()
}

func TestCardsIterator(t *testing.T) {
	boots := NewCardsBoot()
	boots.Init()
	boots.Shuffle()
	i := boots.Interator()

	cnt := 0
	for i.HasNext() {
		v := i.Next()
		fmt.Println(v.ToString())
		cnt++
	}
	fmt.Println(cnt)
}
