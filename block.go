package gofu

import (
	"github.com/codr7/gofu/errors"
)

type Block struct {
	ops []Op
}

func (self Block) Pc() int {
	return len(self.ops)
}

func (self Block) Step(thread *TThread, pc *int) error {
	return self.ops[*pc].Eval(thread, pc)
}

func (self Block) Run(thread *TThread, pc int) error {
	var err error
	
	for {
		if err = self.ops[pc].Eval(thread, &pc); err != nil {
			if err == errors.Stop {
				err = nil
			}
			
			break
		}
	}

	return err
}

func (self *Block) Emit(op Op) int {
	i := len(self.ops)
	self.ops = append(self.ops, op)
	return i
}

func (self Block) Set(idx int, op Op) {
	self.ops[idx] = op
}


func (self *Block) Peek() *Op {
	n := len(self.ops)
	
	if n == 0 {
		return nil
	}
	
	return &self.ops[n-1]
}
