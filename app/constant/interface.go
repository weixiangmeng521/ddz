package constant

import "ddz/app/cards"

type GameInterface interface {
	JoinPlayer(PlayerInterface) bool // 玩家加入游戏
	LeavePlayer(PlayerInterface)     // 玩家离开游戏
	CanStart() bool                  // 是否能开始
	Display()                        // 看牌
	Shuffle()                        // 洗牌
	Licensing()                      // 发牌
	CallLandlord() bool              // 叫地主，如果没人叫，就继续让他们叫
	DealCards([]*cards.Card) error   // 玩家出牌
	Turn()                           // 出牌权交下一位玩家
	GetWiners() []PlayerInterface    // 获取游戏赢家
	GetCurPlayer() PlayerInterface   // 获取当前有出牌权的玩家
	HasGoodGame() bool               // 是不是gg了
}

type PlayerInterface interface {
	GetName() string                      //获取玩家名称
	SetLord()                             // 把玩家变成地主
	Clear()                               // 清空玩家的手牌
	SetFarmer()                           // 把玩家变成龙鸣
	IsLord() bool                         // 看玩家是不是龙鸣
	IsFarmer() bool                       // 看玩家是不是地主
	AcceptCards(cards ...*cards.Card)     //获取牌
	DealCards(cards ...*cards.Card) error // 出一堆牌
	DealCard(card *cards.Card) error      // 出一张牌
	SortCards()                           // 给牌排序
	HasCards() bool                       // 看玩家还有没有手牌
	CheckCards() string                   // 看玩家手上的牌
	CallLord()                            // 叫地主
	NotCall()                             // 不叫
	HasCalledLord() bool                  // 是否叫了地主
	HasWinned() bool                      // 玩家是不是赢了
}
