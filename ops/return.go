package ops

import (
	"fmt"
	"github.com/codr7/gofu"
)

type TReturn struct {}

func Return() TReturn {
	return TReturn{}
}

func (self TReturn) Eval(thread *gofu.TThread, pc *int) error {
	c := thread.PopCall()
	
	if c == nil {
		return fmt.Errorf("No call in progress")
	}
	
	if err := c.Exit(thread, pc); err != nil {
		return err
	}

	f := c.Target().(*gofu.TFunc)
	stack := thread.Stack();
	offs := stack.Len()-1
	
	for i, t := range f.ArgTypes() {
		if v := stack.Peek(offs-i); !gofu.Isa(v.Type(), t) {
			return fmt.Errorf("Invalid result: %v/%v", v, t.Name())
		}
	}

	return nil
}
