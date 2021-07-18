package gofu

type Op interface {
	Eval(pc int, stack *Stack) (int, error)
}
