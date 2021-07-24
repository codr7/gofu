package types

import (
	"github.com/codr7/gofu"
)

func Isa(child, parent gofu.Type) bool {
	if child == parent {
		return true
	}

	return child.Isa(parent)
}
