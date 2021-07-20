package forms

import (
	"github.com/codr7/gofu"
	"github.com/codr7/gofu/ops"
)

type TBindId struct {
	id string
}

func BindId(id string) *TBindId {
	f := new(TBindId)
	f.id = id
	return f
}

func (self TBindId) Emit(scope *gofu.Scope, block *gofu.Block) error {
	i := scope.BindId(self.id)
	block.Emit(ops.BindId(i))
	return nil
}
