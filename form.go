package gofu

type Form interface {
	Compile(scope *Scope, block *Block) error
}

type BForm struct {
	pos TPos
}

func (self *BForm) Init(pos TPos) *BForm {
	self.pos = pos
	return self
}

func (self BForm) Pos() TPos {
	return self.pos
}
