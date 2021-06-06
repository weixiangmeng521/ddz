package scene

import "ddz/app/constant"

// 结束
var GoodGame = func(cxt *SceneFlow) {
	g := cxt.GetGame()
	g.SetState(constant.GameEnd)

	cxt.Log("good game.")

	cxt.Next()
}
