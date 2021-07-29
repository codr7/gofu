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

func (self TCall) CompileArgs(scope *gofu.TScope, block *gofu.TBlock) error {
	for _, a := range self.args {
		if err := a.Compile(scope, block); err != nil {
			return err
		}
	}

	return nil
}

func (self TCall) Compile(scope *gofu.TScope, block *gofu.TBlock) error {
	f := scope.Find(self.target.name)
	
	switch s := f.(type) {
	case gofu.TRegister:
		if err := self.CompileArgs(scope, block); err != nil {
			return err
		}
		
		block.Emit(ops.Get(self.Pos(), s))
		block.Emit(ops.Call(self.Pos(), nil, nil, true))
		return nil
	case gofu.TSlot:
		if !gofu.Isa(s.Type(), types.Target()) {
			return errors.Compile(self.Pos(), "Invalid target: %v", s)
		}
	default:
		return errors.Compile(self.Pos(), "Invalid target: %v", f)
	}

	ts := f.(gofu.TSlot)
	tt := ts.Type()
	tv := ts.Value()
	
	if (gofu.Isa(tt, types.Macro())) {
		m := tv.(*gofu.TMacro)
		
		if len(self.args) != m.ArgCount() {
			return errors.Compile(self.Pos(), "Wrong number of macro arguments: %v/%v", m.Name(), self.args)
		}
		
		return m.Expand(self.Pos(), self.args, scope, block)
	} else if (gofu.Isa(tt, types.Target())) {
		if n := len(self.args) ; n != tt.(types.ITarget).TargetArgCount(tv) {
			return errors.Compile(self.Pos(), "Wrong number of args: %v", n)
		}

		if m, ok := tv.(*gofu.TMulti); ok {
			tv = m.GetFunc(self.args, scope)
		}

		unknowns := 0

		if f, ok := tv.(*gofu.TFunc); ok {
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

		if err := self.CompileArgs(scope, block); err != nil {
			return err
		}
		
		block.Emit(ops.Call(self.Pos(), tt, tv, unknowns > 0))
	} else {
		return errors.Compile(self.Pos(), "Invalid target: %v", ts)
	}
	
	return nil
}

func (self TCall) Slot(scope *gofu.TScope) *gofu.TSlot {
	return nil
}
