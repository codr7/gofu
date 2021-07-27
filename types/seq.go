package types

import (
	"fmt"
	"github.com/codr7/gofu"
)

type TSeq struct {
	gofu.BType
	itemType gofu.Type
}

func (self *TSeq) Init(name string, itemType gofu.Type) *TSeq {
	self.BType.Init(name)
	self.itemType = itemType
	return self
}

func Seq(itemType gofu.Type) *TSeq {
	t := new(TSeq).Init(fmt.Sprintf("Seq<%v>", itemType.Name), itemType)
	t.AddParent(Any(), false)
	return t
}
