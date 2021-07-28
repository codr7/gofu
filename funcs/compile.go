package funcs

import (
	"github.com/codr7/gofu"
	"github.com/codr7/gofu/ops"
)

func CompileBody(body gofu.Form, block *gofu.TBlock) (gofu.FuncBody, error) {
	var scope gofu.TScope
	scope.Init()

	skip := block.Emit(ops.Nop())
	startPc := block.Pc()
	
	if err := body.Compile(&scope, block); err != nil {
		return nil, err
	}

	block.Emit(ops.Return(body.Pos()))
	block.Set(skip, ops.Goto(block.Pc()))
	
	return func(pos gofu.TPos, thread *gofu.TThread, pc *int) error {
		thread.PeekCall().Enter(&scope, thread, pc)
		*pc = startPc
		return nil
	}, nil
	
}
