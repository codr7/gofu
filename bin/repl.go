package main

import (
	"fmt"
	"github.com/codr7/gofu"
	"github.com/codr7/gofu/forms"
	"github.com/codr7/gofu/ops"
	"github.com/codr7/gofu/types"
)

func main() {
	fmt.Println("gofu v1\n")

	var block gofu.Block	
	var scope gofu.Scope

	scope.Init()
	forms.Literal(types.Int, 7).Emit(&scope, &block)

	scope.BindSlot("foo", types.Int, 14)	
	forms.Id("foo").Emit(&scope, &block)

	block.Emit(ops.Push(types.Int, 21))
	forms.BindId("bar").Emit(&scope, &block)
	forms.Id("bar").Emit(&scope, &block)
	
	block.Emit(ops.Stop())
	
	var stack gofu.Stack
	stack.Init(scope.StackDepth())
	block.Eval(0, &stack)
	
	fmt.Printf("%v %v %v\n", stack.Pop(), stack.Pop(), stack.Pop())
}
