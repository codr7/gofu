package gofu

type Block struct {
	ops []Op
}

func (self Block) Eval(pc int, stack *Stack) error {
	var err error

	for {
		op := self.ops[pc]
		pc, err = op.Eval(pc, stack)
		
		if pc == -1 {
			break
		}

	}

	return err
}

func (self *Block) Emit(op Op) {
	self.ops = append(self.ops, op)
}
