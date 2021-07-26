package ops

import (
	"fmt"
	"github.com/codr7/gofu"
)

type TCall struct {
	pos gofu.TPos
	target gofu.Target
}

func Call(pos gofu.TPos, tgt gofu.Target) TCall {
	return TCall{pos: pos, target: tgt}
}

func (self TCall) Eval(thread *gofu.TThread, pc *int) error {
	thread.PushCall(self.pos, self.target)
	*pc++

	if err := self.target.Call(self.pos, thread, pc); err != nil {
		return err
	}

	if c := thread.PeekCall(); c == nil {
		return fmt.Errorf("No call in progress")
	} else if c.ReturnPc() == -1 {
		thread.PopCall()
	}

	return nil
}
