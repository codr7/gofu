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

func (self TBindId) Eval(pc *int, calls *gofu.CallStack, registers []gofu.Slot, stack *gofu.Stack) error {
	if stack.Empty() {
		return fmt.Errorf("Missing value to bind")
	}
	
	registers[self.index] = stack.Pop()
	*pc++
	return nil
}
