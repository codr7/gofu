package types

import (
	"github.com/codr7/gofu"
	"io"
)

type TFunc struct {
	gofu.BType
}

var _func *TFunc

func Func() *TFunc {
	if (_func == nil) {
		_func = new(TFunc)
		_func.Init("Func")
		_func.AddParent(Any(), false)
		_func.AddParent(Target(), true)
	}

	return _func
}

func (self TFunc) DumpValue(val interface{}, out io.Writer) {
	val.(*gofu.TFunc).Dump(out)
}

