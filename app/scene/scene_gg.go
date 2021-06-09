package scene

import (
	"ddz/app/constant"
)

// 结束
var GoodGame = func(cxt *SceneFlow) {
	// 清除钩子
	// clearHooks := func(g constant.GameInterface) {
	// 	g.Off(constant.GAME_PLAYER_LEAVED)
	// }

	g := cxt.GetGame()
	g.SetState(constant.GameEnd)

	cxt.Log("good game.")

	g.On(constant.GAME_PLAYER_LEAVED, func(i ...interface{}) {
		if IsEmptyGame(g) {
			cxt.Info("game restart agin.")

			// TODO: 销毁房间，然后重建
			g.Trigger(constant.GAME_REBUILD, g)
			// clearHooks(g)
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
