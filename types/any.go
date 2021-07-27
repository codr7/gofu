package types

import (
	"github.com/codr7/gofu"
)

type TAny struct {
	gofu.BType
}

var any *TAny

func Any() *TAny {
	if any == nil {
		any = new(TAny)
		any.Init("Any")
		any.AddParent(any, false)
	}
	
	return any
}
