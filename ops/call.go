package ops

import (
	"github.com/codr7/gofu"
	"github.com/codr7/gofu/errors"
	"github.com/codr7/gofu/types"
)

type TCall struct {
	pos gofu.TPos
	target *gofu.TSlot
	check bool
}

func Call(pos gofu.TPos, t gofu.Type, v interface{}, chk bool) TCall {
	op := TCall{pos: pos, check: chk}

	if t != nil || v != nil {
		op.target = gofu.Slot(t, v)
	}

	return op
}

func (self TCall) Eval(thread *gofu.TThread, pc *int) error {
	stack := thread.Stack()
	t := self.target
	
	if t == nil {
		t = stack.Pop()

		if t == nil {
			return errors.Eval(self.pos, "Missing target")
		}
	}

	if !gofu.Isa(t.Type(), types.Target()) {
		return errors.Eval(self.pos, "Invalid target: %v", t)
	}

	*pc++

	if err := t.Type().(types.ITarget).CallTarget(t.Value(), self.pos, thread, pc, self.check); err != nil {
		return err
	}

	return nil
}
