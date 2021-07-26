package ops

import (
	"github.com/codr7/gofu"
)

type TNop struct {}

var nop TNop

func Nop() TNop {
	return nop
}

func (self TNop) Eval(thread *gofu.TThread, pc *int) error {
	*pc++
	return nil
}
