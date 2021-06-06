package scene

import (
	"ddz/app/constant"
)

// 游戏开局
// 玩家在房间等待，等到3个人
// 等所有玩家状态是already，就开始游戏
var StartGame = func(cxt *SceneFlow) {
	g := cxt.GetGame()

	// 清空钩子
	clearHooks := func() {
		g.Off(constant.GAME_JOINED_PLYAER)
		g.Off(constant.GAME_PLAYER_STATE_CHANGED)
	}

	// 加入游戏时触发
	g.On(constant.GAME_JOINED_PLYAER, func(i ...interface{}) {
	})

	// 玩家状态改变时触发
	g.On(constant.GAME_PLAYER_STATE_CHANGED, func(i ...interface{}) {
		if g.CanStart() {
			cxt.Info("Room [%s] created started successfully.", cxt.game.GetName())
			clearHooks()
			cxt.Next()
		}
	})

	// 当某个玩家离开游戏
	g.On(constant.GAME_PLAYER_LEAVED, func(i ...interface{}) {
		// 如果游戏在准备阶段，直接无视玩家退出
		if g.GetState() == constant.GameReady {
			return
		}

		cxt.End()
	})
}
