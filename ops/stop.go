package ops

import (
	"github.com/codr7/gofu"
	"github.com/codr7/gofu/errors"
)

type TStop struct {}

func Stop() *TStop {
	return new(TStop)
}

func (self TStop) Eval(thread *gofu.TThread, pc *int) error {
	return errors.Stop
}
