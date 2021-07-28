package gofu

import (
	"fmt"
	"io"
)

type MacroBody = func(pos TPos, scope *TScope, block *TBlock) error

type TMacro struct {
	name string
	body MacroBody
}

func Macro(name string, body MacroBody) *TMacro {
	return &TMacro{name: name, body: body}
}

func (self TMacro) Expand(pos TPos, scope *TScope, block *TBlock) error {
	return self.body(pos, scope, block)
}

func (self TMacro) Dump(out io.Writer) {
	fmt.Fprintf(out, "Macro(%v)", self.name)
}
