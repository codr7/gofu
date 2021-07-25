package gofu

type Target interface {
	Arity() int
	Call(pos TPos, pc *int, stack *Stack) error
}
