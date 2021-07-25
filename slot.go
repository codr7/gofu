package gofu

import (
	"bytes"
)

type Slot struct {
	_type Type
	value interface{}
}

func (self *Slot) Init(t Type, v interface{}) {
	self._type = t
	self.value = v
}

func (self Slot) Type() Type {
	return self._type
}

func (self Slot) Value() interface{} {
	return self.value
}

func (self Slot) String() string {
	var out bytes.Buffer
	self._type.DumpValue(self.value, &out)
	out.WriteRune(':')
	out.WriteString(self._type.Name())
	return out.String()
}
