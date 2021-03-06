package forms

import (
	"github.com/codr7/gofu"
	"github.com/codr7/gofu/errors"
	"github.com/codr7/gofu/ops"
	"github.com/codr7/gofu/types"
)

type TId struct {
	gofu.BForm
	name string
}

func Id(pos gofu.TPos, name string) TId {
	f := TId{name: name}
	f.BForm.Init(pos)
	return f
}

func (self TId) Name() string {
	return self.name
}

func (self TId) Compile(scope *gofu.TScope, block *gofu.TBlock) error {
	found := scope.Find(self.name)

	if found == nil {
		return errors.Compile(self.Pos(), "Unknown identifier: %v", self.name)
	}

	switch found := found.(type) {
	case gofu.TRegister:
		block.Emit(ops.Get(self.Pos(), found))
	case gofu.TSlot:
		t := found.Type()
		
		if t == types.Macro() {
			return found.Value().(*gofu.TMacro).Expand(self.Pos(), nil, scope, block)
		}

		if gofu.Isa(t, types.Target()) {
			block.Emit(ops.Call(self.Pos(), found.Type(), found.Value(), true))
		} else {
			block.Emit(ops.Push(found.Type(), found.Value()))
		}
	default:
		return errors.Compile(self.Pos(), "Invalid binding: %v", found)
	}
	
	return nil
}

func (self TId) Slot(scope *gofu.TScope) *gofu.TSlot {
	var s *gofu.TSlot
	
	switch found := scope.Find(self.name).(type) {
	case gofu.TRegister:
		s = gofu.Slot(found.Type(), nil)
	case gofu.TSlot:
		s = &found
	}

	return s
}
