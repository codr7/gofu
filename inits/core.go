package inits

import (
	"github.com/codr7/gofu"
	"github.com/codr7/gofu/ops"
	"github.com/codr7/gofu/types"
)

func Core(scope *gofu.TScope) {
	scope.BindSlot("Any", types.Meta(), types.Any())
	scope.BindSlot("Bool", types.Meta(), types.Bool())
	scope.BindSlot("Char", types.Meta(), types.Char())
	scope.BindSlot("Func", types.Meta(), types.Func())
	scope.BindSlot("Int", types.Meta(), types.Int())
	scope.BindSlot("Meta", types.Meta(), types.Meta())
	scope.BindSlot("Multi", types.Meta(), types.Multi())
	scope.BindSlot("Num", types.Meta(), types.Num())
	scope.BindSlot("Seq", types.Meta(), types.Seq(types.Any()))
	scope.BindSlot("Stack", types.Meta(), types.Stack(types.Any()))
	scope.BindSlot("String", types.Meta(), types.String())
	scope.BindSlot("Target", types.Meta(), types.Target())

	scope.BindSlot("_", types.Nil(), nil)
	scope.BindSlot("true", types.Bool(), true)
	scope.BindSlot("false", types.Bool(), false)

	scope.BindSlot("reset",
		types.Macro(),
		gofu.Macro("reset", 0,
			func(pos gofu.TPos, args []gofu.Form, scope *gofu.TScope, block *gofu.TBlock) error {
				block.Emit(ops.Reset())
				return nil
			}))
	
	scope.BindSlot("stack",
		types.Func(),
		gofu.Func("stack", nil, []gofu.Type{types.Stack(types.Any())},
			func(pos gofu.TPos, thread *gofu.TThread, pc *int) error {
				stack := thread.Stack()
				stack.Push(types.Stack(types.Any()), stack)
				return nil
			}))
}
