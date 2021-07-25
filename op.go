package gofu

type Op interface {
	Eval(pc *int, calls *CallStack, registers []Slot, stack *Stack) error
}
