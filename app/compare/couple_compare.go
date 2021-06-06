package compare

import (
	c "ddz/app/cards"
	"ddz/app/constant"
)

// 是不是对子
var IsValidCouple = func(cards ...*c.Card) bool {
	if len(cards) != 2 {
		return false
	}
	if cards[0].Value == "Jack" || cards[1].Value == "Jack" {
		return false
	}
	return IsCardsEqual(cards...)
}

// 对子比较
type CoupleCards struct {
	cards   []*c.Card
	pattern constant.CardsPattern
}

func NewCoupleCards(cards ...*c.Card) *CoupleCards {
	if !IsValidCouple(cards...) {
		return nil
	}
	return &CoupleCards{
		cards:   cards,
		pattern: constant.CoupleCardsPattern,
	}
}

func (t *CoupleCards) GetCards() []*c.Card {
	return t.cards
}

func (t *CoupleCards) GetPattern() constant.CardsPattern {
	return t.pattern
}

func (t *CoupleCards) IsSamePattern(ci constant.CardsCompareInterface) bool {
	return t.GetPattern() == ci.GetPattern()
}

// 对子比较
func (t *CoupleCards) IsGreater(ci constant.CardsCompareInterface) bool {
	if len(ci.GetCards()) == 0 { // for any pattern
		return true
	}

	if ci.GetPattern() == constant.BoomCardsPattern {
		return false
	}

	cur := getSize(t.GetCards()[0].Value)
	tar := getSize(ci.GetCards()[0].Value)

	if cur == -1 || tar == -1 {
		return false
	}
	return cur > tar
}
