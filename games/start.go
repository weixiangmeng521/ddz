package games

import "ddz/flow"

// 开始游戏
func Start() {
	g := NewGame()
	flow := flow.NewFlow(g)
	flow.Reset()
	flow.Start()
}
