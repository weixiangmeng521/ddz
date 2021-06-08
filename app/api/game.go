package api

import (
	"ddz/app/constant"
	"ddz/app/players"
	"fmt"

	gosocketio "github.com/graarh/golang-socketio"
)

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
var GameList = func(c *gosocketio.Channel, i interface{}) error {
	msg := NewMessage().Success()
	msg.SetData(hubs.List())
	c.Emit("game:list", msg)
	return nil
}

// 游戏 api
var JoinGame = func(c *gosocketio.Channel, i interface{}) error {
	msg := ParseMessage(i)
	r := msg.GetData().(string) // room name
	game := hubs.GetRoom(r)

	if game.GetState() != constant.GameReady {
		c.Emit("game:join", msg.Error().SetMessage("game is runing now. u cannot join. state: "+game.GetState().ToString()))
		return nil
	}

	// 获取连接信息
	info := GetConnInfo(c)
	// 加入房间广播
	c.Join(r)

	if info == nil {
		c.Emit("game:join", msg.Error().SetMessage("you didnt have name."))
		return nil
	}

	// 加入游戏
	p := players.NewPlayer(info.name)
	p.SetRoom(r)
	if res := game.JoinPlayer(p); !res {
		msg.Error().SetMessage("the game is full.")
		c.Emit("game:join", msg)
		return nil
	}

	SetPlayer(c, p, func() {
		c.BroadcastTo(r, "game:join", msg.Success().SetData(getPlayersStatus(c)))
	})
	return nil
}

// 开始游戏
var ReadyGame = func(c *gosocketio.Channel, i interface{}) error {
	msg := NewMessage()
	p := GetPlayer(c)
	if p == nil {
		c.Emit("game:ready", msg.Error().SetMessage("cannot ready game, refresh first."))
		fmt.Println("Nil Pointer player.")
		return nil
	}
	r := p.GetRoom() // room
	if p == nil {
		msg.Error().SetMessage("u r not player.")
		c.Emit("game:ready", msg)
		return nil
	}

	p.SetState(constant.Already)
	c.BroadcastTo(r, "game:ready", msg.Success().SetData(getPlayersStatus(c)))
	return nil
}

// 等待游戏
var WaitGame = func(c *gosocketio.Channel, i interface{}) error {
	msg := NewMessage()
	p := GetPlayer(c) // player
	r := p.GetRoom()  // room
	if p == nil {
		msg.Error().SetMessage("u r not player.")
		c.Emit("game:wait", msg)
		return nil
	}
	p.SetState(constant.Waiting)

	c.BroadcastTo(r, "game:wait", msg.Success().SetData(getPlayersStatus(c)))
	return nil
}

// [订阅] 发牌 cards:changed
var CardsChanged = func(c *gosocketio.Channel, i interface{}) error {
	p := GetPlayer(c) // player
	if p == nil {
		fmt.Println("Nil pointer Player: ", p)
		return nil
	}
	g := hubs.GetRoom(p.GetRoom())
	if g == nil {
		fmt.Println("Nil pointer Game: ", p.GetName())
		return nil
	}
	msg := NewMessage()

	// 当牌发生变化时触发
	g.On(constant.GAME_CARDS_CHANGED, func(i ...interface{}) {
		var target *gosocketio.Channel
		arr := c.List(p.GetRoom()) // 每个连接
		gc := NewGameCards()       // 返回给前台的数据

		roleMaps := map[string]int{}
		for _, conn := range arr {
			player := GetPlayer(conn)
			// 设置用户信息
			if player != nil && player.IsLord() {
				roleMaps[player.GetName()] = 1
			}
			if player != nil && !player.IsLord() {
				roleMaps[player.GetName()] = 0
			}
			// 让玩家看对手牌的数量
			if player != nil && conn.Id() != p.GetName() {
				gc.HideCards[conn.Id()] = len(player.GetCards())
			}
			// 获取自己的牌
			if conn.Id() == p.GetName() {
				// 这里响应每个玩家的卡牌，还有地主牌
				gc.MyCards = p.GetCards()
				target = conn
				continue
			}
		}
		gc.LordCards = g.GetLordCards()
		gc.RolesMap = roleMaps
		gc.MyId = p.GetName()
		gc.PlayedCards = g.GetPlayedCards()

		if target != nil {
			target.Emit("cards:changed", msg.Success().SetData(gc))
		}
	})
	return nil
}

// [订阅] 当前玩家的操作 game:options
var GameOptions = func(c *gosocketio.Channel, i interface{}) error {
	g := GetGame(c)
	if g == nil {
		fmt.Println("nil pointer err: GetGame.")
		return nil
	}
	// 游戏出牌权发生改变
	g.On(constant.GAME_TURN_CHANGED, func(i ...interface{}) {
		state := g.GetState()

		p := g.GetCurPlayer()
		if GetConnInfo(c) == nil || p == nil {
			fmt.Println("nil pointer err: GetConnInfo || game.GetCurPlayer")
			return
		}

		// 叫地主时
		if GetConnInfo(c).GetName() == p.GetName() && state == constant.GameStarted {
			c.Emit("game:options", NewPlayerOptions().SetCallLord())
		}
		// 叫地主后，出牌时
		if GetConnInfo(c).GetName() == p.GetName() && state == constant.GameCalled {
			c.Emit("game:options", NewPlayerOptions().SetPlayCards())
		}
	})
	// 游戏状态发生改变 ！暂时取消这块的
	g.On(constant.GAME_STATE_CHANGED, func(i ...interface{}) {
		p := g.GetCurPlayer()
		if GetConnInfo(c) == nil || p == nil {
			fmt.Println("nil pointer err: GetConnInfo || game.GetCurPlayer")
			return
		}
		if g.GetState() == constant.GameCalled {
			c.Emit("game:options", NewPlayerOptions().Clear())
		}
	})

	return nil
}

// 玩家选择的操作
var GameDeal = func(c *gosocketio.Channel, i interface{}) error {
	msg := ParseMessage(i)
	if msg == nil {
		fmt.Println("Parse message err!")
		return nil
	}
	g := GetGame(c)
	if g == nil {
		fmt.Println("Nil pointer err: GetGame")
		return nil
	}
	// 给出牌玩家操作
	if g.GetCurPlayer().GetName() == GetPlayer(c).GetName() {
		m := msg.Data.(map[string]interface{})
		tp := m["type"].(string)    // type
		opt := m["option"].(string) // options

		// 叫地主
		if tp == "call" && opt == "1" && g.GetState() == constant.GameStarted {
			fmt.Printf("%s call lord.\n", g.GetCurPlayer().GetName())
			g.GetCurPlayer().CallLord()
			g.Turn()
		}

		// 不叫
		if tp == "call" && opt == "0" && g.GetState() == constant.GameStarted {
			fmt.Printf("%s not call.\n", g.GetCurPlayer().GetName())
			g.GetCurPlayer().NotCall()
			g.Turn()
		}

		// 响应叫地主
		if tp == "call" {
			c.Emit("game:deal", NewMessage().Success())
		}

		// 获取牌
		if tp == "play" && g.GetState() == constant.GameCalled {
			msg.GetData()
		}

		// 出牌
		if tp == "play" && opt == "1" && g.GetState() == constant.GameCalled {

		}

		// 不出
		if tp == "play" && opt == "0" && g.GetState() == constant.GameCalled {
			// g.DealCards()
		}

		return nil
	}

	c.Emit("game:deal", NewMessage().Error().SetMessage("It's not your turn."))
	return nil
}
