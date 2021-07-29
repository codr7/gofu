package gofu

import (
	"bytes"
)

type TSlot struct {
	_type Type
	value interface{}
}

func Slot(t Type, v interface{}) *TSlot {
	var s TSlot
	return s.Init(t, v)
}
	
func (self *TSlot) Init(t Type, v interface{}) *TSlot {
	self._type = t
	self.value = v
	return self
}

func (self TSlot) Type() Type {
	return self._type
}

func (self TSlot) Value() interface{} {
	return self.value
}

func (self TSlot) String() string {
	var out bytes.Buffer
	self._type.DumpValue(self.value, &out)
	return out.String()
}
