package ops

import (
	"github.com/codr7/gofu"
)

type TCall struct {
	pos gofu.TPos
	target gofu.Target
}

func Call(pos gofu.TPos, tgt gofu.Target) TCall {
	return TCall{pos: pos, target: tgt}
}

func (self TCall) Eval(pc *int, calls *gofu.CallStack, stack *gofu.Stack) error {
	tag := calls.Push(self.pos, self.target)
	*pc++

	if err := self.target.Call(self.pos, pc, stack); err != nil {
		return err
	}

	if err := calls.Pop(tag); err != nil {
		return err
	}
	
	return nil
}
