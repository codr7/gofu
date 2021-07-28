package gofu

type Target interface {
	Name() string
	ArgCount() int
	Applicable(stack *TStack) bool
	Call(pos TPos, thread *TThread, pc *int, check bool) error
}
