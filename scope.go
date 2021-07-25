package gofu

import (
	"github.com/codr7/gollies"
)

type Scope struct {
	bindings gollies.SliceMap
	registerCount int
}

func (self *Scope) Init() {
	self.bindings.Init(gollies.CompareString)
}

func (self *Scope) RegisterCount() int {
	return self.registerCount
}

func (self *Scope) BindId(id string) int {
	i := self.registerCount

	if self.bindings.Add(id, i) != nil {
		return -1
	}

	self.registerCount++
	return i
}

func (self *Scope) BindSlot(id string, t Type, v interface{}) bool {
	var s Slot
	s.Init(t, v)
	return self.bindings.Add(id, s) == nil
}

func (self Scope) Find(id string) interface{} {
	return self.bindings.Find(id)
}
