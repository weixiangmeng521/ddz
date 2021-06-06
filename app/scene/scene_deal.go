package scene

import (
	"ddz/app/cards"
	"ddz/app/constant"
)

// 开始出牌
var DealCards = func(cxt *SceneFlow) {
	g := cxt.GetGame()
	n := g.GetCurPlayer().GetName()
	cxt.Warn("Land lord is %s", n)

	cxt.Info("Deal cards begin\n")

	// 触发一下turn
	ForceUpdateGameTurnChanged(g)

	g.On(constant.GAME_PLAYER_PLAYED_CARDS, func(i ...interface{}) {
		// playedCards := g.GetPlayedCards()
		cp := g.GetCurPlayer().GetPlayedCards()
		if cp != nil {
			cxt.Info("%s >>> [%s]: %s", g.GetCurPlayer().GetName(), cp.GetPattern().ToString(), cards.NewCardsList(cp.GetCards()...).ToString())
			return
		}
		cxt.Info("%s >>> not play", g.GetCurPlayer().GetName())

	})

	// return
	// // 清除钩子
	// g.Off(constant.GAME_TURN_CHANGED)
	// cxt.Next()
}

// for !g.HasGoodGame() {
// 	p := g.GetCurPlayer()
// 	cxt.Info("%s turn to play cards.", p.GetName())
// 	cxt.Info("%s", p.CheckCards())
// 	fmt.Println()

// 	return

// 	// 玩家出牌
// 	// g.DealCards()

// 	// 轮到下一个玩家出牌
// 	// g.Turn()
// }
