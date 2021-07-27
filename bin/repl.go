package main

import (
	"bufio"
	"fmt"
	"github.com/codr7/gofu"
	"github.com/codr7/gofu/inits"
	"github.com/codr7/gofu/ops"
	"github.com/codr7/gofu/parsers"
	"os"
)

func main() {
	fmt.Printf("gofu v%v\n", gofu.VERSION)

	block := gofu.Block()

	scope := gofu.Scope()
	inits.Core(scope)
	
	thread := gofu.Thread(scope)
	parser := gofu.Parser(parsers.Space, parsers.Id)
	in := bufio.NewReader(os.Stdin)
		
	for {
		fmt.Print("  ")
		p := gofu.Pos("repl", 1, 1)
		f, err := parser.Form(&p, in)

		if err == nil {
			if err = f.Compile(scope, block); err == nil {
				block.Emit(ops.Stop())

				if err := block.Run(thread, 0); err != nil {
					fmt.Println(err)
				}
			} else {
				fmt.Println(err)
			}
		} else {
			fmt.Println(err)
		}
	
		fmt.Println(thread.Stack())
	}
}
