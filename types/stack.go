package types

import (
	"fmt"
	"github.com/codr7/gofu"
	"io"
)

type TStack struct {
	gofu.BType
	itemType gofu.Type
}

func Stack(itemType gofu.Type) *TStack {
	t := new(TStack)
	t.Init(fmt.Sprintf("Stack[%v]", itemType.Name()), itemType)
	t.AddParent(Any(), false)
	t.AddParent(Seq(itemType), true)
	return t
}

func (self *TStack) Init(name string, itemType gofu.Type) *TStack {
	self.BType.Init(name)
	self.itemType = itemType
	return self
}

func (self TStack) DumpValue(val interface{}, out io.Writer) {
	val.(*gofu.TStack).Dump(out)
}
