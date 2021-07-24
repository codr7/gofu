package types

import (
	"github.com/codr7/gofu"
)

type TFunc struct {
	gofu.BType
}

var _func *TFunc

func Func() *TFunc {
	if (_func == nil) {
		_func = new(TFunc)
		_func.Init("Func")
		_func.AddParent(Target())
	}

	return _func
}
