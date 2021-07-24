package types

import (
	"github.com/codr7/gofu"
)

type TTarget struct {
	gofu.BasicType
}

var target *TTarget

func Target() *TTarget {
	if target == nil {
		target = new(TTarget)
		target.Init("Target")
	}
	
	return target
}
