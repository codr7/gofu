package types

import (
	"github.com/codr7/gofu"
)

type TMeta struct {
	gofu.BType
}

var meta *TMeta

func Meta() *TMeta {
	if meta == nil {
		meta = new(TMeta)
		meta.Init("Meta")
		meta.AddParent(Target())
	}
	
	return meta
}
