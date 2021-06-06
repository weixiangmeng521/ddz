package api

import (
	"ddz/app/constant"
	"ddz/app/games"
	h "ddz/app/hub"

	gosocketio "github.com/graarh/golang-socketio"
)

var (
	hubs constant.RoomInterface
)

func init() {
	hubs = h.NewRoom()
	createRooms()
}

// ? 模拟建房
// 绑定钩子
func createRooms() {
	createRoom("test1")
	createRoom("test2")
}

func createRoom(name string) {
	hubs.CreateRoom(name, games.NewGame(name))
}

// 玩家提前离场，有钩子会触发
func LeaveRoom(c *gosocketio.Channel) {
	g := GetGame(c)
	if g != nil {
		g.Trigger(constant.GAME_PLAYER_LEAVED, GetPlayer(c))
	}

	DelConnInfo(c) // 删除连接信息
	p := GetPlayer(c)
	DelPlayer(c, p) // 删除玩家身份
	if p == nil {
		return
	}
	//删除游戏中的玩家
	r := p.GetRoom()
	if r != "" {
		g := hubs.GetRoom(r)
		g.LeavePlayer(p)
	}
}
