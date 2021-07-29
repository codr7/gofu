package gofu

import (
	"fmt"
	"io"
)

type FuncBody = func(pos TPos, thread *TThread, _func *TFunc, pc *int, check bool) error

type TFunc struct {
	name string
	argTypes []Type
	resTypes []Type
	body FuncBody
}

func Func(name string, aTypes []Type, rTypes []Type, body FuncBody) *TFunc {
	return &TFunc{name: name, argTypes: aTypes, resTypes: rTypes, body: body}
}

func (self TFunc) Name() string {
	return self.name
}

func (self TFunc) ArgCount() int {
	return len(self.argTypes)
}

func (self TFunc) ResCount() int {
	return len(self.resTypes)
}

func (self TFunc) ArgTypes() []Type {
	return self.argTypes
}

func (self TFunc) ResTypes() []Type {
	return self.resTypes
}

func (self TFunc) Applicable(stack *TStack) bool {
	n := self.ArgCount()
	
	if stack.Len() <  n {
		return false
	}
	
	offs := n-1

	for i, t := range(self.argTypes) {
		if !Isa(stack.Peek(offs-i).Type(), t) {
			return false
		}
	}

	return true
}

func (self *TFunc) Call(pos TPos, thread *TThread, pc *int, check bool) error {	
	return self.body(pos, thread, self, pc, check)

}

func (self TFunc) Dump(out io.Writer) {
	fmt.Fprintf(out, "Func(%v %v %v)", self.name, self.argTypes, self.resTypes)
}
