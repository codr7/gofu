package ops

import (
	"github.com/codr7/gofu"
)

type TBranch struct {
		falsePc int
}

func Branch(fpc int) TBranch {
	return TBranch{falsePc: fpc}
}

func (self TBranch) Eval(thread *gofu.TThread, pc *int) error {
	stack := thread.Stack()
	v := stack.Pop()

	if v.Type().TrueValue(v.Value()) {
		*pc++
	} else {
		*pc = self.falsePc
	}

	return  nil
}
