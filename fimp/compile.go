package fimp

import (
	"github.com/codr7/gofu"
	"github.com/codr7/gofu/ops"
)

func Compile(body gofu.Form, block *gofu.Block) (gofu.Fimp, error) {
	var scope gofu.Scope
	scope.Init()

	skip := block.Emit(ops.Nop())
	startPc := block.Pc()
	
	if err := body.Compile(&scope, block); err != nil {
		return nil, err
	}

	block.Emit(ops.Return())
	block.Set(skip, ops.Goto(block.Pc()))
	
	return func(pos gofu.TPos, thread *gofu.TThread, pc *int) error {
		thread.PeekCall().Enter(&scope, thread, pc)
		*pc = startPc
		return nil
	}, nil
	
}
