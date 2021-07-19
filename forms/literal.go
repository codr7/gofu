package forms

import (
	"github.com/codr7/gofu"
	"github.com/codr7/gofu/ops"
)

type literal struct {
	slot gofu.Slot
}

func Literal(t gofu.Type, v interface{}) *literal {
	f := new(literal)
	f.slot.Init(t, v)
	return f
}

func (self literal) Emit(block *gofu.Block) {
	block.Emit(ops.Push(self.slot.Type, self.slot.Value))
}
