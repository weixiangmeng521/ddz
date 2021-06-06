package cards

import "fmt"

// 2 - 10
// J Q K A
// 每张牌有 Spade, Heart, Diamond, Club
// Joker1, Jack2

type CardType int32

// Cards type
const (
	// For Normal
	Spade   CardType = 0
	Heart   CardType = 1
	Diamond CardType = 2
	Club    CardType = 3
	// special for Jack
	Freak CardType = 4
	Real  CardType = 5
)

// 一般卡牌值
var cardValue = [13]string{
	"3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K", "A", "2",
}

// 特殊卡牌
var specialCardValue = [1]string{
	"Jack",
}

// 是不是合法的卡片值
func isValidCardValue(value string) bool {
	for _, v := range cardValue {
		if v == value {
			return true
		}
	}
	for _, v := range specialCardValue {
		if v == value {
			return true
		}
	}
	return false
}

// 是不是特殊卡牌 Jack
func isSpecialCard(value string) bool {
	for _, v := range specialCardValue {
		if v == value {
			return true
		}
	}
	return false
}

// 卡牌
type Card struct {
	Value string   `json:"value"`
	Type  CardType `json:"type"`
}

func NewCard(v string, t CardType) *Card {
	// 特殊卡牌
	if isSpecialCard(v) {
		if t == Freak || t == Real {
			return &Card{v, t}
		}
		return nil
	}
	// 一般卡牌
	if !isValidCardValue(v) || t == Freak || t == Real {
		fmt.Println("invalid card value: ", v)
		return nil
	}
	return &Card{v, t}
}

func (t *Card) ToString() string {
	mapper := map[CardType]string{
		Spade:   "♠", // 黑桃
		Heart:   "♥", // 红桃
		Diamond: "♦", // 方块
		Club:    "♣", // 梅花
		Freak:   "小",
		Real:    "大",
	}
	return mapper[t.Type] + t.Value
}

// 是否等于
func (t *Card) IsEqual(card *Card) bool {
	return t.Value == card.Value && t.Type == card.Type
}

// 是不是值相同
func (t *Card) IsValueEqual(card *Card) bool {
	return t.Value == card.Value
}
