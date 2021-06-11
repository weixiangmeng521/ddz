package scene

import (
	"ddz/app/constant"
)

// 结束
var GoodGame = func(cxt *SceneFlow) {
	g := cxt.GetGame()
	g.SetState(constant.GameEnd)

	cxt.Log("good game.")

	g.On(constant.GAME_PLAYER_LEAVED, func(i ...interface{}) {
		// 如果玩家走完了就重开
		if IsEmptyGame(g) {
			cxt.Info("game restart agin.")
			// 销毁房间，然后重建
			g.Trigger(constant.GAME_REBUILD, g)
		}
	})

	cxt.Next()
}

// 如果游戏的玩家走完了
func IsEmptyGame(g constant.GameInterface) bool {
	count := 0
	g.MapPlayers(func(i int, pi constant.PlayerInterface) {
		count++
	})
	// fmt.Println("players number :", count)
	return count == 1
}
