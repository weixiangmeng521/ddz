package api

import "ddz/app/cards"

// 游戏场上的牌
type GameCards struct {
	HideCards   map[string]int `json:"hideCards"`
	MyCards     []*cards.Card  `json:"myCards"`
	MyId        string         `json:"myId"`
	LordCards   []*cards.Card  `json:"lordCards"`
	RolesMap    map[string]int `json:"rolesMap"` // 角色map
	PlayedCards []*cards.Card  `json:"playedCards"`
}

func NewGameCards() *GameCards {
	return &GameCards{
		HideCards:   map[string]int{},
		MyCards:     []*cards.Card{},
		MyId:        "",
		LordCards:   []*cards.Card{},
		RolesMap:    map[string]int{},
		PlayedCards: []*cards.Card{},
	}
}

// 玩家选择
type PlayerOptions struct {
	Type    string            `json:"type"`
	Options map[string]string `json:"options"`
	lordMap map[string]string // 叫地主的牌
	dealMap map[string]string // 出牌的按钮
}

func NewPlayerOptions() *PlayerOptions {
	lordMap := map[string]string{
		"call_lord": "1",
		"not_call":  "0",
	}
	dealMap := map[string]string{
		"cannot_afford": "0",
		"play_cards":    "1",
	}

	return &PlayerOptions{
		Type:    "",
		Options: map[string]string{},
		lordMap: lordMap,
		dealMap: dealMap,
	}
}

// 叫地主的选择
func (t *PlayerOptions) SetCallLord() *PlayerOptions {
	t.Type = "call"
	t.Options = t.lordMap
	return t
}

// 叫地主的选择
func (t *PlayerOptions) SetPlayCards() *PlayerOptions {
	t.Type = "play"
	t.Options = t.dealMap
	return t
}

// 清牌
func (t *PlayerOptions) Clear() *PlayerOptions {
	t.Type = "play"
	t.Options = map[string]string{}
	return t
}
