package api

import (
	"ddz/app/constant"
	"ddz/app/games"
	h "ddz/app/hub"
	"ddz/app/players"
	"fmt"

	gosocketio "github.com/graarh/golang-socketio"
)

type EventHandler func(*gosocketio.Channel, interface{}) error

var EventsMap = map[string]EventHandler{}

// bind websocket events
func bind(m string, cb EventHandler) {
	EventsMap[m] = cb
}

// 加入游戏的玩家一个游戏名称
var (
	connections = map[*gosocketio.Channel]*ConnInfo{}
	delConnChan = make(chan *gosocketio.Channel)
	addConnChan = make(chan *ConnInfo)
)

func init() {
	go func() {
		for {
			select {
			case c := <-delConnChan:
				delete(connections, c)

			case c := <-addConnChan:
				connections[c.GetConn()] = c

			case m := <-addPlayerUsersChan:
				playerUsers[m.channel] = m.player
				for _, cb := range m.cb {
					cb()
				}

			case m := <-delPlayerUsersChan:
				for k := range m {
					delete(playerUsers, k)
				}
			}
		}
	}()

	createRooms()
}

// 获取连接的额外信息
func GetConnInfo(c *gosocketio.Channel) *ConnInfo {
	return connections[c]
}

// 添加连接额外信息
func AddConnInfo(c *ConnInfo) {
	addConnChan <- c
}

// 删除连接的额外信息
func DelConnInfo(c *gosocketio.Channel) {
	c.Leave("hub")
	delConnChan <- c
}

// 连接的额外信息
type ConnInfo struct {
	name string
	room string
	conn *gosocketio.Channel
}

func NewConn(name string, room string, conn *gosocketio.Channel) *ConnInfo {
	return &ConnInfo{name, room, conn}
}

func (t *ConnInfo) SetName(s string) {
	t.name = s
}

func (t *ConnInfo) GetName() string {
	return t.name
}

func (t *ConnInfo) SetRoom(s string) {
	t.room = s
}

func (t *ConnInfo) GetRoom() string {
	return t.room
}

func (t *ConnInfo) GetConn() *gosocketio.Channel {
	return t.conn
}

var (
	playerUsers        = map[*gosocketio.Channel]*players.Player{}
	addPlayerUsersChan = make(chan *customPlayerCallInfo)
	delPlayerUsersChan = make(chan map[*gosocketio.Channel]*players.Player)
)

// 获取玩家信息
func GetPlayer(c *gosocketio.Channel) *players.Player {
	return playerUsers[c]
}

// 自定义的player信息
type customPlayerCallInfo struct {
	channel *gosocketio.Channel
	player  *players.Player
	cb      []func()
}

// 设置玩家信息
func SetPlayer(c *gosocketio.Channel, p *players.Player, cb ...func()) {
	m := &customPlayerCallInfo{
		channel: c,
		player:  p,
		cb:      cb,
	}
	addPlayerUsersChan <- m
}

// 删除玩家信息
func DelPlayer(c *gosocketio.Channel, p *players.Player) {
	m := map[*gosocketio.Channel]*players.Player{}
	m[c] = p
	delPlayerUsersChan <- m
}

// 根据连接获取游戏handle
func GetGame(c *gosocketio.Channel) constant.GameInterface {
	p := GetPlayer(c)
	if p == nil {
		fmt.Println("Nil pointer err: GetPlayer()")
		return nil
	}
	if p.GetGame() != nil {
		return p.GetGame()
	}
	r := p.GetRoom()
	return hubs.GetRoom(r)
}

// 获取玩家状态列表
func getPlayersStatus(c *gosocketio.Channel) map[string]string {
	m := map[string]string{}
	p := GetPlayer(c)
	if p == nil {
		return m
	}
	g := hubs.GetRoom(p.GetRoom())
	g.MapPlayers(func(i int, v constant.PlayerInterface) {
		m[v.GetName()] = v.GetState().ToString()
	})
	return m
}

// *** 房间
var (
	hubs = h.NewRoom()
)

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
