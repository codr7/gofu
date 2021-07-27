package gofu

import (
	"fmt"
)

type TCall struct {
	pos TPos
	target Target
	
	registers []TSlot
	stack TStack
	returnPc int
}

func Call(pos TPos, tgt Target) TCall {
	return TCall{pos: pos, target: tgt, returnPc: -1}
}

func (self TCall) Target() Target {
	return self.target
}

func (self TCall) ReturnPc() int {
	return self.returnPc
}

func (self *TCall) Enter(scope *TScope, thread *TThread, pc *int) {
	self.registers = thread.registers
	self.stack = thread.stack

	thread.registers = make([]TSlot, scope.RegisterCount())
	thread.stack.Init(nil)

	self.returnPc = *pc
}

func (self TCall) Exit(thread *TThread, pc *int) error {
	thread.registers = self.registers
	prevStack := thread.stack
	thread.stack = self.stack
	f := self.target.(*TFunc)
	rc := f.ResCount()

	if prevStack.Len() < rc {
		return fmt.Errorf("Not enough results on stack: %v/%v", f.Name(), prevStack)
	}
	
	thread.stack.items = append(thread.stack.items, prevStack.items[:rc]...)
	*pc = self.returnPc
	return nil
}
