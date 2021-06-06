package scene

import (
	"ddz/app/constant"
	"fmt"
)

// 是否有人叫了地主
func hasSomebodyCalled(g constant.GameInterface) bool {
	has := false
	g.MapPlayers(func(i int, pi constant.PlayerInterface) {
		if !has && pi.HasCalledLord() {
			has = true
		}
	})
	return has
}

// 获取地主
func getLandlord(g constant.GameInterface) constant.PlayerInterface {
	var p constant.PlayerInterface
	g.MapPlayers(func(i int, pi constant.PlayerInterface) {
		if pi.IsLord() {
			p = pi
		}
	})
	return p
}

// 强制更新game的出牌权，响应给前端
func ForceUpdateGameTurnChanged(g constant.GameInterface) {
	if g == nil {
		fmt.Println("Nil pointer err GameInterface.")
		return
	}
	g.Trigger(constant.GAME_TURN_CHANGED, g.GetState())
}

// 叫地主
var CallLandlord = func(cxt *SceneFlow) {
	g := cxt.GetGame()

	// 触发一下turn
	ForceUpdateGameTurnChanged(g)

	num := 0 // turn的次数
	// 叫地主
	g.On(constant.GAME_TURN_CHANGED, func(i ...interface{}) {
		num++

		// cxt.Info(">>> %d", num)
		// 3伦下来没人叫地主
		if num == 3 && !hasSomebodyCalled(g) {
			cxt.Warn("Nobody call landLord, shuffle cards.")
			g.Shuffle()
			g.Licensing()
			g.Display()
			num = 0
			return
		}

		// 如果叫了 3伦了
		if num == 3 && hasSomebodyCalled(g) {
			g.CallLandlord()
			lord := getLandlord(g)
			if lord == nil {
				num = 0
				cxt.Err("System err: call landloard err.")
				// cxt.Redo()
				return
			}
			// 新新牌了，地主拿到地主牌
			g.Trigger(constant.GAME_CARDS_CHANGED)
			// 出牌权给地主
			g.ChangeTurn2Lord()

			cxt.Info("%s has became to landloard.", lord.GetName())
			g.Display()

			cxt.Next()
			return
		}
	})
}
