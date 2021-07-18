package ops

import (
	"github.com/codr7/gofu"
)

type stop struct {}

func Stop() *stop {
	return new(stop)
}

func (self stop) Eval(pc int, stack *gofu.Stack) (int, error) {
	return -1, nil
}
