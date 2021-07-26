package types

import (
	"github.com/codr7/gofu"
)

type TBool struct {
	gofu.BType
}

var _bool *TBool

func Bool() *TBool {
	if _bool == nil {
		_bool = new(TBool)
		_bool.Init("Bool")
	}
	
	return _bool
}
