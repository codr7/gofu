package gofu

type TRegister struct {
	index int
	_type Type
}

func Register(i int, t Type) TRegister {
	return TRegister{index: i, _type: t}
}

func (self TRegister) Index() int {
	return self.index
}

func (self TRegister) Type() Type {
	return self._type
}
