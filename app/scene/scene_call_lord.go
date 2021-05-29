package scene

// 叫地主
var CallLandlord = func(cxt *SceneFlow) {
	g := cxt.GetGame()

	// ? 叫了地主的玩家列表, 现在只是模拟
	calledList := []string{
		"Kenny",
	}

	// 设置抢地主的玩家
	for _, p := range Players {
		for _, v := range calledList {
			if p.GetName() == v {
				p.CallLord()
				cxt.Info("%s has became to landloard.", p.GetName())
			}
		}
	}

	// 如果没人叫地主，就卡这个环节，一直不停叫
	if !g.CallLandlord() {
		cxt.Warn("Nobody call landLord. repeat")
		cxt.Redo()
		return
	}

	// 看牌
	g.Display()

	cxt.Next()
}
