package games

import (
	"ddz/cards"
	"ddz/compare"
	"ddz/players"
	"errors"
	"fmt"
)

// 游戏状态
type GameState int8

const (
	Ready   GameState = 0
	Started GameState = 1
	End     GameState = 3
)

type Game struct {
	cardsBoot  *cards.CardsBoot // 牌靴
	players    []*players.Player
	lordCards  []*cards.Card
	curPattern compare.CardsPattern          // 当前局的行牌模式
	curCards   compare.CardsCompareInterface // 当前的局面的牌
	curIndex   int8                          // 当前进行的玩家index
	state      GameState
}

func NewGame() *Game {
	return &Game{
		cardsBoot:  cards.NewCardsBoot(),
		players:    []*players.Player{},
		lordCards:  []*cards.Card{},
		curIndex:   0,
		state:      Ready,
		curPattern: compare.NullPattern,
	}
}

// 游戏加入玩家
func (t *Game) JoinPlayer(p *players.Player) bool {
	// 如果游戏人数满了，就加入失败
	if len(t.players) == 3 {
		return false
	}
	t.players = append(t.players, p)
	return true
}

// 游戏离开玩家
func (t *Game) LeavePlayer(p *players.Player) {
	for i, p := range t.players {
		if p.GetName() == p.GetName() {
			t.players = append(t.players[:i], t.players[i+1:]...)
		}
	}
}

// 能否开局
func (t *Game) CanStart() bool {
	return len(t.players) == 3
}

// 洗牌
func (t *Game) Shuffle() {
	t.state = Started
	for _, p := range t.players {
		p.SetFarmer() // 都变成农民
		p.Clear()     //清空手牌
	}
	t.lordCards = []*cards.Card{}
	t.cardsBoot.Init().Shuffle()
}

// 发牌
func (t *Game) Licensing() {
	boot := t.cardsBoot.Interator()
	for i := 0; i < len(t.players); i++ {
		p := t.players[i]
		if boot.HasLastByNum(3) {
			break
		}
		p.AcceptCards(boot.Next())

		if i == 2 {
			i = -1
		}
	}

	// 给玩家的牌排序
	for _, p := range t.players {
		p.SortCards()
	}

	for boot.HasNext() {
		t.lordCards = append(t.lordCards, boot.Next())
	}
}

func (t *Game) Display() {
	tmp := ""
	for _, p := range t.players {
		tmp += p.GetName() + "\t  " + p.CheckCards() + "\n"
	}
	tmp += "Lordcards\t" + cards.NewCardsList(t.lordCards...).ToString() + "\n"
	fmt.Println(tmp)
}

// 叫地主
func (t *Game) CallLandlord(p *players.Player) {
	p.AcceptCards(t.lordCards...)
	p.SetLord()
}

// 出牌
func (t *Game) DealCards(c []*cards.Card) error {
	// 当前的牌模式 和 玩家出的牌模式是否相同
	pattern := GetCardsPattern(c...)
	obj := ConvertCards(pattern, c)

	// 玩家放弃出牌
	if pattern == compare.NullPattern {
		t.Turn()
		return nil
	}

	// 当前局没有炸弹，如果牌型不匹配
	if t.curPattern != compare.BoomCardsPattern &&
		pattern != compare.BoomCardsPattern &&
		t.curPattern != compare.NullPattern &&
		!t.curCards.IsSamePattern(obj) {
		return errors.New("U dealed mismatched cards type: " + pattern.ToString())
	}

	// 当前局没有炸弹，玩家出的牌是否比场上的牌大
	if t.curPattern != compare.BoomCardsPattern &&
		pattern != compare.BoomCardsPattern &&
		t.curPattern != compare.NullPattern &&
		!obj.IsGreater(t.curCards) {
		return errors.New("U should deal greater than last player's cards: " + cards.NewCardsList(t.curCards.GetCards()...).ToString())
	}

	// 如果场上有炸弹，就比较炸弹
	// 如果当前没有炸弹，就直接炸弹
	if t.curPattern == compare.BoomCardsPattern && !obj.IsGreater(t.curCards) {
		return errors.New("U should deal greater than last player's cards: " + cards.NewCardsList(t.curCards.GetCards()...).ToString())
	}

	// 是否出牌成功
	p := t.players[t.curIndex]
	if err := p.DealCards(c...); err != nil {
		return err
	}
	// 如果玩家出完牌，游戏结束
	if !p.HasCards() {
		t.state = End
		return nil
	}
	// 出牌成功, 则设置 curPattern 和 curCards
	t.curPattern = pattern
	t.curCards = obj

	t.Turn()
	return nil
}

// 出牌权交下一位玩家
func (t *Game) Turn() {
	if t.curIndex == 2 {
		t.curIndex = -1
	}
	t.curIndex++
}

// 获取游戏赢家
func (t *Game) GetWiners() []*players.Player {
	if !t.players[t.curIndex].HasCards() {
		// 地主赢
		if t.players[t.curIndex].IsLord() {
			return []*players.Player{t.players[t.curIndex]}
		}
		// 农民赢
		winer := []*players.Player{}
		for _, p := range t.players {
			if !p.IsLord() {
				winer = append(winer, p)
			}
		}
		return winer
	}
	return nil
}

// 比较牌
func (t *Game) CompareCards(c []*cards.Card) bool {
	// 牌比较
	// t.curCards
	// c

	return true
}
