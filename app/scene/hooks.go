package scene

import (
	"sync"
)

// 钩子
type Hooks struct {
	hooks map[string][]func(...interface{})
	sync.Mutex
}

func NewHooks() *Hooks {
	return &Hooks{}
}

func (t *Hooks) init() {
	if t.hooks == nil {
		t.hooks = map[string][]func(...interface{}){}
	}
}

// 绑定钩子
func (t *Hooks) On(name string, cb func(...interface{})) {
	t.Lock()
	defer t.Unlock()

	t.init()
	if _, ok := t.hooks[name]; !ok {
		t.hooks[name] = []func(...interface{}){cb}
		return
	}
	t.hooks[name] = append(t.hooks[name], cb)
}

// 取消绑定
func (t *Hooks) Off(name string) {
	t.Lock()
	defer t.Unlock()

	t.init()
	if _, ok := t.hooks[name]; !ok {
		return
	}
	t.hooks[name] = []func(...interface{}){}
}

// 清除全部的钩子
func (t *Hooks) Clear() {
	t.hooks = nil
}

// 触发钩子
func (t *Hooks) Trigger(name string, args ...interface{}) {
	t.init()
	if _, ok := t.hooks[name]; !ok {
		return
	}
	for _, fn := range t.hooks[name] {
		fn(args...)
	}
}
