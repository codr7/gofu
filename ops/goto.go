package ops

import (
	"github.com/codr7/gofu"
)

type TGoto struct {
	pc int
}

func Goto(pc int) TGoto {
	return TGoto{pc: pc}
}

func (self TGoto) Eval(thread *gofu.TThread, pc *int) error {
	*pc = self.pc
	return nil
}
