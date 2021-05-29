package scene

import (
	"ddz/app/constant"
	"fmt"

	"github.com/gookit/color"
)

type SceneFlowHandler func(*SceneFlow)

type SceneFlow struct {
	game       constant.GameInterface
	background map[interface{}]interface{}
	handlers   []SceneFlowHandler
	cur        int8
}

func NewSceneFlow(game constant.GameInterface) *SceneFlow {
	return &SceneFlow{
		game:       game,
		background: map[interface{}]interface{}{},
		handlers:   []SceneFlowHandler{},
		cur:        -1,
	}
}

// 重置SceneFlow
func (t *SceneFlow) Reset() {
	t.cur = -1
}

// 添加handler
func (t *SceneFlow) AddHandler(h SceneFlowHandler) {
	t.handlers = append(t.handlers, h)
}

// 进入SceneFlow
func (t *SceneFlow) Next() {
	t.cur++
	if t.cur < int8(len(t.handlers)) {
		t.handlers[t.cur](t)
	}
}

// 重新跑
func (t *SceneFlow) Redo() {
	t.handlers[t.cur](t)
}

// 开始启用这个SceneFlow
func (t *SceneFlow) Start() {
	if len(t.handlers) == 0 {
		return
	}
	t.cur++
	t.handlers[t.cur](t)
}

// 写入值
func (t *SceneFlow) Set(key interface{}, value interface{}) {
	t.background[key] = value
}

// 取出值
func (t *SceneFlow) Get(key interface{}) interface{} {
	return t.background[key]
}

// 获取game
func (t *SceneFlow) GetGame() constant.GameInterface {
	return t.game
}

// 错误信息
func (t *SceneFlow) Err(f string, a ...interface{}) {
	// fmt.Printf("\033[1;31;40m%s\033[0m\n", a)
	color.Error.Printf(f+"\n", a...)

}

// 正常输出
func (t *SceneFlow) Log(f string, a ...interface{}) {
	fmt.Printf(f+"\n", a...)
}

// 信息
func (t *SceneFlow) Info(f string, a ...interface{}) {
	color.Info.Printf(f+"\n", a...)
	// fmt.Printf("\033[1;43;40m%s\033[0m\n", a)
}

// 提示输出
func (t *SceneFlow) Warn(f string, a ...interface{}) {
	// fmt.Printf("\033[1;31;40m%s\033[0m\n", a)
	color.Info.Printf(f+"\n", a...)

}
