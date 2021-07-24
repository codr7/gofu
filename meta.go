package gofu

import (
	"fmt"
)

type TMeta struct {
	name string
	arity int
	items []*TFunc
}

func Meta(name string, arity int) *TMeta {
	return &TMeta{name: name, arity: arity}
}

func (self *TMeta) Arity() int {
	return self.arity;
}

func (self *TMeta) Push(_func *TFunc) {
	self.items = append(self.items, _func)
}

func (self *TMeta) Pop() *TFunc {
	i := len(self.items)-1
	it := self.items[i]
	self.items = self.items[:i]
	return it
}

func (self *TMeta) GetFunc(args []Form) Target {
	//TODO Return first matching implementation from end if possible,
	//if and only if full match or nothing else matches
	return self
}

func (self *TMeta) Call(stack *Stack) error {
	for i := len(self.items)-1; i >= 0; i++ {
		if f := self.items[i]; f.Check(stack) {
			return f.Call(stack)
		}
	}
	
	return fmt.Errorf("No matching function found: %v", self.name)
}

