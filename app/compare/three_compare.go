package compare

import (
	c "ddz/app/cards"
)

// 是不是合法三不带
var IsValidThreeCards = func(cards ...*c.Card) bool {
	if len(cards) != 3 {
		return false
	}
	if !IsCardsEqual(cards...) {
		return false
	}
	return true
}

// 三不带
type ThreeCards struct {
	pattern CardsPattern
	cards   []*c.Card
}

func NewThreeCards(cards ...*c.Card) *ThreeCards {
	if !IsValidThreeCards(cards...) {
		return nil
	}
	return &ThreeCards{
		cards:   cards,
		pattern: ThreeCardsPattern,
	}
}

func (t *ThreeCards) GetCards() []*c.Card {
	return t.cards
}

func (t *ThreeCards) GetPattern() CardsPattern {
	return t.pattern
}

func (t *ThreeCards) IsSamePattern(ci CardsCompareInterface) bool {
	return t.GetPattern() == ci.GetPattern()
}

// 三不带比较大小
func (t *ThreeCards) IsGreater(ci CardsCompareInterface) bool {
	cur := getSize(t.cards[0].Value)
	tar := getSize(ci.GetCards()[0].Value)
	if cur == -1 || tar == -1 {
		return false
	}
	return cur > tar
}
