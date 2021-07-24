package gofu

type Op interface {
	Eval(pc *int, calls *CallStack, stack *Stack) error
}
