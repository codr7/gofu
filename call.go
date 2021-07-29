package gofu

import (
	"fmt"
)

type TCall struct {
	pos TPos
	_func *TFunc
	
	registers []TSlot
	stack TStack
	returnPc int
}

func Call(pos TPos, tgt *TFunc) TCall {
	return TCall{pos: pos, _func: tgt, returnPc: -1}
}

func (self TCall) Func() *TFunc {
	return self._func
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
	rc := self._func.ResCount()

	if prevStack.Len() < rc {
		return fmt.Errorf("Not enough results on stack: %v/%v", self._func, prevStack)
	}
	
	thread.stack.items = append(thread.stack.items, prevStack.items[:rc]...)
	*pc = self.returnPc
	return nil
}
