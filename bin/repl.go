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
	p := gofu.Pos("test", -1, -1)
	forms.Literal(p, &types.Int, 7).Compile(&scope, &block)

	scope.BindSlot("foo", &types.Int, 14)	
	forms.Id(p, "foo").Compile(&scope, &block)

	block.Emit(ops.Push(&types.Int, 21))
	forms.BindId(p, "bar").Compile(&scope, &block)
	forms.Id(p, "bar").Compile(&scope, &block)

	f := gofu.NewFunc("baz", []gofu.Type{&types.Int}, &types.Int, func(stack *gofu.Stack) error {
		fmt.Printf("Inside baz!\n")
		return nil
	})

	scope.BindSlot("baz", types.Func, f)
	c := forms.Call(p, forms.Id(p, "baz"), forms.Literal(p, &types.Int, 28))

	if err := c.Compile(&scope, &block); err != nil {
		fmt.Println(err)
	}
	
	block.Emit(ops.Stop())

	var calls gofu.CallStack
	var stack gofu.Stack
	stack.Init(scope.StackDepth())

	if err := block.Run(0, &calls, &stack); err != nil {
		fmt.Println(err)
	}
	
	fmt.Println(stack)
}
