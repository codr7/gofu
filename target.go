package gofu

type Target interface {
	Arity() int
	Call(pos TPos, pc *int, registers []Slot, stack *Stack) error
}
