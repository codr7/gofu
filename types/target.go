package types

import (
	"github.com/codr7/gofu"
)

type Target interface {
	gofu.Type
	CallValue(val interface{}, stack *gofu.Stack) *gofu.Slot
}
