package forms

import (
	"github.com/codr7/gofu"
	"github.com/codr7/gofu/errors"
	"github.com/codr7/gofu/ops"
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

func (self TId) Compile(scope *gofu.TScope, block *gofu.TBlock) error {
	found := scope.Find(self.name)

	if found == nil {
		return errors.Compile(self.Pos(), "Unknown identifier: %v", self.name)
	}

	switch found := found.(type) {
	case gofu.TRegister:
		block.Emit(ops.Get(self.Pos(), found))
	case gofu.TSlot:
		switch v := found.Value().(type) {
		case *gofu.TMacro:
			return v.Call(self.Pos(), scope, block)
		default:
			block.Emit(ops.Push(found.Type(), found.Value()))
		}
	default:
		return errors.Compile(self.Pos(), "Invalid binding: %v", found)
	}
	
	return nil
}

func (self TId) Slot(scope *gofu.TScope) *gofu.TSlot {
	var s gofu.TSlot
	
	switch found := scope.Find(self.name).(type) {
	case gofu.TRegister:
		s = gofu.Slot(found.Type(), nil)
	case gofu.TSlot:
		s = found
	default:
		return nil
	}

	return &s
}
