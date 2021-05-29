package scene

import (
	"fmt"
	"testing"
)

func TestSceneFlow(t *testing.T) {
	f := NewSceneFlow(nil)
	f.AddHandler(func(cxt *SceneFlow) {
		fmt.Println(1)
		cxt.Next()
	})
	f.AddHandler(func(cxt *SceneFlow) {
		fmt.Println(2)
		cxt.Next()
	})
	f.AddHandler(func(cxt *SceneFlow) {
		fmt.Println(3)
		cxt.Next()
	})
	f.Start()
}
