package gofu

import (
	"fmt"
)

type Call struct {
	pos TPos
	target Target
}

type CallStack struct {
	items []Call
}

func (self *CallStack) Push(pos TPos, tgt Target) int {
	tag := len(self.items)
	self.items = append(self.items, Call{pos: pos, target: tgt})
	return tag
}

func (self *CallStack) Pop(tag int) error {
	n := len(self.items)
	
	if x, y := n, tag+1; x != y {
		return fmt.Errorf("Invalid call stack: %v/%v", x, y)
	}

	self.items = self.items[:n-1]
	return nil
}
