package types

import (
	"fmt"
	"github.com/codr7/gofu"
	"io"
)

type TNil struct {
	gofu.BType
}

var _nil *TNil

func Nil() *TNil {
	if _nil == nil {
		_nil = new(TNil)
		_nil.Init("Nil")
	}
	
	return _nil
}

func (self TNil) TrueValue(val interface{}) bool {
	return false
}

func (self *TNil) DumpValue(val interface{}, out io.Writer) {
	fmt.Fprintf(out, "_")
}
