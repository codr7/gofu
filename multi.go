package gofu

import (
	"fmt"
)

type TMulti struct {
	name string
	arity int
	items []*TFunc
}

func Multi(name string, arity int) *TMulti {
	return &TMulti{name: name, arity: arity}
}

func (self *TMulti) Arity() int {
	return self.arity;
}

func (self *TMulti) Push(_func *TFunc) error {
	if x, y := _func.Arity(), self.arity; x != y {
		return fmt.Errorf("Wrong arity for multi: %v/%v/%v", self.name, x, y)
	}
	
	self.items = append(self.items, _func)
	return nil
}

func (self *TMulti) Pop() *TFunc {
	i := len(self.items)-1
	it := self.items[i]
	self.items = self.items[:i]
	return it
}

func (self *TMulti) GetFunc(args []Form) Target {
	//TODO Return first matching implementation from end if possible,
	//if and only if full match or nothing else matches
	return self
}

func (self *TMulti) Call(pos TPos, pc *int, stack *Stack) error {
	for i := len(self.items)-1; i >= 0; i++ {
		if f := self.items[i]; f.Applicable(stack) {
			return f.Call(pos, pc, stack)
		}
	}
	
	return fmt.Errorf("No matching function found: %v", self.name)
}

