package forms

import (
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
	block.Emit(ops.BindId(i))
	return nil
}
