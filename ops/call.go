package ops

import (
	"github.com/codr7/gofu"
)

type TCall struct {
	target gofu.Target
}

func Call(target gofu.Target) *TCall {
	op := new(TCall)
	op.target = target
	return op
}

func (self TCall) Eval(pc int, stack *gofu.Stack) (int, error) {
	return pc+1, self.target.Call(stack)
}
