package constant

type CardsPattern int32

// 牌的pattern
const (
	AnyPattern                 CardsPattern = iota //任意类型
	NullPattern                                    // 无效的pattern
	BoomCardsPattern                               // 炸弹
	SinglePattren                                  // 单张
	CoupleCardsPattern                             // 对子
	ThreeCardsPattern                              // 三不带
	ThreeBeltOneCardsPattern                       // 三带一
	ThreeBeltTwoCardsPattern                       // 三带二
	StraightCardsPattern                           // 顺子
	StraightCoupleCardsPattern                     // 姊妹对
	StraightThreeCardsPattern                      // 三顺
	FourBeltTwoPattern                             // 四带二
	AirplanePattern                                // 飞机
)

func (t CardsPattern) ToString() string {
	m := map[CardsPattern]string{
		AnyPattern:                 "AnyPattern",
		NullPattern:                "NullPattern",
		BoomCardsPattern:           "BoomCardsPattern",
		SinglePattren:              "SinglePattren",
		CoupleCardsPattern:         "CoupleCardsPattern",
		ThreeCardsPattern:          "ThreeCardsPattern",
		ThreeBeltOneCardsPattern:   "ThreeBeltOneCardsPattern",
		ThreeBeltTwoCardsPattern:   "ThreeBeltTwoCardsPattern",
		StraightCardsPattern:       "StraightCardsPattern",
		StraightCoupleCardsPattern: "StraightCoupleCardsPattern",
		StraightThreeCardsPattern:  "StraightThreeCardsPattern",
		FourBeltTwoPattern:         "FourBeltTwoPattern",
		AirplanePattern:            "AirplanePattern",
	}
	return m[t]
}
