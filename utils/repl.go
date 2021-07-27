package utils

import (
	"bufio"
	"fmt"
	"github.com/codr7/gofu"
	"github.com/codr7/gofu/ops"
	"os"
)

func Repl(scope *gofu.TScope, parser *gofu.TParser, block *gofu.TBlock, thread *gofu.TThread) {
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
