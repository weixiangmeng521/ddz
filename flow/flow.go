package flow

import (
	"ddz/cards"
	"ddz/players"
)

type GameInterface interface {
	JoinPlayer(*players.Player) bool // 玩家加入游戏
	LeavePlayer(*players.Player)     // 玩家离开游戏
	CanStart() bool                  // 是否能开始
	Display()                        // 看牌
	Shuffle()                        // 洗牌
	Licensing()                      // 发牌
	CallLandlord(*players.Player)    // 叫地主
	DealCards([]*cards.Card) error   // 玩家出牌
	CompareCards([]*cards.Card) bool // 比牌
	Turn()                           // 出牌权交下一位玩家
	GetWiners() []*players.Player    // 获取游戏赢家
}

type FlowHandler func(*Flow)

type Flow struct {
	game       GameInterface
	background map[interface{}]interface{}
	handlers   []FlowHandler
	cur        int8
}

func NewFlow(game GameInterface) *Flow {
	return &Flow{
		game:       game,
		background: map[interface{}]interface{}{},
		handlers:   []FlowHandler{},
		cur:        -1,
	}
}

// 重置flow
func (t *Flow) Reset() {
	t.cur = -1
}

// 添加handler
func (t *Flow) AddHandler(h FlowHandler) {
	t.handlers = append(t.handlers, h)
}

// 进入flow
func (t *Flow) Next() {
	t.cur++
	if t.cur < int8(len(t.handlers)) {
		t.handlers[t.cur](t)
	}
}

// 开始启用这个flow
func (t *Flow) Start() {
	if len(t.handlers) == 0 {
		return
	}
	t.cur++
	t.handlers[t.cur](t)
}

// 写入值
func (t *Flow) Set(key interface{}, value interface{}) {
	t.background[key] = value
}

// 取出值
func (t *Flow) Get(key interface{}) interface{} {
	return t.background[key]
}
