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

func (self TBindId) Eval(pc *int, calls *gofu.CallStack, stack *gofu.Stack) error {
	it := stack.Pop()
	stack.Set(self.index, it.Type(), it.Value())
	*pc++
	return nil
}
