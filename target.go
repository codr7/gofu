package gofu

type Target interface {
	ArgumentTypes() []Type
	Call(stack *Stack) error
}
