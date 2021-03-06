package gofu

import (
	"github.com/codr7/gollies"
)

type TScope struct {
	bindings gollies.SliceMap
	registerCount int
}

func Scope() *TScope {
	return new(TScope).Init()
}

func (self *TScope) Init() *TScope {
	self.bindings.Init(gollies.CompareString)
	return self
}

func (self *TScope) RegisterCount() int {
	return self.registerCount
}

func (self *TScope) BindId(id string, t Type) int {
	i := self.registerCount

	if self.bindings.Add(id, Register(i, t)) != nil {
		return -1
	}

	self.registerCount++
	return i
}

func (self *TScope) BindSlot(id string, t Type, v interface{}) bool {
	var s TSlot
	s.Init(t, v)
	return self.bindings.Add(id, s) == nil
}

func (self TScope) Find(id string) interface{} {
	return self.bindings.Find(id)
}
