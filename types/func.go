package types

import (
	"fmt"
	"github.com/codr7/gofu"
)

type TFunc struct {
	gofu.BasicType
}

func (self TFunc) CallValue(val interface{}, stack *gofu.Stack) error {
	f := val.(gofu.TFunc)
	fmt.Printf("Calling %v!\n", f.Name())
	return f.Call(stack)
}

var Func TFunc

func init() {
	Func.Init("Func")
}
