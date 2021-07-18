package gofu

type Stack struct {
	items []Slot
}

func (self *Stack) Push(it Slot) {
	self.items = append(self.items, it)
}

func (self Stack) Peek(offs int) Slot {
	return self.items[len(self.items)-offs-1]
}
