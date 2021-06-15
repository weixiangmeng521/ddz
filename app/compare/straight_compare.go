package compare

import (
	c "ddz/app/cards"
	"ddz/app/constant"
)

// 是不是合法的顺子
var IsValidStraight = func(cards ...*c.Card) bool {
	specialList := []*c.Card{
		c.NewCard("Jack", c.Freak),
		c.NewCard("Jack", c.Real),
		c.NewCard("2", c.Club),
		c.NewCard("2", c.Diamond),
		c.NewCard("2", c.Heart),
		c.NewCard("2", c.Spade),
	}
	// 顺子不应该小于5张牌
	if len(cards) < 5 {
		return false
	}
	// 不应该有大小Jack & 2
	if HasContain(cards, specialList...) {
		return false
	}

	SortCards(cards)

	// 检查是否重复 && 是不是序列
	for i := 1; i < len(cards); i++ {
		if cards[i-1].Value == cards[i].Value {
			return false
		}
		if getSize(cards[i-1].Value)+1 != getSize(cards[i].Value) {
			return false
		}
	}
	return true
}

// 顺子
type StraightCards struct {
	cards   []*c.Card
	pattern constant.CardsPattern
}

func NewStraightCards(cards ...*c.Card) *StraightCards {
	if !IsValidStraight(cards...) {
		return nil
	}
	return &StraightCards{
		cards:   SortCards(cards),
		pattern: constant.StraightCardsPattern,
	}
}

func (t *StraightCards) GetCards() []*c.Card {
	return t.cards
}

func (t *StraightCards) GetPattern() constant.CardsPattern {
	return t.pattern
}

func (t *StraightCards) IsSamePattern(ci constant.CardsCompareInterface) bool {
	if t.GetPattern() != ci.GetPattern() { // 是不是顺子
		return false
	}
	if len(t.GetCards()) == len(ci.GetCards()) { // 顺子长度验证
		return false
	}
	return true
}

func (t *StraightCards) IsGreater(ci constant.CardsCompareInterface) bool {
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
