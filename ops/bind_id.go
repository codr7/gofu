package ops

import (
	"github.com/codr7/gofu"
)

type TBindId struct {
	index int
}

func BindId(index int) *TBindId {
	op := new(TBindId)
	op.index = index
	return op
}

func (self TBindId) Eval(pc int, stack *gofu.Stack) (int, error) {
	it := stack.Pop()
	stack.Set(self.index, it.Type(), it.Value())
	return pc+1, nil
}
