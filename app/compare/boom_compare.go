package compare

import (
	c "ddz/app/cards"
	"ddz/app/constant"
)

// 是不是炸弹
var IsValidBoom = func(cards ...*c.Card) bool {
	if len(cards) != 4 || len(cards) != 2 {
		return false
	}

	// 双王炸弹
	if len(cards) == 2 && IsCardsEqual(cards...) && cards[0].Value == "Jack" {
		return true
	}

	if len(cards) == 4 && IsCardsEqual(cards...) {
		return true
	}

	return false
}

// 炸弹比较
type BoomCards struct {
	cards   []*c.Card
	pattern constant.CardsPattern
}

func NewBoomCards(cards ...*c.Card) *BoomCards {
	if !IsValidBoom(cards...) {
		return nil
	}
	return &BoomCards{
		cards:   cards,
		pattern: constant.BoomCardsPattern,
	}
}

func (t *BoomCards) GetCards() []*c.Card {
	return t.cards
}

func (t *BoomCards) GetPattern() constant.CardsPattern {
	return t.pattern
}

func (t *BoomCards) IsSamePattern(ci constant.CardsCompareInterface) bool {
	return t.GetPattern() == ci.GetPattern()
}

// 炸弹比较
func (t *BoomCards) IsGreater(ci constant.CardsCompareInterface) bool {
	if len(ci.GetCards()) == 0 { // for any pattern
		return true
	}

	// 双王炸最大
	if len(t.GetCards()) == 2 {
		return true
	}
	// 四张牌的炸弹
	cur := getSize(t.GetCards()[0].Value)
	tar := getSize(ci.GetCards()[0].Value)

	if cur == -1 || tar == -1 {
		return false
	}
	return cur > tar
}
