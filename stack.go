package gofu

import (
	"strings"
)

type TStack struct {
	items []TSlot
}

func Stack(items...TSlot) *TStack {
	return new(TStack).Init(items)
}

func (self *TStack) Init(items []TSlot) *TStack {
	self.items = items
	return self
}

func (self TStack) Empty() bool {
	return len(self.items) == 0
}

func (self TStack) Len() int {
	return len(self.items)
}

func (self *TStack) Push(t Type, v interface{}) {
	var s TSlot
	s.Init(t, v)
	self.items = append(self.items, s)
}

func (self TStack) Peek(offs int) *TSlot {
	if len(self.items) == 0 {
		return nil
	}

	return &self.items[len(self.items)-offs-1]
}

func (self *TStack) Pop() *TSlot {
	if len(self.items) == 0 {
		return nil
	}

	i := len(self.items)-1
	it := self.items[i]
	self.items = self.items[:i]
	return &it
}

func (self TStack) String() string {
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
