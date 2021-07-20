package gofu

type Form interface {
	Emit(scope *Scope, block *Block)
}
