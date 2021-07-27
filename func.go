package gofu

import (
)

type Fimp = func(pos TPos, thread *TThread, pc *int) error

type TFunc struct {
	name string
	argTypes []Type
	resTypes []Type
	body Fimp
}

func Func(name string, aTypes []Type, rTypes []Type, body Fimp) *TFunc {
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
	offs := self.ArgCount()-1

	for i, t := range(self.argTypes) {
		if !Isa(stack.Peek(offs-i).Type(), t) {
			return false
		}
	}

	return true
}

func (self TFunc) Call(pos TPos, thread *TThread, pc *int) error {
	return self.body(pos, thread, pc)
}
