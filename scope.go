package gofu

import (
	"github.com/codr7/gollies"
)

type Scope struct {
	bindings gollies.SliceMap
	stackDepth int
}

func (self *Scope) Init() {
	self.bindings.Init(gollies.CompareString)
}

func (self *Scope) StackDepth() int {
	return self.stackDepth
}

func (self *Scope) Push() int {
	i := self.stackDepth
	self.stackDepth++
	return i
}

func (self *Scope) BindId(id string) int {
	i := self.Push()
	self.bindings.Add(id, i)
	return i
}

func (self *Scope) BindSlot(id string, t Type, v interface{}) {
	var s Slot
	s.Init(t, v)
	self.bindings.Add(id, s)
}

func (self Scope) Find(id string) interface{} {
	return self.bindings.Find(id)
}
