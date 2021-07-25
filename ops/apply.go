package ops

import (
	"fmt"
	"github.com/codr7/gofu"
)

type TApply struct {
	pos gofu.TPos
	_func *gofu.TFunc
}

func Apply(pos gofu.TPos, fun *gofu.TFunc) TApply {
	return TApply{pos: pos, _func: fun}
}

func (self TApply) Eval(pc *int, calls *gofu.CallStack, stack *gofu.Stack) error {
	if !self._func.Applicable(stack) {
		return fmt.Errorf("Function is not applicable: %v/%v", self._func.Name(), stack)
	}
	
	*pc++
	return nil
}
