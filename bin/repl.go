package main

import (
	"fmt"
	"github.com/codr7/gofu"
	"github.com/codr7/gofu/ops"
	"github.com/codr7/gofu/types"
)

func main() {
	fmt.Println("gofu v1\n")

	var block gofu.TBlock	

	var scope gofu.TScope
	scope.Init()

	var thread gofu.TThread
	thread.Init(&scope)

	scope.BindSlot("stack", types.Stack(), thread.Stack())
	
	//p := gofu.Pos("repl", 1, 1)
	block.Emit(ops.Push(types.Int(), 7))
	block.Emit(ops.Stop())

	if err := block.Run(&thread, 0); err != nil {
		fmt.Println(err)
	}
	
	fmt.Println(thread.Stack())
}
