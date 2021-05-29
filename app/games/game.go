package games

import (
	"ddz/app/cards"
	"ddz/app/compare"
	"ddz/app/constant"
	"errors"
	"fmt"
	"math/rand"
	"time"
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
	players    []constant.PlayerInterface
	lordCards  []*cards.Card
	curPattern compare.CardsPattern          // 当前局的行牌模式
	curCards   compare.CardsCompareInterface // 当前的局面的牌
	curIndex   int8                          // 当前进行的玩家index
	state      GameState
}

func NewGame() *Game {
	return &Game{
		cardsBoot:  cards.NewCardsBoot(),
		players:    []constant.PlayerInterface{},
		lordCards:  []*cards.Card{},
		curIndex:   0,
		state:      Ready,
		curPattern: compare.NullPattern,
	}
}

// 游戏加入玩家
func (t *Game) JoinPlayer(p constant.PlayerInterface) bool {
	// 如果游戏人数满了，就加入失败
	if len(t.players) == 3 {
		return false
	}
	t.players = append(t.players, p)
	return true
}

// 游戏离开玩家
func (t *Game) LeavePlayer(p constant.PlayerInterface) {
	for i, v := range t.players {
		if v.GetName() == p.GetName() {
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
		p.NotCall()   // 初始都不叫地主
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
	tmp := "\n"
	for _, p := range t.players {
		tmp += p.GetName() + "\t" + p.CheckCards() + "\t\n"
	}
	tmp += "Lordcards\t" + cards.NewCardsList(t.lordCards...).ToString() + "\t\n"
	fmt.Println(tmp)
}

// 叫地主
func (t *Game) CallLandlord() bool {
	// 获取叫了地主的玩家
	called := []constant.PlayerInterface{}
	for _, p := range t.players {
		if p.HasCalledLord() {
			called = append(called, p)
		}
	}
	// 看谁叫了地主， 如果都没叫返回false
	if len(called) == 0 {
		return false
	}

	// 随机获取一个叫了地主的玩家，让他成为地主
	rand.Seed(time.Now().Unix()) // initialize global pseudo random generator
	lord := called[rand.Intn(len(called))]

	// 设置地主先出跑
	for i, p := range t.players {
		if p.IsLord() {
			t.curIndex = int8(i)
		}
	}

	lord.AcceptCards(t.lordCards...)
	lord.SetLord()
	lord.SortCards()
	return true
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

// 获取当前有出牌权的玩家
func (t *Game) GetCurPlayer() constant.PlayerInterface {
	return t.players[t.curIndex]
}

// 获取游戏赢家
func (t *Game) GetWiners() []constant.PlayerInterface {
	if !t.players[t.curIndex].HasCards() {
		// 地主赢
		if t.players[t.curIndex].IsLord() {
			return []constant.PlayerInterface{t.players[t.curIndex]}
		}
		// 农民赢
		winer := []constant.PlayerInterface{}
		for _, p := range t.players {
			if !p.IsLord() {
				winer = append(winer, p)
			}
		}
		return winer
	}
	return nil
}

// 是不是游戏结束了
func (t *Game) HasGoodGame() bool {
	for _, p := range t.players {
		if p.HasWinned() {
			return true
		}
	}
	return false
}

// 比较牌
func (t *Game) CompareCards(c []*cards.Card) bool {
	// 牌比较
	// t.curCards
	// c

	return true
}
