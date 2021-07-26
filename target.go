package gofu

type Target interface {
	Arity() int
	Call(pos TPos, thread *TThread, pc *int) error
}
