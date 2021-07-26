package gofu

import (
	"strings"
)

type Stack struct {
	items []Slot
}

func (self *Stack) Init() {
	self.items = nil
}

func (self Stack) Empty() bool {
	return len(self.items) == 0
}

func (self Stack) Len() int {
	return len(self.items)
}

func (self *Stack) Push(t Type, v interface{}) {
	var s Slot
	s.Init(t, v)
	self.items = append(self.items, s)
}

func (self Stack) Peek(offs int) Slot {
	return self.items[len(self.items)-offs-1]
}

func (self *Stack) Pop() Slot {
	i := len(self.items)-1
	it := self.items[i]
	self.items = self.items[:i]
	return it
}

func (self Stack) String() string {
	var out strings.Builder
	out.WriteRune('[')

	for i, s := range self.items {
		if i > 0 {
			out.WriteRune(' ')
		}

		s.Type().DumpValue(s.Value(), &out)
	}
	
	out.WriteRune(']')
	return out.String()
}
