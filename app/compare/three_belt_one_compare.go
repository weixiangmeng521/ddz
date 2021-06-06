package compare

import (
	c "ddz/app/cards"
	"ddz/app/constant"
)

// 是不是合法的三带一
var IsValidThreeBeltOneCards = func(cards ...*c.Card) bool {
	if len(cards) != 4 {
		return false
	}
	res := MostSort(cards)
	if !IsCardsEqual(res[:3]...) {
		return false
	}
	return true
}

// 三带一
type ThreeBeltOneCards struct {
	cards   []*c.Card
	pattern constant.CardsPattern
}

func NewThreeBeltOneCards(cards ...*c.Card) *ThreeBeltOneCards {
	if !IsValidThreeBeltOneCards(cards...) {
		return nil
	}
	return &ThreeBeltOneCards{
		cards:   MostSort(cards),
		pattern: constant.ThreeBeltOneCardsPattern,
	}
}

func (t *ThreeBeltOneCards) GetCards() []*c.Card {
	return t.cards
}

func (t *ThreeBeltOneCards) GetPattern() constant.CardsPattern {
	return t.pattern
}

func (t *ThreeBeltOneCards) IsSamePattern(ci constant.CardsCompareInterface) bool {
	return t.GetPattern() == ci.GetPattern()
}

// 三带一比较大小
func (t *ThreeBeltOneCards) IsGreater(ci constant.CardsCompareInterface) bool {
	if len(ci.GetCards()) == 0 { // for any pattern
		return true
	}

	if ci.GetPattern() == constant.BoomCardsPattern {
		return false
	}

	cur := getSize(t.cards[0].Value)
	tar := getSize(ci.GetCards()[0].Value)
	if cur == -1 || tar == -1 {
		return false
	}
	return cur > tar
}
