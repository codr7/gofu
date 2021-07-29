package ops

import (
	"github.com/codr7/gofu"
	"github.com/codr7/gofu/errors"
)

type TBindId struct {
	pos gofu.TPos
	index int
	slot gofu.TSlot
}

func BindId(pos gofu.TPos, idx int, t gofu.Type,  v interface{}) TBindId {
	return TBindId{pos: pos, index: idx, slot: *gofu.Slot(t, v)}
}

func (self TBindId) Eval(thread *gofu.TThread, pc *int) error {
	stack := thread.Stack()
	s := self.slot
	
	if v := self.slot.Value(); v == nil {
		if stack.Empty() {
			return errors.Eval(self.pos, "Missing value to bind")
		}
		
		s = *stack.Pop()
		
		if ct, pt := s.Type(), self.slot.Type(); pt != nil && !gofu.Isa(ct, pt) {
			return errors.Eval(self.pos, "Invalid binding: %v/%v", s, pt.Name())
		}
	}

	thread.Set(self.index, s.Type(), s.Value())
	*pc++
	return nil
}
