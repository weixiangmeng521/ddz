package games

import "ddz/app/cards"

func GetDebugCardsBoot() *cards.CardsBoot {
	boot := cards.NewCardsBoot()
	boot.Init()

	return boot
}
