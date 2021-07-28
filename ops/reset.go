package ops

import (
	"github.com/codr7/gofu"
)

type TReset struct {}

var reset TReset

func Reset() TReset {
	return reset
}

func (self TReset) Eval(thread *gofu.TThread, pc *int) error {
	thread.Stack().Init(nil)
	*pc++
	return nil
}
