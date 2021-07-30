package gofu

import (
	"fmt"
	"io"
	"strings"
)

type FuncBody = func(pos TPos, thread *TThread, _func *TFunc, pc *int) error

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

func (self *TFunc) Call(pos TPos, thread *TThread, pc *int) error {	
	return self.body(pos, thread, self, pc)
}

func (self TFunc) Dump(out io.Writer) {
	var as, rs []string

	for _, t := range self.argTypes {
		as = append(as, t.Name())
	}

	for _, t := range self.resTypes {
		rs = append(rs, t.Name())
	}

	fmt.Fprintf(out, "Func(%v %v %v)", self.name, as, rs)
}

func (self *TFunc) SetBody(body FuncBody) {
	self.body = body
}

func (self *TFunc) String() string {	
	var out strings.Builder
	self.Dump(&out)
	return out.String()
}
