package compare

import (
	c "ddz/app/cards"
)

// 是不是对子
var IsValidCouple = func(cards ...*c.Card) bool {
	if len(cards) != 2 {
		return false
	}
	return cards[0].Value == cards[1].Value
}

// 对子比较
type CoupleCards struct {
	cards   []*c.Card
	pattern CardsPattern
}

func NewCoupleCards(cards ...*c.Card) *CoupleCards {
	if !IsValidCouple(cards...) {
		return nil
	}
	return &CoupleCards{
		cards:   cards,
		pattern: CoupleCardsPattern,
	}
}

func (t *CoupleCards) GetCards() []*c.Card {
	return t.cards
}

func (t *CoupleCards) GetPattern() CardsPattern {
	return t.pattern
}

func (t *CoupleCards) IsSamePattern(ci CardsCompareInterface) bool {
	return t.GetPattern() == ci.GetPattern()
}

// 对子比较
func (t *CoupleCards) IsGreater(ci CardsCompareInterface) bool {
	cur := getSize(t.GetCards()[0].Value)
	tar := getSize(ci.GetCards()[0].Value)
	if cur == -1 || tar == -1 {
		return false
	}
	return cur > tar
}
