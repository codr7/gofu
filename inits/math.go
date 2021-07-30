package inits

import (
	"github.com/codr7/gofu"
	"github.com/codr7/gofu/types"
)

func Math(scope *gofu.TScope) {
	scope.BindSlot("+",
		types.Func(),
		gofu.Func("+", []gofu.Type{types.Int(), types.Int()}, []gofu.Type{types.Int()},
			func(pos gofu.TPos, thread *gofu.TThread, _func *gofu.TFunc, pc *int, check bool) error {
				stack := thread.Stack()
				y := stack.Pop()
				x := stack.Peek(0)
				x.Init(types.Int(), x.Value().(int) + y.Value().(int))
				return nil
			}))

	scope.BindSlot("-",
		types.Func(),
		gofu.Func("-", []gofu.Type{types.Int(), types.Int()}, []gofu.Type{types.Int()},
			func(pos gofu.TPos, thread *gofu.TThread, _func *gofu.TFunc, pc *int, check bool) error {
				stack := thread.Stack()
				y := stack.Pop()
				x := stack.Peek(0)
				x.Init(types.Int(), x.Value().(int) - y.Value().(int))
				return nil
			}))
}
