package compare

import c "ddz/cards"

// 是不是炸弹
var IsValidBoom = func(cards ...*c.Card) bool {
	// 双王炸弹
	if len(cards) == 2 && cards[0].Value == "Jack" && cards[1].Value == "Jack" {
		return true
	}
	// 四张牌的炸弹
	if len(cards) != 4 {
		return false
	}
	for i := 1; i < 4; i++ {
		if cards[i].Value != cards[i-1].Value {
			return false
		}
	}
	return true
}

// 炸弹比较
type BoomCards struct {
	cards   []*c.Card
	pattern CardsPattern
}

func NewBoomCards(cards ...*c.Card) *BoomCards {
	if !IsValidBoom(cards...) {
		return nil
	}
	return &BoomCards{
		cards:   cards,
		pattern: BoomCardsPattern,
	}
}

func (t *BoomCards) GetCards() []*c.Card {
	return t.cards
}

func (t *BoomCards) GetPattern() CardsPattern {
	return t.pattern
}

func (t *BoomCards) IsSamePattern(ci CardsCompareInterface) bool {
	return t.GetPattern() == ci.GetPattern()
}

// 炸弹比较
func (t *BoomCards) IsGreater(ci CardsCompareInterface) bool {
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
