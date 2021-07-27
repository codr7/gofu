package ops

import (
	"github.com/codr7/gofu"
	"github.com/codr7/gofu/errors"
)

type TApply struct {
	pos gofu.TPos
	_func *gofu.TFunc
}

func Apply(pos gofu.TPos, fun *gofu.TFunc) TApply {
	return TApply{pos: pos, _func: fun}
}

func (self TApply) Eval(thread *gofu.TThread, pc *int) error {
	stack := thread.Stack()
	
	if !self._func.Applicable(stack) {
		return errors.Eval(self.pos, "Function is not applicable: %v/%v", self._func.Name(), stack)
	}
	
	*pc++
	return nil
}
