package compare

import c "ddz/app/cards"

// 是不是合法的三带一
var IsValidThreeBeltTwoCards = func(cards ...*c.Card) bool {
	if len(cards) != 5 {
		return false
	}
	res := MostSort(cards)
	// 前面三张牌相等
	if !IsCardsEqual(res[:3]...) {
		return false
	}
	// 后面两张牌相等
	if !IsCardsEqual(res[3:]...) {
		return false
	}
	return true
}

// 三带二
type ThreeBeltTwoCards struct {
	cards   []*c.Card
	pattern CardsPattern
}

func NewThreeBeltTwoCards(cards ...*c.Card) *ThreeBeltTwoCards {
	if !IsValidThreeBeltTwoCards(cards...) {
		return nil
	}
	return &ThreeBeltTwoCards{
		cards:   MostSort(cards),
		pattern: ThreeBeltTwoCardsPattern,
	}
}

func (t *ThreeBeltTwoCards) GetCards() []*c.Card {
	return t.cards
}

func (t *ThreeBeltTwoCards) GetPattern() CardsPattern {
	return t.pattern
}

func (t *ThreeBeltTwoCards) IsSamePattern(ci CardsCompareInterface) bool {
	return t.GetPattern() == ci.GetPattern()
}

// 三带二比较大小
func (t *ThreeBeltTwoCards) IsGreater(ci CardsCompareInterface) bool {
	cur := getSize(t.cards[0].Value)
	tar := getSize(ci.GetCards()[0].Value)
	if cur == -1 || tar == -1 {
		return false
	}
	return cur > tar
}
