package ops

import (
	"fmt"
	"github.com/codr7/gofu"
)

type TCheck struct {
	pos gofu.TPos
	_func *gofu.TFunc
}

func Check(pos gofu.TPos, fun *gofu.TFunc) TCheck {
	return TCheck{pos: pos, _func: fun}
}

func (self TCheck) Eval(pc *int, calls *gofu.CallStack, stack *gofu.Stack) error {
	if !self._func.Check(stack) {
		return fmt.Errorf("Function not applicable: %v/%v", self._func.Name(), stack)
	}
	
	*pc++
	return nil
}
