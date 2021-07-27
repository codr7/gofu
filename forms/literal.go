package forms

import (
	"github.com/codr7/gofu"
	"github.com/codr7/gofu/ops"
)

type TLiteral struct {
	gofu.BForm
	slot gofu.TSlot
}

func Literal(pos gofu.TPos, t gofu.Type, v interface{}) TLiteral {
	var f TLiteral
	f.BForm.Init(pos)
	f.slot.Init(t, v)
	return f
}

func (self TLiteral) Compile(scope *gofu.TScope, block *gofu.TBlock) error {
	block.Emit(ops.Push(self.slot.Type(), self.slot.Value()))
	return nil
}

func (self TLiteral) Slot() *gofu.TSlot {
	return &self.slot
}
