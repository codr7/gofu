package ops

import (
	"fmt"
	"github.com/codr7/gofu"
)

type TGet struct {
	register gofu.TRegister
}

func Get(reg gofu.TRegister) TGet {
	return TGet{register: reg}
}

func (self TGet) Eval(thread *gofu.TThread, pc *int) error {
	i := self.register.Index()
	it := thread.Get(i)

	if t := self.register.Type(); t != nil && !gofu.Isa(it.Type(), t) {
		return fmt.Errorf("Invalid type for bound value: %v/%v", it, t.Name())
	}
	
	thread.Stack().Push(it.Type(), it.Value())
	*pc++
	return nil
}
