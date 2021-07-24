package gofu

type Form interface {
	Compile(scope *Scope, block *Block) error
}
