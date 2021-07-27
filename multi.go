package gofu

import (
	"fmt"
	"log"
)

type TMulti struct {
	name string
	argCount int
	funcs []*TFunc
}

func Multi(name string, argCount int, funcs...*TFunc) *TMulti {
	if argCount == 0 {
		log.Fatal("Invalid arg count: 0")
	}
	
	var m = TMulti{name: name, argCount: argCount}

	for _, f := range(funcs) {
		if err := m.Push(f); err != nil {
			log.Fatal(err)
		}
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

func (self *TMulti) GetFunc(args []Form, scope *TScope) Target {
	var matches []*TFunc
	nfs := len(self.funcs)
	
	for fi := nfs-1; fi >= 0; fi-- {
		f := self.funcs[fi]
		match := true
		
		for ai, a := range(args) {
			if s := a.Slot(scope); s != nil && !Isa(s.Type(), f.argTypes[ai]) {
				match = false
			}

			if !match {
				break
			}
		}

		if match {
			matches = append(matches, f)
		}
	}

	nms := len(matches) 

	if nms == 1 {
		return matches[0]
	}

	if nms > 0 && nms < nfs {
		return Multi(self.name, self.argCount, matches...)
	}
	
	return self
}

func (self *TMulti) Call(pos TPos, thread *TThread, pc *int) error {
	stack := thread.Stack()
	
	for i := len(self.funcs)-1; i >= 0; i-- {
		if f := self.funcs[i]; f.Applicable(stack) {
			return f.Call(pos, thread, pc)
		}
	}
	
	return fmt.Errorf("No matching function found: %v", self.name)
}

