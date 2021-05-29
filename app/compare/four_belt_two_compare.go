package compare

import c "ddz/app/cards"

// 四张带两张：四张一样的牌+两张单牌。(注意：四带二不是炸弹)。如：4444+65
var IsValidFourBeltTwo = func(cards ...*c.Card) bool {
	// 必须6张牌
	if len(cards) != 6 {
		return false
	}

	// 四代二排序
	set := map[string][]*c.Card{}
	for _, card := range cards {
		if _, ok := set[card.Value]; !ok {
			set[card.Value] = []*c.Card{card}
			continue
		}
		set[card.Value] = append(set[card.Value], card)
	}

	arr := [][]*c.Card{}
	for _, list := range set {
		arr = append(arr, list)
	}

	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			tmp := arr[i]
			if len(arr[i]) < len(arr[j]) {
				arr[i] = arr[j]
				arr[j] = tmp
			}
		}
	}

	res := []*c.Card{}
	for _, v := range arr {
		res = append(res, v...)
	}

	// 前面四个要相同
	if !IsCardsEqual(res[:4]...) {
		return false
	}

	return true
}

type FourBeltTwoCards struct {
	cards   []*c.Card
	pattern CardsPattern
}

func NewFourBeltTwoCards(cards ...*c.Card) *FourBeltTwoCards {
	if !IsValidFourBeltTwo(cards...) {
		return nil
	}
	return &FourBeltTwoCards{
		cards:   MostSort(cards),
		pattern: FourBeltTwoPattern,
	}
}

func (t *FourBeltTwoCards) GetCards() []*c.Card {
	return t.cards
}

func (t *FourBeltTwoCards) GetPattern() CardsPattern {
	return t.pattern
}

func (t *FourBeltTwoCards) IsSamePattern(ci CardsCompareInterface) bool {
	if t.GetPattern() != ci.GetPattern() { // 是不是姊妹对
		return false
	}
	if len(t.GetCards()) == len(ci.GetCards()) { // 姊妹对长度验证
		return false
	}
	return true
}

func (t *FourBeltTwoCards) IsGreater(ci CardsCompareInterface) bool {
	cur := getSize(t.GetCards()[0].Value)
	tar := getSize(ci.GetCards()[0].Value)
	if cur == -1 || tar == -1 {
		return false
	}
	return cur > tar
}
