package types

import (
	"github.com/codr7/gofu"
	"github.com/codr7/gofu/errors"
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
	stack := thread.Stack()
	f := val.(*gofu.TFunc)
	
	if check && !f.Applicable(stack) {
		return errors.Eval(pos, "Func is not applicable: %v/%v", f, stack)
	}

	return f.Call(pos, thread, pc)		
}
