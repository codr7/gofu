package types

import (
	"github.com/codr7/gofu"
)

type TInt struct {
	gofu.BType
}

var _int *TInt

func Int() *TInt {
	if _int == nil {
		_int = new(TInt)
		_int.Init("Int")
		_int.AddParent(Any(), false)
		_int.AddParent(Num(), true)
	}
	
	return _int
}

func (self TInt) TrueValue(val interface{}) bool {
	return val.(int) != 0
}
