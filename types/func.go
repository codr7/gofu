package types

import (
	"github.com/codr7/gofu"
	"io"
)

type TFunc struct {
	TTarget
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

func (self TFunc) TargetArgCount(val interface{}) int {
	return val.(*gofu.TFunc).ArgCount()
}

func (self TFunc) TargetApplicable(val interface{}, stack *gofu.TStack) bool {
	return val.(*gofu.TFunc).Applicable(stack)	
}

func (self TFunc) CallTarget(val interface{}, pos gofu.TPos, thread *gofu.TThread, pc *int, check bool) error {
	return val.(*gofu.TFunc).Call(pos, thread, pc, check)		
}
