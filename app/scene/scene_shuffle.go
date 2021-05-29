package scene

// 洗牌
var ShuffleCards = func(cxt *SceneFlow) {
	g := cxt.GetGame()

	cxt.Info("Shuffle cards successfully.")
	g.Shuffle()
	g.Licensing()
	g.Display()

	cxt.Next()
}
