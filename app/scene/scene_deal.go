package scene

import (
	"fmt"
)

// 开始出牌
var DealCards = func(cxt *SceneFlow) {
	cxt.Info("Deal cards begin")
	g := cxt.GetGame()

	for !g.HasGoodGame() {
		p := g.GetCurPlayer()
		cxt.Info("%s turn to play cards.", p.GetName())
		cxt.Info("%s", p.CheckCards())
		fmt.Println()

		return

		// 玩家出牌
		// g.DealCards()

		// 轮到下一个玩家出牌
		// g.Turn()
	}

	cxt.Next()
}
