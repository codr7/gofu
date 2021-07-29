package ops

import (
	"github.com/codr7/gofu"
	"github.com/codr7/gofu/errors"
)

type TDrop struct {
	pos gofu.TPos
	n int
}

func Drop(pos gofu.TPos, n int) TDrop {
	return TDrop{pos: pos, n: n}
}

func (self TDrop) Eval(thread *gofu.TThread, pc *int) error {
	if s := thread.Stack(); !s.Drop(self.n) {
		return errors.Eval(self.pos, "Not enough values on stack: %v", s)
	}
	
	*pc++
	return nil
}
