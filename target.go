package gofu

type Target interface {
	Arity() int
	Call(stack *Stack) error
}
