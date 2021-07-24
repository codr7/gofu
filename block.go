package gofu

import (
	"github.com/codr7/gofu/errors"
)

type Block struct {
	ops []Op
}

func (self Block) Step(pc *int, calls *CallStack, stack *Stack) error {
	return self.ops[*pc].Eval(pc, calls, stack)
}

func (self Block) Run(pc int, calls *CallStack, stack *Stack) error {
	var err error
	
	for {
		if err = self.ops[pc].Eval(&pc, calls, stack); err != nil {
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
