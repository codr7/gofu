package gofu

import (
	"fmt"
	"io"
)

type Type interface {
	Name() string
	DumpValue(val interface{}, out io.Writer)
	Isa(parent Type) bool
}

type BasicType struct {
	name string
	parentTypes []Type
}

func (self *BasicType) Init(name string) *BasicType {
	self.name = name
	return self
}

func (self BasicType) Name() string {
	return self.name
}

func (self BasicType) String() string {
	return self.name
}

func (self BasicType) DumpValue(val interface{}, out io.Writer) {
	fmt.Fprintf(out, "%v", val)
}

func (self BasicType) Isa(parent Type) bool {
	for _, t := range self.parentTypes {
		if t == parent || t.Isa(parent) {
			return true
		}
	}

	return false
}
