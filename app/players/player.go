package players

import (
	c "ddz/app/cards"
	"ddz/app/compare"
	"errors"
	"strings"
)

// 玩家
type Player struct {
	hasCalledLord bool // 是否叫地主
	name          string
	role          RoleType
	cards         []*c.Card // 农民17张牌，地主20张，地主需要抢地主
}

// 创建一个player
func NewPlayer(name string) *Player {
	return &Player{
		hasCalledLord: false,
		name:          name,
		role:          Farmer,
		cards:         []*c.Card{},
	}
}

// 叫地主
func (t *Player) CallLord() {
	t.hasCalledLord = true
}

// 不叫
func (t *Player) NotCall() {
	t.hasCalledLord = false
}

// 是否叫了地主
func (t *Player) HasCalledLord() bool {
	return t.hasCalledLord
}

// 获取玩家名字
func (t *Player) GetName() string {
	return t.name
}

// 把农民变成地主
func (t *Player) SetLord() {
	t.role = Lord
}

// 把玩家变成农民
func (t *Player) SetFarmer() {
	t.role = Farmer
}

// 是不是地主
func (t *Player) IsLord() bool {
	return t.role == Lord
}

// 是不是龙鸣
func (t *Player) IsFarmer() bool {
	return t.role == Lord
}

// 接收牌
func (t *Player) AcceptCards(cards ...*c.Card) {
	for _, card := range cards {
		t.cards = append(t.cards, card)
	}
}

// 出牌
func (t *Player) DealCards(cards ...*c.Card) error {
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
func (t *Player) DealCard(card *c.Card) error {
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

// 清空手牌
func (t *Player) Clear() {
	t.cards = []*c.Card{}
}

// 牌排序
func (t *Player) SortCards() {
	compare.SortCards(t.cards)
}

// 玩家还有没有牌
func (t *Player) HasCards() bool {
	return len(t.cards) != 0
}

// 是不是牌出完了
func (t *Player) HasWinned() bool {
	return len(t.cards) == 0
}
