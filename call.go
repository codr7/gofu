package gofu

import (
)

type TCall struct {
	pos TPos
	target Target
	
	registers []Slot
	stack Stack
	returnPc int
}

func Call(pos TPos, tgt Target) TCall {
	return TCall{pos: pos, target: tgt, returnPc: -1}
}

func (self TCall) ReturnPc() int {
	return self.returnPc
}

func (self *TCall) Enter(scope *Scope, thread *TThread, pc *int) {
	self.registers = thread.registers
	self.stack = thread.stack

	thread.registers = make([]Slot, scope.RegisterCount())
	thread.stack.Init()

	self.returnPc = *pc
}

func (self TCall) Exit(thread *TThread, pc *int) {
	thread.registers = self.registers
	prevStack := thread.stack
	thread.stack = self.stack
	thread.stack.items = append(thread.stack.items, prevStack.items...)
	*pc = self.returnPc
}
