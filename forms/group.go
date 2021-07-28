package forms

import (
	"github.com/codr7/gofu"
)

type TGroup struct {
	gofu.BForm
	members []gofu.Form
}

func Group(pos gofu.TPos, members []gofu.Form) TGroup {
	var f TGroup
	f.BForm.Init(pos)
	f.members = members
	return f
}

func (self TGroup) Compile(scope *gofu.TScope, block *gofu.TBlock) error {
	for _, f := range self.members {
		if err := f.Compile(scope, block); err != nil {
			return err
		}
	}
	
	return nil
}

func (self TGroup) Slot(scope *gofu.TScope) *gofu.TSlot {
	return nil
}
