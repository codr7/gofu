package gofu

import (
	"fmt"
	"io"
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
	n := len(self.items)
	
	if n == 0 {
		return nil
	}

	return &self.items[n-offs-1]
}

func (self *TStack) Pop() *TSlot {
	n := len(self.items)

	if n == 0 {
		return nil
	}

	i := n-1
	it := self.items[i]
	self.items = self.items[:i]
	return &it
}

func (self *TStack) Drop(n int) bool {
	max := len(self.items)

	if max < n {
		return false
	}

	self.items = self.items[:max-n]
	return true
}

func (self *TStack) Dump(out io.Writer) {
	fmt.Fprint(out, "[")

	for i, s := range self.items {
		if i > 0 {
			fmt.Fprint(out, " ")
		}

		if s.Value() == self {
			fmt.Fprint(out, "[^]")
		}
		
		s.Type().DumpValue(s.Value(), out)
	}
	
	fmt.Fprint(out, "]")
}

func (self *TStack) String() string {
	var out strings.Builder
	self.Dump(&out)
	return out.String()
}
