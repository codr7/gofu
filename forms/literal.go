package forms

import (
	"github.com/codr7/gofu"
	"github.com/codr7/gofu/ops"
)

type TLiteral struct {
	slot gofu.Slot
}

func Literal(t gofu.Type, v interface{}) TLiteral {
	var f TLiteral
	f.slot.Init(t, v)
	return f
}

func (self TLiteral) Compile(scope *gofu.Scope, block *gofu.Block) error {
	scope.Push()
	block.Emit(ops.Push(self.slot.Type(), self.slot.Value()))
	return nil
}
