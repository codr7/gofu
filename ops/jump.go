package ops

import (
	"github.com/codr7/gofu"
)

type TJump struct {
	pc int
}

func Jump(pc int) TJump {
	return TJump{pc: pc}
}

func (self TJump) Eval(thread *gofu.TThread, pc *int) error {
	*pc = self.pc
	return nil
}
