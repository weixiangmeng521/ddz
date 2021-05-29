package games

import (
	c "ddz/cards"
	"ddz/compare"
)

// // 前置操作， 比较牌时，检查牌是否合法
// type CardsCompareInterface interface {
// 	GetCards() []*c.Card
// 	GetPattern() compare.CardsPattern
// 	IsSamePattern(CardsCompareInterface) bool
// 	IsGreater(CardsCompareInterface) bool
// }

var (
	cardsMap = map[compare.CardsPattern]func(cards ...*c.Card) bool{}
)

func init() {
	cardsMap[compare.BoomCardsPattern] = compare.IsValidBoom
	cardsMap[compare.SinglePattren] = compare.IsValidSingle
	cardsMap[compare.CoupleCardsPattern] = compare.IsValidCouple
	cardsMap[compare.ThreeCardsPattern] = compare.IsValidThreeCards
	cardsMap[compare.ThreeBeltOneCardsPattern] = compare.IsValidThreeBeltOneCards
	cardsMap[compare.ThreeBeltTwoCardsPattern] = compare.IsValidThreeBeltTwoCards
	cardsMap[compare.StraightCardsPattern] = compare.IsValidStraight
	cardsMap[compare.StraightCoupleCardsPattern] = compare.IsValidStraightCouple
	cardsMap[compare.StraightThreeCardsPattern] = compare.IsValidStraightThree
	cardsMap[compare.FourBeltTwoPattern] = compare.IsValidFourBeltTwo
	cardsMap[compare.AirplanePattern] = compare.IsValidAirplane
}

// 获取牌的类型
func GetCardsPattern(c ...*c.Card) compare.CardsPattern {
	for k, fn := range cardsMap {
		if fn(c...) {
			return k
		}
	}
	return compare.NullPattern
}

// 获取 卡牌对象
func ConvertCards(pattern compare.CardsPattern, cards []*c.Card) compare.CardsCompareInterface {
	var obj compare.CardsCompareInterface
	switch pattern {
	case compare.BoomCardsPattern:
		obj = compare.NewBoomCards(cards...)
	case compare.SinglePattren:
		obj = compare.NewSingleCard(cards...)
	case compare.CoupleCardsPattern:
		obj = compare.NewCoupleCards(cards...)
	case compare.ThreeCardsPattern:
		obj = compare.NewThreeCards(cards...)
	case compare.ThreeBeltOneCardsPattern:
		obj = compare.NewThreeBeltOneCards(cards...)
	case compare.ThreeBeltTwoCardsPattern:
		obj = compare.NewThreeBeltTwoCards(cards...)
	case compare.StraightCardsPattern:
		obj = compare.NewStraightCards(cards...)
	case compare.StraightCoupleCardsPattern:
		obj = compare.NewStraightCoupleCards(cards...)
	case compare.StraightThreeCardsPattern:
		obj = compare.NewStraightThreeCards(cards...)
	case compare.FourBeltTwoPattern:
		obj = compare.NewFourBeltTwoCards(cards...)
	case compare.AirplanePattern:
		obj = compare.NewValidAirplaneCards(cards...)
	}
	return obj
}
