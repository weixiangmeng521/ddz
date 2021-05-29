package scene

import (
	"ddz/app/constant"
	"ddz/app/players"
)

// 游戏玩家
var Players = []constant.PlayerInterface{
	players.NewPlayer("Kenny"),
	players.NewPlayer("Kyle"),
	players.NewPlayer("Cartman"),
}

// 初始化
func CreateSceneFlow(g constant.GameInterface) {
	t := NewSceneFlow(g)
	t.AddHandler(StartGame)
	t.AddHandler(ShuffleCards)
	t.AddHandler(CallLandlord)
	t.AddHandler(DealCards)
	t.AddHandler(GoodGame)
	t.Reset()
	t.Start()
}
