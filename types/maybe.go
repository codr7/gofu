package types

import (
	"fmt"
	"github.com/codr7/gofu"
	"io"
)

type TMaybe struct {
	gofu.BType
	valueType gofu.Type
}

func Maybe(valueType gofu.Type) *TMaybe {
	t := new(TMaybe)
	t.Init(fmt.Sprintf("Maybe[%v]", valueType.Name()), valueType)
	t.AddParent(Any(), false)
	valueType.AddParent(t, false)
	Nil().AddParent(t, false)
	return t
}

func (self *TMaybe) Init(name string, valueType gofu.Type) *TMaybe {
	self.BType.Init(name)
	self.valueType = valueType
	return self
}

func (self TMaybe) TrueValue(val interface{}) bool {
	if val == nil {
		return false
	}

	return self.valueType.TrueValue(val)
}

func (self TMaybe) DumpValue(val interface{}, out io.Writer) {
	if val == nil {
		fmt.Fprintf(out, "Maybe(_)")
	} else {
		fmt.Fprintf(out, "Maybe(")
		self.valueType.DumpValue(val, out)
		fmt.Fprintf(out, ")")
	}
}
