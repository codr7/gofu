package ops

import (
	"github.com/codr7/gofu"
)

type TBindId struct {
	index int
}

func BindId(idx int) TBindId {
	return TBindId{index: idx}
}

func (self TBindId) Eval(pc *int, calls *gofu.CallStack, registers []gofu.Slot, stack *gofu.Stack) error {
	registers[self.index] = stack.Pop()
	*pc++
	return nil
}
