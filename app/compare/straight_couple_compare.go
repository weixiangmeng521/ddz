package compare

import (
	c "ddz/app/cards"
)

// 是不是合法的姊妹对
var IsValidStraightCouple = func(cards ...*c.Card) bool {
	specialList := []*c.Card{
		c.NewCard("Jack", c.Freak),
		c.NewCard("Jack", c.Real),
		c.NewCard("2", c.Club),
		c.NewCard("2", c.Diamond),
		c.NewCard("2", c.Heart),
		c.NewCard("2", c.Spade),
	}
	// 姊妹对不应该小于6张牌
	if len(cards) < 6 {
		return false
	}
	// 是双数
	if len(cards)%2 != 0 {
		return false
	}

	// 不应该有大小Jack & 2
	if HasContain(cards, specialList...) {
		return false
	}

	SortCards(cards)

	sequence := []int{}
	// 检查是不是aabb格式
	for i := 0; i < len(cards)-1; i += 2 {
		slice := cards[i : i+2]
		if !IsCardsEqual(slice...) {
			return false
		}
		sequence = append(sequence, getSize(slice[0].Value))
	}
	// 检查是不是序列
	for i := 1; i < len(sequence)-1; i++ {
		if sequence[i-1]+1 != sequence[i] {
			return false
		}
	}

	return true
}

type StraightCoupleCards struct {
	cards   []*c.Card
	pattern CardsPattern
}

func NewStraightCoupleCards(cards ...*c.Card) *StraightCoupleCards {
	if !IsValidStraightCouple(cards...) {
		return nil
	}
	return &StraightCoupleCards{
		cards:   SortCards(cards),
		pattern: StraightCoupleCardsPattern,
	}
}

func (t *StraightCoupleCards) GetCards() []*c.Card {
	return t.cards
}

func (t *StraightCoupleCards) GetPattern() CardsPattern {
	return t.pattern
}

func (t *StraightCoupleCards) IsSamePattern(ci CardsCompareInterface) bool {
	if t.GetPattern() != ci.GetPattern() { // 是不是姊妹对
		return false
	}
	if len(t.GetCards()) == len(ci.GetCards()) { // 姊妹对长度验证
		return false
	}
	return true
}

func (t *StraightCoupleCards) IsGreater(ci CardsCompareInterface) bool {
	cur := getSize(t.GetCards()[0].Value)
	tar := getSize(ci.GetCards()[0].Value)
	if cur == -1 || tar == -1 {
		return false
	}
	return cur > tar
}
