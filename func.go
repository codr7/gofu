package gofu

type FuncImp = func(stack *Stack) error

type Func struct {
	name string
	argumentTypes []Type
	returnType  Type
	implementation  FuncImp
}

func NewFunc(name string, aTypes []Type, rType Type, imp FuncImp) *Func {
	return &Func{name: name, argumentTypes: aTypes, returnType: rType, implementation: imp}
}

func (self Func) Name() string {
	return self.name
}

func (self Func) ArgumentTypes() []Type {
	return self.argumentTypes
}

func (self Func) Call(stack *Stack) error {
	return self.implementation(stack)
}

