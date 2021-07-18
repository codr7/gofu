package ops

import (
	"github.com/codr7/gofu"
)

type push struct {
	slot gofu.Slot
}

func Push(t gofu.Type, v interface{}) *push {
	op := new(push)
	op.slot.Init(t, v)
	return op
}

func (self push) Eval(pc int, stack *gofu.Stack) (int, error) {
	stack.Push(self.slot)
	return pc+1, nil
}
