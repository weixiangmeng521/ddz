package compare

import (
	c "ddz/app/cards"
	"ddz/app/constant"
)

// 是不是单张
var IsValidSingle = func(cards ...*c.Card) bool {
	return len(cards) == 1
}

// 单张比较
type SingleCard struct {
	cards   []*c.Card
	pattern constant.CardsPattern
}

func NewSingleCard(cards ...*c.Card) *SingleCard {
	if !IsValidSingle(cards...) {
		return nil
	}
	return &SingleCard{
		cards:   cards,
		pattern: constant.SinglePattren,
	}
}

func (t *SingleCard) GetCards() []*c.Card {
	return t.cards
}

func (t *SingleCard) GetPattern() constant.CardsPattern {
	return t.pattern
}

func (t *SingleCard) IsSamePattern(ci constant.CardsCompareInterface) bool {
	return t.GetPattern() == ci.GetPattern()
}

func (t *SingleCard) IsGreater(ci constant.CardsCompareInterface) bool {
	if len(ci.GetCards()) == 0 { // for any pattern
		return true
	}

	if ci.GetPattern() == constant.BoomCardsPattern {
		return false
	}

	// 大小马比较
	if t.GetCards()[0].Value == "Jack" && ci.GetCards()[0].Value == "Jack" {
		return t.GetCards()[0].Type == c.Real && ci.GetCards()[0].Type == c.Freak
	}
	// 一般的牌比较
	cur := getSize(t.cards[0].Value)
	tar := getSize(ci.GetCards()[0].Value)
	if cur == -1 || tar == -1 {
		return false
	}
	return cur > tar
}
