package gofu

type Target interface {
	ArgCount() int
	Call(pos TPos, thread *TThread, pc *int) error
}
