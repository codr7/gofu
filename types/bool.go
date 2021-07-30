package types

import (
	"fmt"
	"github.com/codr7/gofu"
	"io"
)

type TBool struct {
	gofu.BType
}

var _bool *TBool

func Bool() *TBool {
	if _bool == nil {
		_bool = new(TBool)
		_bool.Init("Bool")
		_bool.AddParent(Any(), false)
	}
	
	return _bool
}

func (self TBool) TrueValue(val interface{}) bool {
	return val.(bool)
}

func (self TBool) DumpValue(val interface{}, out io.Writer) {
	if val.(bool) {
		fmt.Fprint(out, "t")
	} else {
		fmt.Fprint(out, "f")
	}	
}
