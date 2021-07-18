package gofu

type Type interface {
	Name() string
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

