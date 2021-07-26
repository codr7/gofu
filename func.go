package gofu

import (
)

type Fimp = func(pos TPos, thread *TThread, pc *int) error

type TFunc struct {
	name string
	argumentTypes []Type
	returnType Type
	body Fimp
}

func Func(name string, aTypes []Type, rType Type, body Fimp) *TFunc {
	return &TFunc{name: name, argumentTypes: aTypes, returnType: rType, body: body}
}

func (self TFunc) Name() string {
	return self.name
}

func (self TFunc) Arity() int {
	return len(self.argumentTypes)
}

func (self TFunc) ArgumentTypes() []Type {
	return self.argumentTypes
}

func (self TFunc) Applicable(thread *TThread) bool {
	stack := thread.Stack()
	offs := self.Arity()-1

	for i, t := range(self.argumentTypes) {
		if !Isa(stack.Peek(offs-i).Type(), t) {
			return false
		}
	}

	return true
}

func (self TFunc) Call(pos TPos, thread *TThread, pc *int) error {
	return self.body(pos, thread, pc)
}
