package gofu

type Op interface {
	Eval(thread *TThread, pc *int) error
}
