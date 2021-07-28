package main

import (
	"fmt"
	"github.com/codr7/gofu"
	"github.com/codr7/gofu/inits"
	"github.com/codr7/gofu/parsers"
	"github.com/codr7/gofu/utils"
)

func main() {
	fmt.Printf("gofu v%v\n", gofu.VERSION)

	scope := gofu.Scope()
	inits.Core(scope)
	inits.Math(scope)
	
	block := gofu.Block()
	thread := gofu.Thread(scope)

	utils.Repl(scope, parsers.Any(), block, thread)
}
