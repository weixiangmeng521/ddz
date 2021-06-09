package constant

import (
	"ddz/app/cards"
)

type RoomInterface interface {
	CreateRoom(string, GameInterface) // 创建房间
	DelRoom(string)                   // 删除房间
	GetRoom(string) GameInterface     // 获取房间的game
	List() []string                   // 房间列表
	HasRoom(string) bool              // 是否存在这个房间
	On(string, func(string))          // 绑定事件
	Off(string)                       // 解绑
}

type GameInterface interface {
	SetName(string)                        // 设置游戏名称
	GetName() string                       // 获取游戏名称
	GetState() GameState                   // 获取游戏当前状态
	Debug()                                // debug模式
	GetCards() CardsCompareInterface       // 获取出牌信息
	GetPlayedCards() []*cards.Card         // 获取场上的牌
	SetState(GameState)                    // 设置当前的游戏状态
	JoinPlayer(PlayerInterface) bool       // 玩家加入游戏
	LeavePlayer(PlayerInterface)           // 玩家离开游戏
	CanStart() bool                        // 是否能开始
	ChangeTurn2Lord()                      // 出牌权交给地主
	Display()                              // 看牌
	Shuffle()                              // 洗牌
	Licensing()                            // 发牌
	CallLandlord() bool                    // 叫地主，如果没人叫，就继续让他们叫
	DealCards([]*cards.Card) error         // 玩家出牌
	Turn()                                 // 出牌权交下一位玩家
	GetWiners() []PlayerInterface          // 获取游戏赢家
	GetCurPlayer() PlayerInterface         // 获取当前有出牌权的玩家
	HasGoodGame() bool                     // 是不是gg了
	MapPlayers(func(int, PlayerInterface)) // 便利玩家
	On(string, func(...interface{}))       // 绑定事件
	Trigger(string, ...interface{})        // 触发钩子
	Off(string)                            // 取消绑定
	GetLordCards() []*cards.Card           // 获取最后3张牌
}

type PlayerInterface interface {
	GetGame() GameInterface                // 玩家所在的游戏
	SetGame(GameInterface)                 // 设置玩家所在的游戏
	GetName() string                       //获取玩家名称
	SetLord()                              // 把玩家变成地主
	Clear()                                // 清空玩家的手牌
	SetFarmer()                            // 把玩家变成龙鸣
	IsLord() bool                          // 看玩家是不是龙鸣
	IsFarmer() bool                        // 看玩家是不是地主
	SetPlayedCards(CardsCompareInterface)  // 设置当前出的牌
	GetPlayedCards() CardsCompareInterface // 获取当前出的牌
	AcceptCards(cards ...*cards.Card)      // 获取牌
	DealCards(cards ...*cards.Card) error  // 出一堆牌
	DealCard(card *cards.Card) error       // 出一张牌
	SortCards()                            // 给牌排序
	HasCards() bool                        // 看玩家还有没有手牌
	CheckCards() string                    // 看玩家手上的牌
	GetCards() []*cards.Card               // 获取卡牌对象
	CallLord()                             // 叫地主
	NotCall()                              // 不叫
	HasCalledLord() bool                   // 是否叫了地主
	HasWinned() bool                       // 玩家是不是赢了
	SetState(StateType)                    // 设置玩家状态
	GetState() StateType                   // 获取玩家状态
	SetRoom(string)                        // 设置房间
	GetRoom() string                       // 获取房间
	On(string, func(...interface{}))       // 绑定事件
	Off(string)                            // 取消绑定
}

// 前置操作， 比较牌时，检查牌是否合法
type CardsCompareInterface interface {
	GetCards() []*cards.Card
	GetPattern() CardsPattern
	IsSamePattern(CardsCompareInterface) bool
	IsGreater(CardsCompareInterface) bool
}

type MessageInterface interface {
	Success() MessageInterface
	Error() MessageInterface
	Warn() MessageInterface
	SetCode(int8) MessageInterface
	GetCode() int8
	SetMessage(string) MessageInterface
	GetMessage() string
	SetData(interface{}) MessageInterface
	GetData() interface{}
}
