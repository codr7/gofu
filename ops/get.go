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
	it := stack.Get(self.index)
	stack.Push(it.Type(), it.Value())
	*pc++
	return nil
}
