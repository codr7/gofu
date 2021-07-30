package forms

import (
	"github.com/codr7/gofu"
	"github.com/codr7/gofu/errors"
	"github.com/codr7/gofu/ops"
)

type TBindId struct {
	gofu.BForm
	id string
	_type gofu.Type
	value gofu.Form
}

func BindId(pos gofu.TPos, id string, t gofu.Type, v gofu.Form) TBindId {
	f := TBindId{id: id, _type: t, value: v}
	f.BForm.Init(pos)
	return f
}

func (self TBindId) Compile(scope *gofu.TScope, block *gofu.TBlock) error {
	if v, ok := self.value.(TLiteral); ok {
		s := v.Slot(scope)
		
		if x, y := s.Type(), self._type;  y != nil && !gofu.Isa(x, y) {
			return errors.Compile(self.Pos(), "Incompatible type: %v/%v", x, y)
		}

		scope.BindSlot(self.id, s.Type(), s.Value())
		return nil
	}
	
	i := scope.BindId(self.id, self._type)

	if i == -1 {
		return errors.Compile(self.Pos(), "Dup binding: %v", self.id)
	}

	if err := self.value.Compile(scope, block); err != nil {
		return err
	}
	
	block.Emit(ops.BindId(self.Pos(), i, self._type))
	return nil
}

func (self TBindId) Slot(scope *gofu.TScope) *gofu.TSlot {
	return nil
}


