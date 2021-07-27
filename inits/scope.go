package inits

import (
	"github.com/codr7/gofu"
	"github.com/codr7/gofu/types"
)

func Scope(scope *gofu.TScope) {
	scope.BindSlot("true", types.Bool(), true)
	scope.BindSlot("false", types.Bool(), false)
	
	scope.BindSlot("stack",
		types.Func(),
		gofu.Func("stack", nil, []gofu.Type{types.Int()},
			func(pos gofu.TPos, thread *gofu.TThread, pc *int) error {
				stack := thread.Stack()
				stack.Push(types.Stack(types.Any()), stack)
				return nil
			}))
}
