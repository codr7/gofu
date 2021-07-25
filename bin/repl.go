package main

import (
	"fmt"
	"github.com/codr7/gofu"
	"github.com/codr7/gofu/ops"
	"github.com/codr7/gofu/types"
)

func main() {
	fmt.Println("gofu v1\n")

	var block gofu.Block	
	var scope gofu.Scope

	scope.Init()
	//p := gofu.Pos("repl", 1, 1)
	block.Emit(ops.Push(types.Int(), 7))
	block.Emit(ops.Stop())

	var calls gofu.CallStack
	var stack gofu.Stack
	stack.Init(scope.StackDepth())

	if err := block.Run(0, &calls, &stack); err != nil {
		fmt.Println(err)
	}
	
	fmt.Println(stack)
}
