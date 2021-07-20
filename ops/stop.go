package ops

import (
	"github.com/codr7/gofu"
)

type TStop struct {}

func Stop() *TStop {
	return new(TStop)
}

func (self TStop) Eval(pc int, stack *gofu.Stack) (int, error) {
	return -1, nil
}
