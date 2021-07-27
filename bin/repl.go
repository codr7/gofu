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

	parser := parsers.Chain(parsers.Space, parsers.Id)
	block := gofu.Block()
	thread := gofu.Thread(scope)

	utils.Repl(scope, parser, block, thread)
}
