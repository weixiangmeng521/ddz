package games

import (
	"ddz/app/cards"
	"ddz/app/compare"
	"ddz/app/constant"
	"ddz/app/scene"
	"errors"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Game struct {
	name          string
	cardsBoot     *cards.CardsBoot                                 // 牌靴
	players       []constant.PlayerInterface                       // 游戏玩家
	lordCards     []*cards.Card                                    // 地主的牌
	curCards      constant.CardsCompareInterface                   // 当前的局面的牌
	curIndex      int8                                             // 当前进行的玩家index
	state         constant.GameState                               // 当前游戏状态
	flow          *scene.SceneFlow                                 // 游戏进行流程
	hooks         map[constant.GameHookType][]func(...interface{}) // 钩子
	isCalledLoard bool                                             // 是否叫完地主
	isSortCards   bool                                             // 是否洗牌
	debug         bool                                             // 设置debug模式

	sync.Mutex  // 锁
	scene.Hooks // 钩子
}

func NewGame(name string) (g *Game) {
	g = &Game{
		name:          name,
		cardsBoot:     cards.NewCardsBoot(),
		players:       []constant.PlayerInterface{},
		lordCards:     []*cards.Card{},
		curIndex:      0,
		curCards:      compare.NewAnyCards(),
		debug:         false,
		isCalledLoard: false,
		state:         constant.GameReady,
		hooks:         map[constant.GameHookType][]func(...interface{}){},
		isSortCards:   false,
	}
	// 创建游戏，直接进入flow
	g.flow = scene.CreateSceneFlow(g)
	return
}

// 设置游戏名称
func (t *Game) SetName(name string) {
	t.name = name
}

// 获取房间名称
func (t *Game) GetName() string {
	return t.name
}

// 设置debug模式
func (t *Game) Debug() {
	t.cardsBoot = GetDebugCardsBoot()
	t.debug = true
}

// 设置状态
func (t *Game) SetState(state constant.GameState) {
	t.state = state
	t.Trigger(constant.GAME_STATE_CHANGED, state)
}

// 获取当前的状态
func (t *Game) GetState() constant.GameState {
	return t.state
}

// 获取当前的牌信息
func (t *Game) GetCards() constant.CardsCompareInterface {
	return t.curCards
}

// 获取当前场上的牌
func (t *Game) GetPlayedCards() []*cards.Card {
	if t.curCards == nil {
		return nil
	}
	return t.curCards.GetCards()
}

// 获取地主牌
func (t *Game) GetLordCards() []*cards.Card {
	return t.lordCards
}

// 游戏加入玩家
func (t *Game) JoinPlayer(p constant.PlayerInterface) bool {
	t.Lock()
	defer t.Unlock()
	// 如果游戏人数满了，就加入失败
	if len(t.players) == 3 {
		return false
	}
	// 相同id 不能同时加入
	for _, v := range t.players {
		if v.GetName() == p.GetName() {
			return false
		}
	}
	// 给玩家写入房间信息
	p.SetGame(t)

	t.players = append(t.players, p)
	// 给加入的玩家添加钩子
	t.hookPlayer(p)
	// 触发钩子
	t.Trigger(constant.GAME_JOINED_PLYAER)
	return true
}

// 游戏离开玩家
func (t *Game) LeavePlayer(p constant.PlayerInterface) {
	t.Lock()
	defer t.Unlock()
	for i, v := range t.players {
		if v.GetName() == p.GetName() {
			t.players = append(t.players[:i], t.players[i+1:]...)
			// 清除玩家所在的游戏指针
			p.SetGame(nil)
			// 给玩家清除钩子
			t.clearPlayerHooks(p)
			// 触发钩子
			t.Trigger(constant.GAME_LEAVED_PLYAER)
			return
		}
	}
}

// 能否开局
func (t *Game) CanStart() bool {
	// 3个玩家才能开始游戏
	if len(t.players) != 3 {
		return false
	}
	// 玩家要准备好了才能开始游戏
	for _, p := range t.players {
		if p.GetState() != constant.Already {
			return false
		}
	}
	return true
}

// 遍历玩家
func (t *Game) MapPlayers(cb func(k int, p constant.PlayerInterface)) {
	for k, v := range t.players {
		cb(k, v)
	}
}

// 洗牌
func (t *Game) Shuffle() {
	for _, p := range t.players {
		p.NotCall()   // 初始都不叫地主
		p.SetFarmer() // 都变成农民
		p.Clear()     //清空手牌
	}
	t.lordCards = []*cards.Card{}

	if t.debug {
		return
	}
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
	if t.isSortCards {
		for _, p := range t.players {
			p.SortCards()
		}
	}

	for boot.HasNext() {
		t.lordCards = append(t.lordCards, boot.Next())
	}

	t.Trigger(constant.GAME_CARDS_CHANGED) // 触发洗完牌的钩子
}

func (t *Game) Display() {
	tmp := "\n"
	for _, p := range t.players {
		tmp += p.GetName() + "\t" + p.CheckCards() + "\t\n"
	}
	tmp += "Lordcards\t" + cards.NewCardsList(t.lordCards...).ToString() + "\t\n"
	fmt.Println(tmp)
}

// 把出牌权给地主
func (t *Game) ChangeTurn2Lord() {
	for i, p := range t.players {
		if p.IsLord() {
			t.curIndex = int8(i)
		}
	}
}

// 叫地主
func (t *Game) CallLandlord() bool {
	if t.isCalledLoard == true {
		return false
	}

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

	lord.AcceptCards(t.lordCards...)
	lord.SetLord()
	if t.isSortCards {
		lord.SortCards()
	}

	t.isCalledLoard = true
	// 设置游戏状态
	t.SetState(constant.GameCalled)
	return true
}

// 出牌,
func (t *Game) DealCards(c []*cards.Card) error {
	// 如果一轮下来，其他玩家都不出牌, 当前牌型变any
	played := []constant.PlayerInterface{}
	isPre := false
	t.MapPlayers(func(k int, p constant.PlayerInterface) {
		if p.GetPlayedCards() != nil {
			played = append(played, p)
		}
	})
	if len(played) == 1 {
		if played[0].GetName() == t.GetCurPlayer().GetName() {
			t.curCards = compare.NewAnyCards()
			isPre = true
		}
	}

	// 当前的牌模式 和 玩家出的牌模式是否相同
	pattern := GetCardsPattern(c...)
	obj := ConvertCards(pattern, c)

	// 如果场上的牌是任意类型
	if pattern == constant.AnyPattern && !isPre {
		t.curCards = obj
		t.GetCurPlayer().SetPlayedCards(obj)
		t.Trigger(constant.GAME_PLAYER_PLAYED_CARDS)
		t.Turn()
		return nil
	}

	// 如果场上是任意牌类型，玩家选择不出
	if t.curCards.GetPattern() == constant.AnyPattern && pattern == constant.NullPattern {
		return errors.New(t.GetCurPlayer().GetName() + " must deal cards.")
	}

	// 玩家放弃出牌
	if pattern == constant.NullPattern {
		t.GetCurPlayer().SetPlayedCards(nil)
		t.Trigger(constant.GAME_PLAYER_PLAYED_CARDS)
		t.Turn()
		return nil
	}

	// 如果牌型不匹配
	if t.curCards.GetPattern() != constant.NullPattern &&
		!t.curCards.IsSamePattern(obj) &&
		obj.GetPattern() != constant.BoomCardsPattern {
		return errors.New("U dealed mismatched cards type: " + t.GetCards().GetPattern().ToString() + " and " + pattern.ToString())
	}

	// 玩家出的牌是否比场上的牌大
	if t.curCards.GetPattern() != constant.NullPattern &&
		!obj.IsGreater(t.curCards) {
		return errors.New("U should deal greater than last player's cards: " + cards.NewCardsList(t.curCards.GetCards()...).ToString())
	}

	// 比较牌
	if !obj.IsGreater(t.curCards) {
		return errors.New("U should deal greater than last player's cards: " + cards.NewCardsList(t.curCards.GetCards()...).ToString())
	}

	// 是否出牌成功
	p := t.players[t.curIndex]
	if err := p.DealCards(c...); err != nil {
		return err
	}
	// 如果玩家出完牌，游戏结束
	if !p.HasCards() {
		t.SetState(constant.GameEnd)
		return nil
	}
	// 出牌成功, 则设置 curPattern 和 curCards
	t.curCards = obj

	t.GetCurPlayer().SetPlayedCards(obj)
	t.Trigger(constant.GAME_PLAYER_PLAYED_CARDS)

	t.Turn()
	return nil
}

// 出牌权交下一位玩家
func (t *Game) Turn() {
	if t.curIndex == 2 {
		t.curIndex = -1
	}
	t.curIndex++
	// 触发钩子
	t.Trigger(constant.GAME_TURN_CHANGED, t.state)
}

// 获取当前有出牌权的玩家
func (t *Game) GetCurPlayer() constant.PlayerInterface {
	if len(t.players) < 3 {
		fmt.Println("Game has been over, because game dont have 3 player enough.")
		return nil
	}
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
// ? 这里有两种情况，一种是提前结束，一种是完美结束
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

// 把player给hooked
func (t *Game) hookPlayer(p constant.PlayerInterface) {
	cb := func(args ...interface{}) {
		t.Trigger(constant.GAME_PLAYER_STATE_CHANGED)
	}
	p.On(constant.PLAYER_STATE_CHANGED, cb)
}

// 清除玩家的钩子
func (t *Game) clearPlayerHooks(p constant.PlayerInterface) {
	p.Off(constant.PLAYER_STATE_CHANGED)
}

// 重开游戏
func (t *Game) Restart() {
	// t.Clear()
	// g = &Game{
	// 	name:          name,
	// 	cardsBoot:     cards.NewCardsBoot(),
	// 	players:       []constant.PlayerInterface{},
	// 	lordCards:     []*cards.Card{},
	// 	curIndex:      0,
	// 	curCards:      compare.NewAnyCards(),
	// 	debug:         false,
	// 	isCalledLoard: false,
	// 	state:         constant.GameReady,
	// 	hooks:         map[constant.GameHookType][]func(...interface{}){},
	// 	isSortCards:   false,
	// }

	t.players = []constant.PlayerInterface{}
	t.lordCards = []*cards.Card{}
	t.curIndex = 0
	t.curCards = compare.NewAnyCards()
	t.isCalledLoard = false
	t.state = constant.GameReady
	t.isSortCards = false

}
