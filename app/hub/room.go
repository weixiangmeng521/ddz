package hub

import (
	"ddz/app/constant"
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
	t.trigger(Created, name)
}

// 消房
func (t *Room) DelRoom(name string) {
	t.Lock()
	defer t.Unlock()
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
