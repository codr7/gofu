package gofu

import (
	"fmt"
	"io"
)

type Type interface {
	Name() string

	AddParent(parent Type, rec bool)
	AddParentsTo(child Type)
	Isa(parent Type) Type

	TrueValue(val interface{}) bool
	DumpValue(val interface{}, out io.Writer)
}

type BType struct {
	name string
	parentTypes map[Type]Type
}

func (self *BType) Init(name string) *BType {
	self.name = name
	self.parentTypes = make(map[Type]Type)
	return self
}

func (self BType) Name() string {
	return self.name
}

func (self BType) DumpValue(val interface{}, out io.Writer) {
	fmt.Fprintf(out, "%v", val)
}

func (self BType) TrueValue(val interface{}) bool {
	return true
}

func (self BType) Isa(parent Type) Type {
	return self.parentTypes[parent]
}

func (self *BType) AddParent(parent Type, rec bool) {
	self.parentTypes[parent] = parent

	if (rec) {
		parent.AddParentsTo(self)
	}
}

func (self BType) AddParentsTo(child Type) {
	for p, _ := range(self.parentTypes) {
		child.AddParent(p, false)
	}
}

func Isa(child, parent Type) bool {
	if child == parent {
		return true
	}

	return child.Isa(parent) != nil
}
