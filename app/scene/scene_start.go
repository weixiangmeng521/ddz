package scene

// 开局
var StartGame = func(cxt *SceneFlow) {
	g := cxt.GetGame()

	// ? 玩家加入游戏, 现在只是模拟
	for _, p := range Players {
		cxt.Log(p.GetName() + " joined games.")
		g.JoinPlayer(p)
	}
	// 玩家不满，无法开始游戏
	if !g.CanStart() {
		cxt.Err("cannot start game, because lack of players.")
		cxt.Redo()
		return
	}
	cxt.Info("Game started successfully.")
	cxt.Next()
}
