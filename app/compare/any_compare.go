package compare

import (
	c "ddz/app/cards"
	"ddz/app/constant"
)

// 是不是任意牌？ps: 这个牌只能系统判定
var IsValidAny = func(cards ...*c.Card) bool {
	return false
}

// any比较
type AnyCards struct {
	cards   []*c.Card
	pattern constant.CardsPattern
}

func NewAnyCards() *AnyCards {
	return &AnyCards{
		cards:   []*c.Card{},
		pattern: constant.AnyPattern,
	}
}

func (t *AnyCards) GetCards() []*c.Card {
	return t.cards
}

func (t *AnyCards) GetPattern() constant.CardsPattern {
	return t.pattern
}

func (t *AnyCards) IsSamePattern(ci constant.CardsCompareInterface) bool {
	return true
}

func (t *AnyCards) IsGreater(ci constant.CardsCompareInterface) bool {
	return false
}
