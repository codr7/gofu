package ops

import (
	"github.com/codr7/gofu"
	"github.com/codr7/gofu/errors"
)

type TBindId struct {
	pos gofu.TPos
	index int
	_type gofu.Type
}

func BindId(pos gofu.TPos, idx int, typ gofu.Type) TBindId {
	return TBindId{pos: pos, index: idx, _type: typ}
}

func (self TBindId) Eval(thread *gofu.TThread, pc *int) error {
	stack := thread.Stack()
			
	s := stack.Pop()

	if s == nil {
		return errors.Eval(self.pos, "Missing value")
	}

	if ct, pt := s.Type(), self._type; pt != nil && !gofu.Isa(ct, pt) {
		return errors.Eval(self.pos, "Invalid binding: %v/%v", s, pt.Name())
	}

	thread.Set(self.index, s.Type(), s.Value())
	*pc++
	return nil
}
