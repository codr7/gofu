package types

import (
	"github.com/codr7/gofu"
	"io"
)

type TMeta struct {
	gofu.BType
}

var meta *TMeta

func Meta() *TMeta {
	if meta == nil {
		meta = new(TMeta)
		meta.Init("Meta")
		meta.AddParent(Any(), false)
	}
	
	return meta
}

func (self *TMeta) DumpValue(val interface{}, out io.Writer) {
	io.WriteString(out, val.(gofu.Type).Name())
}
