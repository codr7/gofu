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
	
	c.Exit(thread, pc)
	return  nil
}
