package gofu

import (
	"fmt"
	"io"
)

type MacroBody = func(pos TPos, args []Form, scope *TScope, block *TBlock) error

type TMacro struct {
	name string
	argCount int
	body MacroBody
}

func Macro(name string, argCount int, body MacroBody) *TMacro {
	return &TMacro{name: name, argCount: argCount, body: body}
}

func (self TMacro) Name() string {
	return self.name
}

func (self TMacro) Expand(pos TPos, args []Form, scope *TScope, block *TBlock) error {	
	return self.body(pos, args, scope, block)
}

func (self TMacro) Dump(out io.Writer) {
	fmt.Fprintf(out, "Macro(%v)", self.name)
}

func (self TMacro) ArgCount() int {
	return self.argCount
}
