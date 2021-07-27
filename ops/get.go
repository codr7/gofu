package ops

import (
	"github.com/codr7/gofu"
	"github.com/codr7/gofu/errors"
)

type TGet struct {
	pos gofu.TPos
	register gofu.TRegister
}

func Get(pos gofu.TPos, reg gofu.TRegister) TGet {
	return TGet{pos: pos, register: reg}
}

func (self TGet) Eval(thread *gofu.TThread, pc *int) error {
	i := self.register.Index()
	it := thread.Get(i)

	if ct, pt := it.Type(), self.register.Type(); pt != nil && !gofu.Isa(ct, pt) {
		return errors.Eval(self.pos, "Invalid bound value: %v/%v", it, pt.Name())
	}
	
	thread.Stack().Push(it.Type(), it.Value())
	*pc++
	return nil
}
