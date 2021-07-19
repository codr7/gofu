package gofu

type Form interface {
	Emit(block *Block)
}
