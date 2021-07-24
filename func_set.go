package gofu

import (
	"fmt"
)

type TFuncSet struct {
	name string
	arity int
	items []*TFunc
}

func FuncSet(name string, arity int) *TFuncSet {
	return &TFuncSet{name: name, arity: arity}
}

func (self *TFuncSet) Arity() int {
	return self.arity;
}

func (self *TFuncSet) Push(_func *TFunc) {
	self.items = append(self.items, _func)
}

func (self *TFuncSet) Pop() *TFunc {
	i := len(self.items)-1
	it := self.items[i]
	self.items = self.items[:i]
	return it
}

func (self *TFuncSet) GetFunc(args []Form) Target {
	//TODO Return first matching implementation if possible
	return self
}

func (self *TFuncSet) Call(stack *Stack) error {
	n := len(self.items)

	if n > 0 {
		//TODO Pick first matching implementation from end
		return self.items[n-1].Call(stack)
	}

	return fmt.Errorf("No matching function found: %v", self.name)
}

