package forms

import (
	"fmt"
	"github.com/codr7/gofu"
	"github.com/codr7/gofu/ops"
)

type id struct {
	name string
}

func Id(name string) *id {
	f := new(id)
	f.name = name
	return f
}

func (self id) Emit(scope *gofu.Scope, block *gofu.Block) error {
	found := scope.Find(self.name)

	if found == nil {
		return fmt.Errorf("Unknown identifier: %v", self.name)
	}

	switch found := found.(type) {
	case int:
		block.Emit(ops.Get(found))
	case gofu.Slot:
		block.Emit(ops.Push(found.Type(), found.Value()))
	default:
		return fmt.Errorf("Invalid binding: %v", found)
	}
	
	return nil
}
