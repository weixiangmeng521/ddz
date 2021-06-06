package games

import (
	c "ddz/app/cards"
	"ddz/app/compare"
	"ddz/app/constant"
)

var (
	cardsMap = map[constant.CardsPattern]func(cards ...*c.Card) bool{
		constant.BoomCardsPattern:           compare.IsValidBoom,
		constant.SinglePattren:              compare.IsValidSingle,
		constant.CoupleCardsPattern:         compare.IsValidCouple,
		constant.ThreeCardsPattern:          compare.IsValidThreeCards,
		constant.ThreeBeltOneCardsPattern:   compare.IsValidThreeBeltOneCards,
		constant.ThreeBeltTwoCardsPattern:   compare.IsValidThreeBeltTwoCards,
		constant.StraightCardsPattern:       compare.IsValidStraight,
		constant.StraightCoupleCardsPattern: compare.IsValidStraightCouple,
		constant.StraightThreeCardsPattern:  compare.IsValidStraightThree,
		constant.FourBeltTwoPattern:         compare.IsValidFourBeltTwo,
		constant.AirplanePattern:            compare.IsValidAirplane,
	}
)

// 获取牌的类型
func GetCardsPattern(c ...*c.Card) constant.CardsPattern {
	for k, fn := range cardsMap {
		if fn(c...) {
			return k
		}
	}
	return constant.NullPattern
}

// 获取 卡牌对象
func ConvertCards(pattern constant.CardsPattern, cards []*c.Card) constant.CardsCompareInterface {
	var obj constant.CardsCompareInterface
	switch pattern {
	case constant.BoomCardsPattern:
		obj = compare.NewBoomCards(cards...)
	case constant.SinglePattren:
		obj = compare.NewSingleCard(cards...)
	case constant.CoupleCardsPattern:
		obj = compare.NewCoupleCards(cards...)
	case constant.ThreeCardsPattern:
		obj = compare.NewThreeCards(cards...)
	case constant.ThreeBeltOneCardsPattern:
		obj = compare.NewThreeBeltOneCards(cards...)
	case constant.ThreeBeltTwoCardsPattern:
		obj = compare.NewThreeBeltTwoCards(cards...)
	case constant.StraightCardsPattern:
		obj = compare.NewStraightCards(cards...)
	case constant.StraightCoupleCardsPattern:
		obj = compare.NewStraightCoupleCards(cards...)
	case constant.StraightThreeCardsPattern:
		obj = compare.NewStraightThreeCards(cards...)
	case constant.FourBeltTwoPattern:
		obj = compare.NewFourBeltTwoCards(cards...)
	case constant.AirplanePattern:
		obj = compare.NewValidAirplaneCards(cards...)
	}
	return obj
}
