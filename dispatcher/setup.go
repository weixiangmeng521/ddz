package dispatcher

import (
	"ddz/app/api"
	"fmt"
	"log"
	"net/http"

	gosocketio "github.com/graarh/golang-socketio"
	"github.com/graarh/golang-socketio/transport"
)

func CreateServer() {
	server := gosocketio.NewServer(transport.GetDefaultWebsocketTransport())

	server.On(gosocketio.OnConnection, HandleOnConnection)
	server.On(gosocketio.OnDisconnection, HandleOnDisconnection)
	server.On(gosocketio.OnError, HandleOnError)

	//	绑定api注册事件
	for k, v := range api.EventsMap {
		server.On(k, v)
	}

	//setup http server
	serveMux := http.NewServeMux()
	serveMux.Handle("/ws/", server)

	fmt.Println("Listening: localhost:9527")
	log.Panic(http.ListenAndServe("localhost:9527", serveMux))
}

// 玩家加入游戏
var HandleOnConnection = func(c *gosocketio.Channel, args interface{}) {
	//client id is unique
	fmt.Println("New client connected, client id is ", c.Id())
	// fmt.Println(args)
	// c.Join(publicHub) // join the public channel)
	conn := api.NewConn("", publicHub, c)
	api.AddConnInfo(conn)
}

// 玩家离开游戏
var HandleOnDisconnection = func(c *gosocketio.Channel) {
	fmt.Println("New client disconnected")
	// c.Leave(publicHub) // leave the public channel
	api.LeaveRoom(c)

}

// 玩家遇到错误
var HandleOnError = func(c *gosocketio.Channel) {
	fmt.Println("Error occurs")
}
