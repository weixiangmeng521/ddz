package flow

import (
	"fmt"
	"testing"
)

func TestFlow(t *testing.T) {
	f := NewFlow(nil)
	f.AddHandler(func(cxt *Flow) {
		fmt.Println(1)
		cxt.Next()
	})
	f.AddHandler(func(cxt *Flow) {
		fmt.Println(2)
		cxt.Next()
	})
	f.AddHandler(func(cxt *Flow) {
		fmt.Println(3)
		cxt.Next()
	})
	f.Start()
}
