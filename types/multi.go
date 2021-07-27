package types

import (
	"github.com/codr7/gofu"
)

type TMulti struct {
	gofu.BType
}

var multi *TMulti

func Multi() *TMulti {
	if multi == nil {
		multi = new(TMulti)
		multi.Init("Multi")
		multi.AddParent(Any(), false)
		multi.AddParent(Target(), true)
	}
	
	return multi
}
