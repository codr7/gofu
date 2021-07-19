package main

import (
	"fmt"
	"github.com/codr7/gofu"
	"github.com/codr7/gofu/forms"
	"github.com/codr7/gofu/ops"
)

func main() {
	fmt.Println("gofu v1\n")

	var intType gofu.BasicType
	intType.Init("Int")
	
	var block gofu.Block
	forms.Literal(intType, 42).Emit(&block)
	block.Emit(ops.Stop())
	
	var stack gofu.Stack
	block.Eval(0, &stack)
	fmt.Printf("Value: %v\n", stack.Peek(0))
}
