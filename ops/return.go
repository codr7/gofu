package ops

import (
	"github.com/codr7/gofu"
	"github.com/codr7/gofu/errors"
)

type TReturn struct {
	pos gofu.TPos
}

func Return(pos gofu.TPos) TReturn {
	return TReturn{pos: pos}
}

func (self TReturn) Eval(thread *gofu.TThread, pc *int) error {
	c := thread.PopCall()
	
	if c == nil {
		return errors.Eval(self.pos, "No call in progress")
	}
	
	if err := c.Exit(thread, pc); err != nil {
		return err
	}

	f := c.Target().(*gofu.TFunc)
	stack := thread.Stack();
	offs := stack.Len()-1

	if x, y := stack.Len(), f.ResCount(); x < y {
		return errors.Eval(self.pos, "Invalid result: %v/%v/%v", f.Name(), x, y)
	}
	
	for i, t := range f.ArgTypes() {
		if v := stack.Peek(offs-i); !gofu.Isa(v.Type(), t) {
			return errors.Eval(self.pos, "Invalid result: %v/%v/%v", f.Name(), v, t.Name())
		}
	}

	return nil
}
