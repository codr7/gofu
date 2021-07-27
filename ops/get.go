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

	if ct, pt := it.Type(), self.register.Type(); pt != nil && !gofu.Isa(ct, pt) {
		return fmt.Errorf("Invalid bound value: %v/%v", it, pt.Name())
	}
	
	thread.Stack().Push(it.Type(), it.Value())
	*pc++
	return nil
}
