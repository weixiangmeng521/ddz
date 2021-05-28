package players

import (
	"ddz/cards"
	"errors"
	"strings"
)

type Player struct {
	role  RoleType
	cards []*cards.Card // 农民17张牌，地主20张，地主需要抢地主
}

// 创建一个player
func NewPlayer() *Player {
	return &Player{
		role:  Farmer,
		cards: []*cards.Card{},
	}
}

// 把农民变成地主
func (t *Player) SetLord() {
	t.role = Lord
}

// 是不是地主
func (t *Player) IsLord() bool {
	return t.role == Lord
}

// 接收牌
func (t *Player) AcceptCards(cards ...*cards.Card) {
	for _, card := range cards {
		t.cards = append(t.cards, card)
	}
}

// 出牌
func (t *Player) DealCards(cards ...*cards.Card) error {
	// 检测是否有这些牌
	condition := len(cards)
	for _, v := range cards {
		for _, card := range t.cards {
			if v.IsEqual(card) {
				condition--
			}
		}
	}

	arr := []string{}
	for _, v := range t.cards {
		arr = append(arr, v.ToString())
	}
	if condition != 0 {
		return errors.New("Cannot put cards: " + strings.Join(arr, ","))
	}

	for _, v := range cards {
		t.DealCard(v)
	}
	return nil
}

// 出单张牌
func (t *Player) DealCard(card *cards.Card) error {
	i := -1
	for k, v := range t.cards {
		if v.IsEqual(card) {
			i = k
		}
	}
	if i == -1 {
		return errors.New("cannot put card: " + card.ToString())
	}
	t.cards = append(t.cards[:i], t.cards[i+1:]...)
	return nil
}

// 查看手牌
func (t *Player) CheckCards() string {
	arr := []string{}
	for _, v := range t.cards {
		arr = append(arr, v.ToString())
	}
	return strings.Join(arr, ",")
}
