package utils

import (
	"bufio"
	"fmt"
	"github.com/codr7/gofu"
	"github.com/codr7/gofu/ops"
	"os"
)

func Repl(scope *gofu.TScope, parser gofu.Parser, block *gofu.TBlock, thread *gofu.TThread) {
	in := bufio.NewReader(os.Stdin)
		
	for {
		fmt.Print("  ")
		p := gofu.Pos("repl", 1, 1)
		f, err := parser(&p, in)

		if err != nil {
			fmt.Println(err)
		} else if f != nil {
			startPc := block.Pc()
			
			if err = f.Compile(scope, block); err == nil {
				block.Emit(ops.Stop())
				
				if err := block.Run(thread, startPc); err != nil {
					fmt.Println(err)
				}
			} else {
				fmt.Println(err)
			}
		}
	
		fmt.Println(thread.Stack())
	}
}
