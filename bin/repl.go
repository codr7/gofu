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
	var thread gofu.TThread
	thread.Init(&scope)

	if err := block.Run(&thread, 0); err != nil {
		fmt.Println(err)
	}
	
	fmt.Println(thread.Stack())
}
