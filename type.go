package gofu

import (
	"fmt"
	"io"
)

type Type interface {
	Name() string
	DumpValue(val interface{}, out io.Writer)
	HasParent(parent Type) bool
}

type BType struct {
	name string
	parentTypes []Type
}

func (self *BType) Init(name string) *BType {
	self.name = name
	return self
}

func (self BType) Name() string {
	return self.name
}

func (self BType) String() string {
	return self.name
}

func (self BType) DumpValue(val interface{}, out io.Writer) {
	fmt.Fprintf(out, "%v", val)
}

func (self BType) HasParent(parent Type) bool {
	for _, t := range self.parentTypes {
		if Isa(t, parent) {
			return true
		}
	}

	return false
}

func (self BType) AddParent(parent Type) {
	self.parentTypes = append(self.parentTypes, parent)
}

func Isa(child, parent Type) bool {
	if child == parent {
		return true
	}

	return child.HasParent(parent)
}
