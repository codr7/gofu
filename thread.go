package gofu

import (
)

type TThread struct {
	calls []TCall
	registers []Slot
	stack Stack
}

func Thread(scope *Scope) *TThread {
	return new(TThread).Init(scope)
}

func (self *TThread) Init(scope *Scope) *TThread {
	self.registers = make([]Slot, scope.RegisterCount())
	return self
}

func (self *TThread) PopCall() *TCall {
	if len(self.calls) == 0 {
		return nil
	}

	i := len(self.calls)-1
	c := self.calls[i]
	self.calls = self.calls[:i]
	return &c
}

func (self *TThread) PushCall(pos TPos, tgt Target) {
	self.calls = append(self.calls, Call(pos, tgt))
}

func (self *TThread) PeekCall() *TCall {
	n := len(self.calls)
	
	if n == 0 {
		return nil
	}
	
	return &self.calls[n-1]
}

func (self *TThread) Stack() *Stack {
	return &self.stack
}

func (self *TThread) Get(idx int) Slot {
	return self.registers[idx]
}

func (self *TThread) Set(idx int, t Type, v interface{}) {
	self.registers[idx].Init(t, v)
}
