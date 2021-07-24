package types

import (
	"github.com/codr7/gofu"
)

type TFuncSet struct {
	gofu.BasicType
}

var funcSet *TFuncSet

func FuncSet() *TFuncSet {
	if funcSet == nil {
		funcSet = new(TFuncSet)
		funcSet.Init("FuncSet")
		funcSet.AddParent(Target())
	}
	
	return funcSet
}
