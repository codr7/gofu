package gofu

type Slot struct {
	_type Type
	value interface{}
}

func (self *Slot) Init(t Type, v interface{}) *Slot {
	self._type = t
	self.value = v
	return self
}

func (self *Slot) Type() Type {
	return self._type
}

func (self *Slot) Value() interface{} {
	return self.value
}
