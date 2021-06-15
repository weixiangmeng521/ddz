package api

func init() {
	bind("name:set", SetName)
	bind("game:list", GameList)
	bind("game:join", JoinGame)
	bind("game:ready", ReadyGame)
	bind("game:wait", WaitGame)
	bind("game:options", GameOptions)
	bind("game:options[confirm]", GameOptionsConfirm)
	bind("game:deal", GameDeal)
	bind("game:good_game", GoodGame)
	bind("cards:changed", CardsChanged)
}
