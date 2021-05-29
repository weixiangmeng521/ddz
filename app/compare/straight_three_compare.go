package compare

import (
	c "ddz/app/cards"
)

// 三顺
var IsValidStraightThree = func(cards ...*c.Card) bool {
	specialList := []*c.Card{
		c.NewCard("Jack", c.Freak),
		c.NewCard("Jack", c.Real),
		c.NewCard("2", c.Club),
		c.NewCard("2", c.Diamond),
		c.NewCard("2", c.Heart),
		c.NewCard("2", c.Spade),
	}
	// 三顺不应该小于6张牌
	if len(cards) < 6 {
		return false
	}
	// 三的倍数张牌
	if len(cards)%3 != 0 {
		return false
	}
	// 不应该有大小Jack & 2
	if HasContain(cards, specialList...) {
		return false
	}

	SortCards(cards)
	if !IsSequence(cards) {
		return false
	}

	return true
}

type StraightThreeCards struct {
	cards   []*c.Card
	pattern CardsPattern
}

func NewStraightThreeCards(cards ...*c.Card) *StraightThreeCards {
	if !IsValidStraightThree(cards...) {
		return nil
	}
	return &StraightThreeCards{
		cards:   MostSort(cards),
		pattern: StraightThreeCardsPattern,
	}
}

func (t *StraightThreeCards) GetCards() []*c.Card {
	return t.cards
}

func (t *StraightThreeCards) GetPattern() CardsPattern {
	return t.pattern
}

func (t *StraightThreeCards) IsSamePattern(ci CardsCompareInterface) bool {
	if t.GetPattern() != ci.GetPattern() { // 是不是姊妹对
		return false
	}
	if len(t.GetCards()) == len(ci.GetCards()) { // 姊妹对长度验证
		return false
	}
	return true
}

func (t *StraightThreeCards) IsGreater(ci CardsCompareInterface) bool {
	cur := getSize(t.GetCards()[0].Value)
	tar := getSize(ci.GetCards()[0].Value)
	if cur == -1 || tar == -1 {
		return false
	}
	return cur > tar
}
