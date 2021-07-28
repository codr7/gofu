package inits

import (
	"github.com/codr7/gofu"
	"github.com/codr7/gofu/types"
)

func Math(scope *gofu.TScope) {
	scope.BindSlot("+",
		types.Func(),
		gofu.Func("+", []gofu.Type{types.Int(), types.Int()}, []gofu.Type{types.Int()},
			func(pos gofu.TPos, thread *gofu.TThread, pc *int) error {
				stack := thread.Stack()
				stack.Push(types.Int(), stack.Pop().Value().(int) + stack.Pop().Value().(int))
				return nil
			}))
}
