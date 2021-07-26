package gofu

import (
	"github.com/codr7/gollies"
)

type TScope struct {
	bindings gollies.SliceMap
	registerCount int
}

func (self *TScope) Init() {
	self.bindings.Init(gollies.CompareString)
}

func (self *TScope) RegisterCount() int {
	return self.registerCount
}

func (self *TScope) BindId(id string) int {
	i := self.registerCount

	if self.bindings.Add(id, i) != nil {
		return -1
	}

	self.registerCount++
	return i
}

func (self *TScope) BindSlot(id string, t Type, v interface{}) bool {
	var s Slot
	s.Init(t, v)
	return self.bindings.Add(id, s) == nil
}

func (self TScope) Find(id string) interface{} {
	return self.bindings.Find(id)
}
