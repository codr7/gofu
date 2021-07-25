package gofu

type Fimp = func(stack *Stack) error

type TFunc struct {
	name string
	argumentTypes []Type
	returnType Type
	implementation Fimp
}

func Func(name string, aTypes []Type, rType Type, imp Fimp) *TFunc {
	return &TFunc{name: name, argumentTypes: aTypes, returnType: rType, implementation: imp}
}

func (self TFunc) Name() string {
	return self.name
}

func (self TFunc) Arity() int {
	return len(self.argumentTypes)
}

func (self TFunc) ArgumentTypes() []Type {
	return self.argumentTypes
}

func (self TFunc) Applicable(stack *Stack) bool {
	offs := self.Arity()-1

	for i, t := range(self.argumentTypes) {
		if !Isa(stack.Peek(offs-i).Type(), t) {
			return false
		}
	}

	return true
}

func (self TFunc) Call(stack *Stack) error {
	return self.implementation(stack)
}
