package ops

import (
	"fmt"
	"github.com/codr7/gofu"
)

type TBindId struct {
	pos gofu.TPos
	index int
}

func BindId(pos gofu.TPos, idx int) TBindId {
	return TBindId{pos: pos, index: idx}
}

func (self TBindId) Eval(thread *gofu.TThread, pc *int) error {
	stack := thread.Stack()
	
	if stack.Empty() {
		return fmt.Errorf("Missing value to bind")
	}

	s := stack.Pop()
	thread.Set(self.index, s.Type(), s.Value())
	*pc++
	return nil
}
