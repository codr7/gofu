package ops

import (
	"fmt"
	"github.com/codr7/gofu"
)

type TBindId struct {
	pos gofu.TPos
	index int
	_type gofu.Type
	value *gofu.TSlot
}

func BindId(pos gofu.TPos, idx int, t gofu.Type,  v *gofu.TSlot) TBindId {
	return TBindId{pos: pos, index: idx, _type: t, value: v}
}

func (self TBindId) Eval(thread *gofu.TThread, pc *int) error {
	stack := thread.Stack()
	

	v := self.value

	if v == nil {
		if stack.Empty() {
			return fmt.Errorf("Missing value to bind")
		}
		
		v = stack.Pop()
	}

	thread.Set(self.index, v.Type(), v.Value())
	*pc++
	return nil
}
