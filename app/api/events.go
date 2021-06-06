package api

import (
	gosocketio "github.com/graarh/golang-socketio"
)

func init() {
	bind("name:set", SetName)
	bind("room:list", RoomList)
	bind("game:join", JoinGame)
	bind("game:ready", ReadyGame)
	bind("game:wait", WaitGame)
	bind("game:options", GameOptions)
	bind("game:deal", GameDeal)
	bind("cards:changed", CardsChanged)
}

// 设置用户名
// ! 暂时模拟获取用户名
var SetName = func(c *gosocketio.Channel, i interface{}) error {
	// if i == nil {
	msg := NewMessage()
	info := GetConnInfo(c)
	if info == nil {
		c.Emit("name:set", msg.Error())
		return nil
	}
	info.SetName(c.Id())
	c.Emit("name:set", msg.Success())
	return nil

	// msg := ParseMessage(i)
	// info := GetConnInfo(c)
	// if info == nil {
	// 	c.Emit("name:set", msg.Error())
	// 	return nil
	// }
	// info.SetName(msg.GetData().(string))

	// msg = NewMessage()
	// c.Emit("name:set", msg.Success())
	// return nil
}

// 获取房间列表
var RoomList = func(c *gosocketio.Channel, i interface{}) error {
	msg := NewMessage().Success()
	msg.SetData(hubs.List())
	c.Emit("room:list", msg)
	return nil
}
