package players

import (
	c "ddz/app/cards"
	"ddz/app/compare"
	"ddz/app/constant"
	"ddz/app/scene"
	"errors"
	"strings"
)

// 玩家
type Player struct {
	hasCalledLord bool                           // 是否叫地主
	room          string                         // 玩家所属房间名
	name          string                         //玩家名称
	role          constant.RoleType              // 玩家角色
	state         constant.StateType             // 玩家状态
	playedCards   constant.CardsCompareInterface // 玩家出的牌
	cards         []*c.Card                      // 农民17张牌，地主20张，地主需要抢地主
	game          constant.GameInterface         // 玩家所在的游戏
	scene.Hooks
}

// 创建一个player
func NewPlayer(name string) *Player {
	return &Player{
		hasCalledLord: false,
		state:         constant.Waiting,
		name:          name,
		room:          "",
		role:          constant.Farmer,
		cards:         []*c.Card{},
	}
}

// 设置玩家所在的游戏指针
func (t *Player) SetGame(g constant.GameInterface) {
	t.game = g
}

// 获取玩家所在的游戏指针
func (t *Player) GetGame() constant.GameInterface {
	return t.game
}

// 设置玩家状态
func (t *Player) SetState(s constant.StateType) {
	t.state = s
	t.Trigger(constant.PLAYER_STATE_CHANGED, s)
}

// 获取状态
func (t *Player) GetState() constant.StateType {
	return t.state
}

// 设置玩家房间
func (t *Player) SetRoom(s string) {
	t.room = s
}

// 获取玩家房间名
func (t *Player) GetRoom() string {
	return t.room
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
	t.role = constant.Lord
}

// 把玩家变成农民
func (t *Player) SetFarmer() {
	t.role = constant.Farmer
}

// 是不是地主
func (t *Player) IsLord() bool {
	return t.role == constant.Lord
}

// 是不是龙鸣
func (t *Player) IsFarmer() bool {
	return t.role == constant.Lord
}

// 接收牌
func (t *Player) AcceptCards(cards ...*c.Card) {
	for _, card := range cards {
		t.cards = append(t.cards, card)
	}
}

// 设置当前出的牌
func (t *Player) SetPlayedCards(cp constant.CardsCompareInterface) {
	t.playedCards = cp
}

// 获取当前出的牌
func (t *Player) GetPlayedCards() constant.CardsCompareInterface {
	return t.playedCards
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

// 获取牌
func (t *Player) GetCards() []*c.Card {
	return t.cards
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
