package ops

import (
	"github.com/codr7/gofu"
)

type TNop struct {}

func Nop() *TNop {
	return new(TNop)
}

func (self TNop) Eval(thread *gofu.TThread, pc *int) error {
	*pc++
	return nil
}
