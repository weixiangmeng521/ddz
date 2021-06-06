package compare

import (
	c "ddz/app/cards"
	"fmt"
)

// 关系大小的List
var relateList = []string{"3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K", "A", "2", "Jack"}

// 获取牌在list中的位置
func getSize(s string) int {
	for i, v := range relateList {
		if v == s {
			return i
		}
	}
	return -1
}

// 牌排序
func SortCards(cards []*c.Card) []*c.Card {
	for i := 0; i <= len(cards); i++ {
		for j := i + 1; j < len(cards); j++ {
			tmp := cards[i]
			if getSize(cards[i].Value) > getSize(cards[j].Value) {
				cards[i] = cards[j]
				cards[j] = tmp
			}
		}
	}
	return cards
}

// 相同的牌放一起的排序
func MostSort(cards []*c.Card) []*c.Card {
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
		SortCards(list)
		arr = append(arr, list)
	}

	// group 排序
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
	return res
}

// 是不是这些牌面值相等
func IsCardsEqual(cards ...*c.Card) bool {
	for i := 1; i < len(cards); i++ {
		if !cards[i-1].IsValueEqual(cards[i]) {
			return false
		}
	}
	return true
}

// 是不是包含这些卡牌
func HasContain(list []*c.Card, requires ...*c.Card) bool {
	for _, v := range list {
		for _, card := range requires {
			if v.IsEqual(card) {
				return true
			}
		}
	}
	return false
}

// 是不是序列：格式如下
// 112233
// 123
func IsSequence(cards []*c.Card) bool {
	tmp := 0 // 获取112233中，11的长度
	for cards[tmp].Value == cards[tmp+1].Value {
		tmp++
	}
	prev := -1

	// cards.NewCardsList(cards...).Display()

	for i := 0; i < len(cards)-1; i++ {
		j := i
		size := 0
		for j < len(cards)-1 && cards[j].Value == cards[j+1].Value {
			j++
			size++
		}
		i += size
		if size != tmp {
			fmt.Println(size)
			fmt.Println(tmp)
			return false // 不符合规律
		}
		// 设置初始值
		if prev == -1 {
			prev = getSize(cards[i].Value)
			continue
		}
		cur := getSize(cards[i].Value)
		if cur-1 != prev {
			return false // 不是升序
		}
		prev = cur
	}

	return true
}
