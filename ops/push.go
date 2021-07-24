package ops

import (
	"github.com/codr7/gofu"
)

type TPush struct {
	slot gofu.Slot
}

func Push(t gofu.Type, v interface{}) TPush {
	var op TPush
	op.slot.Init(t, v)
	return op
}

func (self TPush) Eval(pc *int, calls *gofu.CallStack, stack *gofu.Stack) error {
	stack.Push(self.slot)
	*pc++
	return  nil
}
