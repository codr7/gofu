package gofu

import (
	"github.com/codr7/gofu/errors"
)

type Block struct {
	ops []Op
}

func (self Block) Eval(pc int, calls *CallStack, stack *Stack) error {
	var err error
	
	for {
		op := self.ops[pc]
		
		
		if err = op.Eval(&pc, calls, stack); err != nil {
			if err == errors.Stop {
				err = nil
			}
			
			break
		}
	}

	return err
}

func (self *Block) Emit(op Op) {
	self.ops = append(self.ops, op)
}
