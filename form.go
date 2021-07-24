package gofu

type Form interface {
	Compile(scope *Scope, block *Block) error
}

type BasicForm struct {
	pos TPos
}

func (self *BasicForm) Init(pos TPos) *BasicForm {
	self.pos = pos
	return self
}

func (self BasicForm) Pos() TPos {
	return self.pos
}
