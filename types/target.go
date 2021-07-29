package types

import (
	"github.com/codr7/gofu"
)

type ITarget interface {
	TargetArgCount(val interface{}) int
	TargetApplicable(val interface{}, stack *gofu.TStack) bool
	CallTarget(val interface{}, pos gofu.TPos, thread *gofu.TThread, pc *int, check bool) error
}

type TTarget struct {
	gofu.BType
}

var target *TTarget

func Target() *TTarget {
	if target == nil {
		target = new(TTarget)
		target.Init("Target")
		target.AddParent(Any(), false)
	}
	
	return target
}
