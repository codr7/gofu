package forms

import (
	"fmt"
	"github.com/codr7/gofu"
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
		switch a := a.(type) {
		case TLiteral:
			block.Emit(ops.Push(a.slot.Type(), a.slot.Value()))
		default:
			return fmt.Errorf("Invalid argument: %v", a)
		}
	}

	f := scope.Find(self.target.name)
	
	switch s := f.(type) {
	case int:
		return fmt.Errorf("Dynamic calls are not implemented: %v", f)
	case gofu.TSlot:
		if !gofu.Isa(s.Type(), types.Target()) {
			return fmt.Errorf("Invalid target: %v", s)
		}
		
		f = s.Value()
	default:
		return fmt.Errorf("Invalid target: %v", f)
	}

	switch f := f.(type) {
	case gofu.Target:
		if n := len(self.args) ; n != f.ArgCount() {
			return fmt.Errorf("Wrong number of args: %v", n)
		}

		if m, ok := f.(*gofu.TMulti); ok {
			f = m.GetFunc(self.args, scope)
		}

		if f, ok := f.(*gofu.TFunc); ok {
			ats := f.ArgTypes()
			unknowns := 0

			for i, a := range(self.args) {
				if s := a.Slot(scope); s == nil {
					unknowns++					
				} else if !gofu.Isa(s.Type(), ats[i]) {
					return fmt.Errorf("Wrong argument type: %v/%v", s.Type().Name(), ats[i].Name())
				}
			}

			if unknowns > 0 {
				block.Emit(ops.Apply(self.Pos(), f))
			}
		}
		
		block.Emit(ops.Call(self.Pos(), f))
	default:
		return fmt.Errorf("Invalid target: %v", f)
	}
	
	return nil
}

func (self TCall) Slot(scope *gofu.TScope) *gofu.TSlot {
	return nil
}
