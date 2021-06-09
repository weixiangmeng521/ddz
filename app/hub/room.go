package hub

import (
	"ddz/app/constant"
	"ddz/app/games"
	"sync"
)

const (
	Created = "created"
	Removed = "removed"
)

type Room struct {
	set   map[string]constant.GameInterface
	hooks map[string][]func(string)
	sync.Mutex
}

func NewRoom() *Room {
	return &Room{
		hooks: map[string][]func(string){},
		set:   map[string]constant.GameInterface{},
	}
}

// 开房
func (t *Room) CreateRoom(name string, g constant.GameInterface) {
	t.Lock()
	defer t.Unlock()
	t.set[name] = g
	t.bindDestory()
	t.trigger(Created, name)
}

// 消房
func (t *Room) DelRoom(name string) {
	t.Lock()
	defer t.Unlock()
	t.bindDestory()
	delete(t.set, name)
	t.trigger(Removed, name)
}

// 获取房间信息
func (t *Room) GetRoom(name string) constant.GameInterface {
	return t.set[name]
}

// 查看所有的房子
func (t *Room) List() []string {
	list := []string{}
	for k := range t.set {
		list = append(list, k)
	}
	return list
}

// 有没有这个房子
func (t *Room) HasRoom(name string) bool {
	_, ok := t.set[name]
	return ok
}

// 绑定hook
func (t *Room) On(name string, fn func(string)) {
	if _, ok := t.hooks[name]; !ok {
		t.hooks[name] = []func(string){fn}
		return
	}
	t.hooks[name] = append(t.hooks[name], fn)
}

// 绑定hook
func (t *Room) Off(name string) {
	if _, ok := t.hooks[name]; !ok {
		return
	}
	t.hooks[name] = []func(string){}
}

// 触发钩子
func (t *Room) trigger(name string, arg string) {
	if _, ok := t.hooks[name]; !ok {
		return
	}
	for _, fn := range t.hooks[name] {
		fn(arg)
	}
}

func (t *Room) bindDestory() {
	for _, g := range t.set {
		g.Off(constant.GAME_REBUILD)
		g.On(constant.GAME_REBUILD, t.onRebuild)
	}

}

// 重新建房
func (t *Room) onRebuild(i ...interface{}) {
	if len(i) != 1 {
		return
	}
	g := i[0].(constant.GameInterface)
	for _, v := range t.set {
		if v.GetName() == g.GetName() {
			t.DelRoom(g.GetName())
		}
	}
	t.CreateRoom(g.GetName(), games.NewGame(g.GetName()))
}

// GAME_DESTORY
