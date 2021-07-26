package gofu

import (
	"fmt"
)

type TMulti struct {
	name string
	argCount int
	funcs []*TFunc
}

func Multi(name string, argCount int, funcs...*TFunc) *TMulti {
	var m = TMulti{name: name, argCount: argCount}

	for _, f := range(funcs) {
		m.Push(f)
	}
	
	return &m
}

func (self *TMulti) ArgCount() int {
	return self.argCount;
}

func (self *TMulti) Push(_func *TFunc) error {
	if x, y := _func.ArgCount(), self.argCount; x != y {
		return fmt.Errorf("Wrong arg count for multi: %v/%v/%v", self.name, x, y)
	}
	
	self.funcs = append(self.funcs, _func)
	return nil
}

func (self *TMulti) Pop() *TFunc {
	i := len(self.funcs)-1
	it := self.funcs[i]
	self.funcs = self.funcs[:i]
	return it
}

func (self *TMulti) GetFunc(args []Form) Target {
	//TODO Return first matching implementation from end if possible,
	//if and only if full match or nothing else matches
	return self
}

func (self *TMulti) Call(pos TPos, thread *TThread, pc *int) error {
	for i := len(self.funcs)-1; i >= 0; i-- {
		if f := self.funcs[i]; f.Applicable(thread) {
			return f.Call(pos, thread, pc)
		}
	}
	
	return fmt.Errorf("No matching function found: %v", self.name)
}

