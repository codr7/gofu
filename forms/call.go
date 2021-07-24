package forms

import (
	"fmt"
	"github.com/codr7/gofu"
	"github.com/codr7/gofu/ops"
	"github.com/codr7/gofu/types"
)

type TCall struct {
	_func TId
	arguments []gofu.Form
}

func Call(_func TId, args...gofu.Form) TCall {
	return TCall{_func: _func, arguments: args}
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

	f := scope.Find(self._func.name)
	
	switch s := f.(type) {
	case int:
		return fmt.Errorf("Dynamic calls not supported yet: %v", f)
	case gofu.Slot:
		f = s.Value()
	default:
		return fmt.Errorf("Invalid target: %v", f)
	}

	switch f := f.(type) {
	case gofu.Target:
		ats := f.ArgumentTypes()
		arity := len(ats)
		
		if n := len(self.arguments) ; n != arity {
			return fmt.Errorf("Wrong number of arguments: %v", n)
		}

		for i, a := range(self.arguments) {
			switch a := a.(type) {
			case TLiteral:
				if x, y := a.slot.Type(), ats[i]; !types.Isa(x, y) {
					return fmt.Errorf("Wrong argument type: %v/%v", x, y)
				}
			}
		}
		
		block.Emit(ops.Call(f))
	default:
		return fmt.Errorf("Invalid target: %v", f)
	}
	
	return nil
}