package gofu

type Stack struct {
	items []Slot
}

func (self *Stack) Init(n int) *Stack {
	self.items = make([]Slot, n)
	return self
}

func (self *Stack) Push(it Slot) {
	self.items = append(self.items, it)
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

func (self *Stack) Get(index int) Slot {
	return self.items[index]
}

func (self *Stack) Set(index int, t Type, v interface{}) {
	self.items[index].Init(t, v)
}
