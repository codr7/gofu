package types

import (
	"github.com/codr7/gofu"
	"io"
)

type TMacro struct {
	gofu.BType
}

var macro *TMacro

func Macro() *TMacro {
	if (macro == nil) {
		macro = new(TMacro)
		macro.Init("Macro")
		macro.AddParent(Any(), false)
	}

	return macro
}

func (self TMacro) DumpValue(val interface{}, out io.Writer) {
	val.(*gofu.TMacro).Dump(out)
}

