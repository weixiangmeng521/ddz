package api

func init() {
	bind("name:set", SetName)
	bind("game:list", GameList)
	bind("game:join", JoinGame)
	bind("game:ready", ReadyGame)
	bind("game:wait", WaitGame)
	bind("game:options", GameOptions)
	bind("game:deal", GameDeal)
	bind("cards:changed", CardsChanged)
}
