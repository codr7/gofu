package forms

import (
	"fmt"
	"github.com/codr7/gofu"
	"github.com/codr7/gofu/ops"
)

type TBindId struct {
	gofu.BForm
	id string
}

func BindId(pos gofu.TPos, id string) TBindId {
	f := TBindId{id: id}
	f.BForm.Init(pos)
	return f
}

func (self TBindId) Compile(scope *gofu.Scope, block *gofu.Block) error {
	i := scope.BindId(self.id)

	if i == -1 {
		return fmt.Errorf("Duplicate binding: %v", self.id) 
	}
	
	block.Emit(ops.BindId(self.Pos(), i))
	return nil
}

