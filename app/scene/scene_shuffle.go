package scene

import "ddz/app/constant"

// 洗牌
var ShuffleCards = func(cxt *SceneFlow) {
	g := cxt.GetGame()

	g.SetState(constant.GameStarted)

	cxt.Info("Shuffle cards successfully.")
	g.Shuffle()
	g.Licensing()
	g.Display()

	cxt.Next()
}
