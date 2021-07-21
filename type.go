package gofu

import (
	"fmt"
	"io"
)

type Type interface {
	Name() string
	DumpValue(val interface{}, out io.Writer)
}

type BasicType struct {
	name string
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
