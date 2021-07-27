package ops

import (
	"github.com/codr7/gofu"
	"github.com/codr7/gofu/errors"
	"github.com/codr7/gofu/types"
)

type TCall struct {
	pos gofu.TPos
	target gofu.Target
}

func Call(pos gofu.TPos, tgt gofu.Target) TCall {
	return TCall{pos: pos, target: tgt}
}

func (self TCall) Eval(thread *gofu.TThread, pc *int) error {
	stack := thread.Stack()
	t := self.target
	
	if t == nil {
		s := stack.Pop()

		if s == nil {
			return errors.Eval(self.pos, "Missing target")
		}

		if !gofu.Isa(s.Type(), types.Target()) {
			return errors.Eval(self.pos, "Invalid target: %v", s)
		}

		t = s.Value().(gofu.Target)
	}
	
	thread.PushCall(self.pos, t)
	*pc++

	if err := t.Call(self.pos, thread, pc); err != nil {
		return err
	}

	if c := thread.PeekCall(); c == nil {
		return errors.Eval(self.pos, "No call in progress")
	} else if c.ReturnPc() == -1 {
		thread.PopCall()
	}

	return nil
}
