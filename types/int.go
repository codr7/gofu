package types

import (
	"github.com/codr7/gofu"
)

type TInt struct {
	gofu.BasicType
}

var _int *TInt

func Int() *TInt {
	if _int == nil {
		_int = new(TInt)
		_int.Init("Int")
		_int.AddParent(Num())
	}
	
	return _int
}
