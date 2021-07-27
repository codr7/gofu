package forms

import (
	"fmt"
	"github.com/codr7/gofu"
	"github.com/codr7/gofu/ops"
)

type TBindId struct {
	gofu.BForm
	id string
	_type gofu.Type
}

func BindId(pos gofu.TPos, id string, t gofu.Type) TBindId {
	f := TBindId{id: id, _type: t}
	f.BForm.Init(pos)
	return f
}

func (self TBindId) Compile(scope *gofu.TScope, block *gofu.TBlock) error {
	i := scope.BindId(self.id, self._type)

	if i == -1 {
		return fmt.Errorf("Duplicate binding: %v", self.id) 
	}
	
	block.Emit(ops.BindId(self.Pos(), i))
	return nil
}

func (self TBindId) Slot() *gofu.TSlot {
	return nil
}


