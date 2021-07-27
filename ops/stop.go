package ops

import (
	"github.com/codr7/gofu"
)

type TStop struct {}

var stop TStop

func Stop() TStop {
	return stop
}

func (self TStop) Eval(thread *gofu.TThread, pc *int) error {
	return gofu.Stop()
}
