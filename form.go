package gofu

type Form interface {
	Compile(scope *TScope, block *TBlock) error
	Slot() *TSlot
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
