package forms

import (
	"github.com/codr7/gofu"
	"github.com/codr7/gofu/errors"
	"github.com/codr7/gofu/ops"
	"github.com/codr7/gofu/types"
)

type TCall struct {
	gofu.BForm
	target TId
	args []gofu.Form
}

func Call(pos gofu.TPos, target TId, args...gofu.Form) TCall {
	f := TCall{target: target, args: args}
	f.BForm.Init(pos)
	return f
}

func (self TCall) Compile(scope *gofu.TScope, block *gofu.TBlock) error {
	for _, a := range self.args {
		if err := a.Compile(scope, block); err != nil {
			return err
		}
	}

	f := scope.Find(self.target.name)
	
	switch s := f.(type) {
	case gofu.TRegister:
		block.Emit(ops.Get(self.Pos(), s))
		block.Emit(ops.Call(self.Pos(), nil, true))
		return nil
	case gofu.TSlot:
		if !gofu.Isa(s.Type(), types.Target()) {
			return errors.Compile(self.Pos(), "Invalid target: %v", s)
		}
		
		f = s.Value()
	default:
		return errors.Compile(self.Pos(), "Invalid target: %v", f)
	}

	switch f := f.(type) {
	case gofu.TMacro:
		return f.Expand(self.Pos(), self.args, scope, block)
	case gofu.Target:		
		if n := len(self.args) ; n != f.ArgCount() {
			return errors.Compile(self.Pos(), "Wrong number of args: %v", n)
		}

		if m, ok := f.(*gofu.TMulti); ok {
			f = m.GetFunc(self.args, scope)
		}

		unknowns := 0

		if f, ok := f.(*gofu.TFunc); ok {
			ats := f.ArgTypes()

			for i, a := range(self.args) {
				if s := a.Slot(scope); s == nil {
					unknowns++					
				} else if !gofu.Isa(s.Type(), ats[i]) {
					return errors.Compile(self.Pos(),
						"Wrong argument type: %v/%v", s.Type().Name(), ats[i].Name())
				}
			}
		}
		
		block.Emit(ops.Call(self.Pos(), f, unknowns > 0))
	default:
		return errors.Compile(self.Pos(), "Invalid target: %v", f)
	}
	
	return nil
}

func (self TCall) Slot(scope *gofu.TScope) *gofu.TSlot {
	return nil
}
