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
	arguments []gofu.Form
}

func Call(pos gofu.TPos, target TId, args...gofu.Form) TCall {
	f := TCall{target: target, arguments: args}
	f.BForm.Init(pos)
	return f
}

func (self TCall) Compile(scope *gofu.Scope, block *gofu.Block) error {
	for _, a := range self.arguments {
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
	case gofu.Slot:
		if gofu.Isa(s.Type(), types.Target()) {
			return fmt.Errorf("Invalid call target: %v", s)
		}
		
		f = s.Value()
	default:
		return fmt.Errorf("Invalid target: %v", f)
	}

	switch f := f.(type) {
	case gofu.Target:
		if n := len(self.arguments) ; n != f.Arity() {
			return fmt.Errorf("Wrong number of arguments: %v", n)
		}

		if m, ok := f.(*gofu.TMulti); ok {
			f = m.GetFunc(self.arguments)
		}

		if f, ok := f.(*gofu.TFunc); ok {
			ats := f.ArgumentTypes()
			unknowns := 0

			for i, a := range(self.arguments) {
				switch a := a.(type) {
				case TLiteral:
					if x, y := a.slot.Type(), ats[i]; !gofu.Isa(x, y) {
						return fmt.Errorf("Wrong argument type: %v/%v", x, y)
					}
				default:
					unknowns++
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
