package compare

import (
	"ddz/app/cards"
	"fmt"
	"strings"
	"testing"
)

// 测试这些牌是否相等
func TestIsCardsEqual(t *testing.T) {
	res := IsCardsEqual(cards.NewCard("3", cards.Spade), cards.NewCard("3", cards.Spade), cards.NewCard("3", cards.Spade))
	if res != true {
		t.Error()
	}
	res = IsCardsEqual(cards.NewCard("4", cards.Spade), cards.NewCard("3", cards.Spade), cards.NewCard("3", cards.Spade))
	if res != false {
		t.Error()
	}
	res = IsCardsEqual(cards.NewCard("3", cards.Spade), cards.NewCard("3", cards.Spade), cards.NewCard("5", cards.Spade))
	if res != false {
		t.Error()
	}
}

// 牌排序
func TestSortCards(t *testing.T) {
	testCase := []*cards.Card{
		cards.NewCard("3", cards.Spade),
		cards.NewCard("4", cards.Spade),
		cards.NewCard("5", cards.Spade),
		cards.NewCard("4", cards.Spade),
		cards.NewCard("K", cards.Spade),
		cards.NewCard("Jack", cards.Freak),
		cards.NewCard("4", cards.Spade),
		cards.NewCard("A", cards.Spade),
		cards.NewCard("2", cards.Spade),
		cards.NewCard("J", cards.Spade),
	}

	SortCards(testCase)
	arr := []string{}
	for _, v := range testCase {
		arr = append(arr, v.ToString())
	}
	fmt.Println(strings.Join(arr, ","))
}

// 单张`
func TestSingleCompare(t *testing.T) {
	obj1 := NewSingleCard(cards.NewCard("3", cards.Spade))
	obj2 := NewSingleCard(cards.NewCard("4", cards.Spade))
	obj3 := NewSingleCard(cards.NewCard("J", cards.Spade))
	obj4 := NewSingleCard(cards.NewCard("Jack", cards.Freak))

	res := obj2.IsGreater(obj1) //4 ,3
	if res != true {
		t.Error()
	}

	res = obj2.IsGreater(obj3) // 4, j
	if res != false {
		t.Error()
	}

	res = obj2.IsGreater(obj4) // 4, j
	if res != false {
		t.Error()
	}
}

func TestValidSingle(t *testing.T) {
	var testCase = []*cards.Card{
		cards.NewCard("3", cards.Spade),
		cards.NewCard("4", cards.Spade),
		cards.NewCard("J", cards.Spade),
		cards.NewCard("Jack", cards.Spade),
	}
	if IsValidSingle(testCase...) != false {
		t.Error()
	}
	if IsValidSingle(testCase[0]) != true {
		t.Error()
	}
}

// 对子
func TestCoupleCompare(t *testing.T) {
	couple1 := []*cards.Card{
		cards.NewCard("3", cards.Spade),
		cards.NewCard("3", cards.Spade),
	}

	couple2 := []*cards.Card{
		cards.NewCard("4", cards.Spade),
		cards.NewCard("4", cards.Spade),
	}

	couple3 := []*cards.Card{
		cards.NewCard("J", cards.Spade),
		cards.NewCard("J", cards.Spade),
	}

	// cur := cards.NewCard("4", cards.Spade)
	obj1 := NewCoupleCards(couple1...)
	obj2 := NewCoupleCards(couple2...)
	obj3 := NewCoupleCards(couple3...)

	if obj2.IsGreater(obj1) != true {
		t.Error()
	} // 44, 33
	if obj2.IsGreater(obj2) != false {
		t.Error()
	} // 44, 44
	if obj2.IsGreater(obj3) != false {
		t.Error()
	} // 44, jj
}

func TestValidCouple(t *testing.T) {
	var testCase = []*cards.Card{
		cards.NewCard("3", cards.Spade),
		cards.NewCard("3", cards.Spade),
	}
	if IsValidCouple(testCase...) != true {
		t.Error()
	}

	testCase = []*cards.Card{
		cards.NewCard("3", cards.Spade),
		cards.NewCard("4", cards.Spade),
	}
	if IsValidCouple(testCase...) != false {
		t.Error()
	}

	if IsValidCouple(testCase[0]) != false {
		t.Error()
	}

}

func TestValidBoom(t *testing.T) {
	var testCase = []*cards.Card{
		cards.NewCard("3", cards.Spade),
		cards.NewCard("3", cards.Spade),
		cards.NewCard("3", cards.Spade),
		cards.NewCard("3", cards.Spade),
	}
	if IsValidBoom(testCase...) != true {
		t.Error()
	}

	testCase = []*cards.Card{
		cards.NewCard("3", cards.Spade),
		cards.NewCard("4", cards.Spade),
		cards.NewCard("3", cards.Spade),
		cards.NewCard("4", cards.Spade),
	}
	if IsValidBoom(testCase...) != false {
		t.Error()
	}

	testCase = []*cards.Card{
		cards.NewCard("Jack", cards.Freak),
		cards.NewCard("Jack", cards.Real),
	}
	if IsValidBoom(testCase...) != true {
		t.Error()
	}

	testCase = []*cards.Card{
		cards.NewCard("Jack", cards.Freak),
		cards.NewCard("4", cards.Spade),
	}
	if IsValidBoom(testCase...) != false {
		t.Error()
	}
}

func TestBoomCompare(t *testing.T) {
	obj1 := NewBoomCards(cards.NewCard("3", cards.Spade), cards.NewCard("3", cards.Spade), cards.NewCard("3", cards.Spade), cards.NewCard("3", cards.Spade))
	obj2 := NewBoomCards(cards.NewCard("2", cards.Spade), cards.NewCard("2", cards.Spade), cards.NewCard("2", cards.Spade), cards.NewCard("2", cards.Spade))
	obj3 := NewBoomCards(cards.NewCard("Jack", cards.Freak), cards.NewCard("Jack", cards.Real))

	if obj2.IsGreater(obj1) != true {
		t.Error()
	}
	if obj2.IsGreater(obj2) != false {
		t.Error()
	}
	if obj2.IsGreater(obj3) != false {
		t.Error()
	}
}

func TestIsValidThreeBeltOneCards(t *testing.T) {
	res := IsValidThreeBeltOneCards(
		cards.NewCard("3", cards.Spade),
		cards.NewCard("3", cards.Spade),
		cards.NewCard("4", cards.Spade),
		cards.NewCard("3", cards.Spade),
	)
	if res != true {
		t.Error()
	}

	res = IsValidThreeBeltOneCards(
		cards.NewCard("5", cards.Spade),
		cards.NewCard("3", cards.Spade),
		cards.NewCard("5", cards.Spade),
		cards.NewCard("5", cards.Spade),
	)
	if res != true {
		t.Error()
	}

	res = IsValidThreeBeltOneCards(
		cards.NewCard("5", cards.Spade),
		cards.NewCard("A", cards.Spade),
		cards.NewCard("A", cards.Spade),
		cards.NewCard("5", cards.Spade),
	)
	if res != false {
		t.Error()
	}
}

func BenchmarkMostSort(t *testing.B) {
	for i := 0; i < 5; i++ {
		testCase := []*cards.Card{
			cards.NewCard("2", cards.Spade),
			cards.NewCard("3", cards.Spade),
			cards.NewCard("2", cards.Spade),
			cards.NewCard("2", cards.Spade),
		}
		res := MostSort(testCase)
		arr := []string{}
		for _, v := range res {
			arr = append(arr, v.ToString())
		}
		if strings.Join(arr, ",") != "♠2,♠2,♠2,♠3" {
			t.Error(strings.Join(arr, ","))
		}

		arr = []string{}

		testCase = []*cards.Card{
			cards.NewCard("3", cards.Spade),
			cards.NewCard("Q", cards.Spade),
			cards.NewCard("3", cards.Spade),
			cards.NewCard("3", cards.Spade),
		}
		res = MostSort(testCase)
		for _, v := range res {
			arr = append(arr, v.ToString())
		}
		if strings.Join(arr, ",") != "♠3,♠3,♠3,♠Q" {
			t.Error(strings.Join(arr, ","))
		}

	}
}

func TestMostSort(t *testing.T) {
	testCase := []*cards.Card{
		cards.NewCard("2", cards.Spade),
		cards.NewCard("3", cards.Spade),
		cards.NewCard("2", cards.Spade),
		cards.NewCard("2", cards.Spade),
	}
	res := MostSort(testCase)
	arr := []string{}
	for _, v := range res {
		arr = append(arr, v.ToString())
	}
	fmt.Println(strings.Join(arr, ","))
	arr = []string{}

	testCase = []*cards.Card{
		cards.NewCard("3", cards.Spade),
		cards.NewCard("Q", cards.Spade),
		cards.NewCard("3", cards.Spade),
		cards.NewCard("3", cards.Spade),
	}
	res = MostSort(testCase)
	for _, v := range res {
		arr = append(arr, v.ToString())
	}
	fmt.Println(strings.Join(arr, ","))
	arr = []string{}

	testCase = []*cards.Card{
		cards.NewCard("Q", cards.Spade),
		cards.NewCard("3", cards.Spade),
		cards.NewCard("Q", cards.Spade),
		cards.NewCard("Q", cards.Spade),
	}
	res = MostSort(testCase)
	for _, v := range res {
		arr = append(arr, v.ToString())
	}
	fmt.Println(strings.Join(arr, ","))
	arr = []string{}

	testCase = []*cards.Card{
		cards.NewCard("Q", cards.Spade),
		cards.NewCard("Q", cards.Spade),
		cards.NewCard("Q", cards.Spade),
		cards.NewCard("Q", cards.Spade),
	}
	res = MostSort(testCase)
	for _, v := range res {
		arr = append(arr, v.ToString())
	}
	fmt.Println(strings.Join(arr, ","))
	arr = []string{}

	testCase = []*cards.Card{
		cards.NewCard("Q", cards.Spade),
		cards.NewCard("Q", cards.Spade),
		cards.NewCard("3", cards.Spade),
		cards.NewCard("Q", cards.Spade),
		cards.NewCard("3", cards.Spade),
	}
	res = MostSort(testCase)
	for _, v := range res {
		arr = append(arr, v.ToString())
	}
	fmt.Println(strings.Join(arr, ","))
	arr = []string{}

	testCase = []*cards.Card{
		cards.NewCard("A", cards.Spade),
		cards.NewCard("A", cards.Spade),
		cards.NewCard("3", cards.Spade),
		cards.NewCard("A", cards.Spade),
		cards.NewCard("3", cards.Spade),
		cards.NewCard("J", cards.Spade),
		cards.NewCard("J", cards.Spade),
	}
	res = MostSort(testCase)
	for _, v := range res {
		arr = append(arr, v.ToString())
	}
	fmt.Println(strings.Join(arr, ","))
	arr = []string{}
}

func TestThreeBeltOneCardsIsGreater(t *testing.T) {
	testCase := []*cards.Card{
		cards.NewCard("Q", cards.Spade),
		cards.NewCard("3", cards.Spade),
		cards.NewCard("Q", cards.Spade),
		cards.NewCard("Q", cards.Spade),
	}
	obj1 := NewThreeBeltOneCards(testCase...) // QQQ3

	testCase = []*cards.Card{
		cards.NewCard("3", cards.Spade),
		cards.NewCard("Q", cards.Spade),
		cards.NewCard("3", cards.Spade),
		cards.NewCard("3", cards.Spade),
	}
	obj2 := NewThreeBeltOneCards(testCase...) // 333Q

	if obj1.IsGreater(obj2) != true {
		t.Error()
	}
	if obj1.IsGreater(obj1) != false {
		t.Error()
	}
}

func TestIsValidThreeBeltTwoCards(t *testing.T) {
	testCase := []*cards.Card{
		cards.NewCard("Q", cards.Spade),
		cards.NewCard("3", cards.Spade),
		cards.NewCard("3", cards.Spade),
		cards.NewCard("Q", cards.Spade),
		cards.NewCard("Q", cards.Spade),
	}
	res := IsValidThreeBeltTwoCards(testCase...) // QQQ3
	if res != true {
		t.Error()
	}

	testCase = []*cards.Card{
		cards.NewCard("Q", cards.Spade),
		cards.NewCard("3", cards.Spade),
		cards.NewCard("3", cards.Spade),
		cards.NewCard("Q", cards.Spade),
	}
	res = IsValidThreeBeltTwoCards(testCase...) // QQQ3
	if res != false {
		t.Error()
	}
}

func TestIsValidStraight(t *testing.T) {
	testCase := []*cards.Card{
		cards.NewCard("3", cards.Spade),
		cards.NewCard("5", cards.Spade),
		cards.NewCard("6", cards.Spade),
		cards.NewCard("7", cards.Spade),
		cards.NewCard("9", cards.Spade),
	}
	if IsValidStraight(testCase...) != false {
		t.Error()
	}

	testCase = []*cards.Card{
		cards.NewCard("7", cards.Spade),
		cards.NewCard("3", cards.Spade),
		cards.NewCard("4", cards.Spade),
		cards.NewCard("5", cards.Spade),
		cards.NewCard("6", cards.Spade),
	}
	if IsValidStraight(testCase...) != true {
		t.Error()
	}

	testCase = []*cards.Card{
		cards.NewCard("Jack", cards.Real),
		cards.NewCard("3", cards.Spade),
		cards.NewCard("4", cards.Spade),
		cards.NewCard("5", cards.Spade),
		cards.NewCard("6", cards.Spade),
	}
	if IsValidStraight(testCase...) != false {
		t.Error()
	}

	testCase = []*cards.Card{
		cards.NewCard("6", cards.Spade),
		cards.NewCard("3", cards.Spade),
		cards.NewCard("4", cards.Spade),
		cards.NewCard("5", cards.Spade),
		cards.NewCard("6", cards.Spade),
	}
	if IsValidStraight(testCase...) != false {
		t.Error()
	}
}

func TestStraightCardsIsGreater(t *testing.T) {
	testCase := []*cards.Card{
		cards.NewCard("6", cards.Spade),
		cards.NewCard("7", cards.Spade),
		cards.NewCard("8", cards.Spade),
		cards.NewCard("9", cards.Spade),
		cards.NewCard("10", cards.Spade),
	}
	obj1 := NewStraightCards(testCase...)
	testCase1 := []*cards.Card{
		cards.NewCard("9", cards.Spade),
		cards.NewCard("10", cards.Spade),
		cards.NewCard("J", cards.Spade),
		cards.NewCard("Q", cards.Spade),
		cards.NewCard("K", cards.Spade),
	}
	obj2 := NewStraightCards(testCase1...)

	if obj1.IsGreater(obj2) != false {
		t.Error()
	}
	if obj2.IsGreater(obj1) != true {
		t.Error()
	}
	if obj2.IsGreater(obj2) != false {
		t.Error()
	}
}

func TestIsValidStraightCouple(t *testing.T) {
	testCase := []*cards.Card{
		cards.NewCard("6", cards.Spade),
		cards.NewCard("6", cards.Spade),
		cards.NewCard("7", cards.Spade),
		cards.NewCard("7", cards.Spade),
		cards.NewCard("8", cards.Spade),
		cards.NewCard("8", cards.Spade),
	}
	cards.NewCardsList(testCase...).Display()
	obj1 := IsValidStraightCouple(testCase...)
	if obj1 != true {
		t.Error()
	}

	testCase = []*cards.Card{
		cards.NewCard("6", cards.Spade),
		cards.NewCard("8", cards.Spade),
		cards.NewCard("7", cards.Spade),
		cards.NewCard("6", cards.Spade),
		cards.NewCard("8", cards.Spade),
		cards.NewCard("7", cards.Spade),
	}
	cards.NewCardsList(testCase...).Display()
	obj1 = IsValidStraightCouple(testCase...)
	if obj1 != true {
		t.Error()
	}

	testCase = []*cards.Card{
		cards.NewCard("J", cards.Spade),
		cards.NewCard("Q", cards.Spade),
		cards.NewCard("K", cards.Spade),
		cards.NewCard("J", cards.Spade),
		cards.NewCard("10", cards.Spade),
		cards.NewCard("9", cards.Spade),
	}
	cards.NewCardsList(testCase...).Display()
	obj1 = IsValidStraightCouple(testCase...)
	if obj1 != false {
		t.Error()
	}

	testCase = []*cards.Card{
		cards.NewCard("J", cards.Spade),
		cards.NewCard("J", cards.Spade),
		cards.NewCard("9", cards.Spade),
		cards.NewCard("9", cards.Spade),
		cards.NewCard("A", cards.Spade),
		cards.NewCard("A", cards.Spade),
	}
	cards.NewCardsList(testCase...).Display()
	obj1 = IsValidStraightCouple(testCase...)
	if obj1 != false {
		t.Error()
	}
}

func TestIsValidStraightThree(t *testing.T) {
	testCase := []*cards.Card{
		cards.NewCard("J", cards.Spade),
		cards.NewCard("J", cards.Spade),
		cards.NewCard("Q", cards.Spade),
		cards.NewCard("Q", cards.Spade),
		cards.NewCard("K", cards.Spade),
		cards.NewCard("K", cards.Spade),
		cards.NewCard("A", cards.Spade),
		cards.NewCard("A", cards.Spade),
	}
	cards.NewCardsList(testCase...).Display()
	obj1 := IsValidStraightThree(testCase...)
	if obj1 != false {
		t.Error()
	}

	testCase = []*cards.Card{
		cards.NewCard("J", cards.Spade),
		cards.NewCard("J", cards.Spade),
		cards.NewCard("J", cards.Spade),
		cards.NewCard("Q", cards.Spade),
		cards.NewCard("Q", cards.Spade),
		cards.NewCard("Q", cards.Spade),
	}
	cards.NewCardsList(testCase...).Display()
	obj1 = IsValidStraightThree(testCase...)
	if obj1 != true {
		t.Error()
	}
}

func TestIsValidFourBeltTwo(t *testing.T) {
	testCase := []*cards.Card{
		cards.NewCard("J", cards.Spade),
		cards.NewCard("J", cards.Spade),
		cards.NewCard("K", cards.Spade),
		cards.NewCard("Q", cards.Spade),
		cards.NewCard("J", cards.Spade),
		cards.NewCard("J", cards.Spade),
	}
	cards.NewCardsList(testCase...).Display()
	obj1 := IsValidFourBeltTwo(testCase...)
	if obj1 != true {
		t.Error()
	}

	testCase = []*cards.Card{
		cards.NewCard("J", cards.Spade),
		cards.NewCard("J", cards.Spade),
		cards.NewCard("J", cards.Spade),
		cards.NewCard("Q", cards.Spade),
		cards.NewCard("Q", cards.Spade),
		cards.NewCard("Q", cards.Spade),
	}
	cards.NewCardsList(testCase...).Display()
	obj1 = IsValidFourBeltTwo(testCase...)
	if obj1 != false {
		t.Error()
	}
}

func TestIsValidAirplane(t *testing.T) {
	testCase := []*cards.Card{
		cards.NewCard("J", cards.Spade),
		cards.NewCard("J", cards.Spade),
		cards.NewCard("9", cards.Spade),
		cards.NewCard("9", cards.Spade),
		cards.NewCard("Q", cards.Spade),
		cards.NewCard("Q", cards.Spade),
		cards.NewCard("10", cards.Spade),
		cards.NewCard("10", cards.Spade),
		cards.NewCard("9", cards.Spade),
		cards.NewCard("10", cards.Spade),
	}
	// cards.NewCardsList(testCase...).Display()
	obj1 := IsValidAirplane(testCase...)
	if obj1 != true {
		t.Error()
	}
}

func TestIsSequence(t *testing.T) {
	testCase := []*cards.Card{
		cards.NewCard("9", cards.Spade),
		cards.NewCard("9", cards.Spade),
		cards.NewCard("10", cards.Spade),
		cards.NewCard("10", cards.Spade),
		cards.NewCard("J", cards.Spade),
		cards.NewCard("J", cards.Spade),
		cards.NewCard("Q", cards.Spade),
		cards.NewCard("Q", cards.Spade),
	}
	obj1 := IsSequence(testCase)
	if obj1 != true {
		t.Error()
	}

	testCase = []*cards.Card{
		cards.NewCard("9", cards.Spade),
		cards.NewCard("9", cards.Spade),
		cards.NewCard("9", cards.Spade),
		cards.NewCard("10", cards.Spade),
		cards.NewCard("10", cards.Spade),
		cards.NewCard("10", cards.Spade),
	}
	obj1 = IsSequence(testCase)
	if obj1 != true {
		t.Error()
	}
}
