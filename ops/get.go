package ops

import (
	"github.com/codr7/gofu"
)

type TGet struct {
	index int
}

func Get(index int) *TGet {
	op := new(TGet)
	op.index = index
	return op
}

func (self TGet) Eval(pc int, stack *gofu.Stack) (int, error) {
	stack.Push(stack.Get(self.index))
	return pc+1, nil
}
