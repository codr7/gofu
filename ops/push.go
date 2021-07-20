package ops

import (
	"github.com/codr7/gofu"
)

type TPush struct {
	slot gofu.Slot
}

func Push(t gofu.Type, v interface{}) *TPush {
	op := new(TPush)
	op.slot.Init(t, v)
	return op
}

func (self TPush) Eval(pc int, stack *gofu.Stack) (int, error) {
	stack.Push(self.slot)
	return pc+1, nil
}
