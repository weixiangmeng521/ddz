package compare

import (
	c "ddz/app/cards"
)

// 是不是飞机
// 三顺 + 同数量的单牌（或同数量的对牌）
var IsValidAirplane = func(cards ...*c.Card) bool {
	// 给牌归类
	set := map[string][]*c.Card{}
	for _, card := range cards {
		if _, ok := set[card.Value]; !ok {
			set[card.Value] = []*c.Card{card}
			continue
		}
		set[card.Value] = append(set[card.Value], card)
	}
	base := [][]*c.Card{}        // 基础牌
	extraSingle := [][]*c.Card{} // 带的单张
	extraCouple := [][]*c.Card{} // 带的对子
	for _, v := range set {
		if len(v) > 3 {
			return false // 不符合飞机333444的要求
		}
		if len(v) == 3 {
			base = append(base, v)
		}
		if len(v) == 2 {
			extraCouple = append(extraCouple, v)
		}
		if len(v) == 1 {
			extraSingle = append(extraSingle, v)
		}
	}
	// 没有飞机牌，333444
	if len(base) < 2 {
		return false
	}
	// 检查是否符合模式
	if len(extraSingle) != 0 && len(extraCouple) != 0 {
		return false
	}
	// 检查飞机是否成立333444
	list := []*c.Card{}
	for _, v := range base {
		list = append(list, v...)
	}

	SortCards(list)
	if !IsSequence(list) {
		return false
	}

	return true
}

// 炸弹比较
type ValidAirplaneCards struct {
	cards   []*c.Card
	pattern CardsPattern
}

func NewValidAirplaneCards(cards ...*c.Card) *ValidAirplaneCards {
	if !IsValidAirplane(cards...) {
		return nil
	}
	return &ValidAirplaneCards{
		cards:   MostSort(cards),
		pattern: AirplanePattern,
	}
}

func (t *ValidAirplaneCards) GetCards() []*c.Card {
	return t.cards
}

func (t *ValidAirplaneCards) GetPattern() CardsPattern {
	return t.pattern
}

func (t *ValidAirplaneCards) IsSamePattern(ci CardsCompareInterface) bool {
	if len(t.cards) != len(ci.GetCards()) {
		return false
	}
	return t.GetPattern() == ci.GetPattern()
}

func (t *ValidAirplaneCards) IsGreater(ci CardsCompareInterface) bool {
	cur := getSize(t.GetCards()[0].Value)
	tar := getSize(ci.GetCards()[0].Value)
	if cur == -1 || tar == -1 {
		return false
	}
	return cur > tar
}
