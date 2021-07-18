package gofu

type Slot struct {
	Type Type
	Value interface{}
}

func (self *Slot) Init(t Type, v interface{}) *Slot {
	self.Type = t
	self.Value = v
	return self
}
