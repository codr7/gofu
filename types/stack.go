package types

import (
	"github.com/codr7/gofu"
	"io"
)

type TStack struct {
	gofu.BType
}

var stack *TStack

func Stack() *TStack {
	if stack == nil {
		stack = new(TStack)
		stack.Init("Stack")
		stack.AddParent(Seq(), true)
	}
	
	return stack
}

func (self *TStack) DumpValue(val interface{}, out io.Writer) {
	io.WriteString(out, val.(*gofu.TStack).String())
}
