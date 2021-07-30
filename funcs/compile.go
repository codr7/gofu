package funcs

import (
	"github.com/codr7/gofu"
	"github.com/codr7/gofu/ops"
)

func CompileBody(body gofu.Form, block *gofu.TBlock) (gofu.FuncBody, error) {
	scope := gofu.Scope()

	skip := block.Emit(ops.Nop())
	startPc := block.Pc()
	
	if err := body.Compile(scope, block); err != nil {
		return nil, err
	}

	block.Emit(ops.Return(body.Pos()))
	block.Set(skip, ops.Goto(block.Pc()))
	
	return func(pos gofu.TPos, thread *gofu.TThread, _func *gofu.TFunc, pc *int) error {
		thread.PushCall(pos, _func).Enter(scope, thread, pc)
		*pc = startPc
		return nil
	}, nil
	
}
