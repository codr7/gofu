package types

import (
	"fmt"
	"github.com/codr7/gofu"
	"io"
)

type TString struct {
	gofu.BType
}

var _string *TString

func String() *TString {
	if _string == nil {
		_string = new(TString)
		_string.Init("String")
		_string.AddParent(Any(), false)
		_string.AddParent(Seq(Char()), true)
	}
	
	return _string
}

func (self *TString) DumpValue(val interface{}, out io.Writer) {
	fmt.Fprintf(out, "'%v'", val)
}
