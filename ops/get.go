package ops

import (
	"github.com/codr7/gofu"
)

type TGet struct {
	index int
}

func Get(idx int) TGet {
	return TGet{index: idx}
}

func (self TGet) Eval(thread *gofu.TThread, pc *int) error {
	it := thread.Get(self.index)
	thread.Stack().Push(it.Type(), it.Value())
	*pc++
	return nil
}
