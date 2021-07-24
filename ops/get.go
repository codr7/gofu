package ops

import (
	"github.com/codr7/gofu"
)

type TGet struct {
	index int
}

func Get(idx int) TGet {
	return TGet{index: idx}
}

func (self TGet) Eval(pc *int, calls *gofu.CallStack, stack *gofu.Stack) error {
	stack.Push(stack.Get(self.index))
	*pc++
	return nil
}
