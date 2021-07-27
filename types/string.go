package types

import (
	"fmt"
	"github.com/codr7/gofu"
	"io"
)

type TString struct {
	gofu.BType
}

var string *TString

func String() *TString {
	if string == nil {
		string = new(TString)
		string.Init("String")
		stack.AddParent(Seq(), true)
	}
	
	return string
}

func (self *TString) DumpValue(val interface{}, out io.Writer) {
	fmt.Fprintf(out, "'%v'", val)
}
