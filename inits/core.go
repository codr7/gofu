package inits

import (
	"github.com/codr7/gofu"
	"github.com/codr7/gofu/errors"
	"github.com/codr7/gofu/forms"
	"github.com/codr7/gofu/funcs"
	"github.com/codr7/gofu/ops"
	"github.com/codr7/gofu/types"
)

func Core(scope *gofu.TScope) {
	scope.BindSlot("Any", types.Meta(), types.Any())
	scope.BindSlot("Bool", types.Meta(), types.Bool())
	scope.BindSlot("Char", types.Meta(), types.Char())
	scope.BindSlot("Func", types.Meta(), types.Func())
	scope.BindSlot("Int", types.Meta(), types.Int())
	scope.BindSlot("Meta", types.Meta(), types.Meta())
	scope.BindSlot("Multi", types.Meta(), types.Multi())
	scope.BindSlot("Num", types.Meta(), types.Num())
	scope.BindSlot("Seq", types.Meta(), types.Seq(types.Any()))
	scope.BindSlot("Stack", types.Meta(), types.Stack(types.Any()))
	scope.BindSlot("String", types.Meta(), types.String())
	scope.BindSlot("Target", types.Meta(), types.Target())

	scope.BindSlot("_", types.Nil(), nil)
	scope.BindSlot("t", types.Bool(), true)
	scope.BindSlot("f", types.Bool(), false)

	scope.BindSlot("bind",
		types.Macro(),
		gofu.Macro("bind", 2,
			func(pos gofu.TPos, args []gofu.Form, scope *gofu.TScope, block *gofu.TBlock) error {
				id := args[0].(forms.TId).Name()
				v := args[1]

				if s := v.Slot(scope); s == nil {
					i := scope.BindId(id, nil)
					
					if i == -1 {
						return errors.Compile(pos, "Duplicate binding: %v", id)
					}
					
					if err := v.Compile(scope, block); err != nil {
						return err
					}
				
					block.Emit(ops.BindId(pos, i, nil))
				} else {
					scope.BindSlot(id, s.Type(), s.Value())
				}
				
				return nil
			}))

	scope.BindSlot("d",
		types.Macro(),
		gofu.Macro("d", 0,
			func(pos gofu.TPos, args []gofu.Form, scope *gofu.TScope, block *gofu.TBlock) error {
				block.Emit(ops.Drop(pos, 1))
				return nil
			}))

	scope.BindSlot("func",
		types.Macro(),
		gofu.Macro("func", 4,
			func(pos gofu.TPos, args []gofu.Form, scope *gofu.TScope, block *gofu.TBlock) error {
				id := args[0].(forms.TId).Name()
				afs, rfs := args[1], args[2]
				var ats, rts []gofu.Type

				for _, f := range afs.(forms.TGroup).Members() {
					found := scope.Find(f.(forms.TId).Name())
					s := found.(gofu.TSlot)
					ats = append(ats, s.Value().(gofu.Type))
				}

				for _, f := range rfs.(forms.TGroup).Members() {
					found := scope.Find(f.(forms.TId).Name())
					s := found.(gofu.TSlot)
					rts = append(rts, s.Value().(gofu.Type))
				}
				f := gofu.Func(id, ats, rts, nil)

				if !scope.BindSlot(id, types.Func(), f) {
					return errors.Compile(pos, "Duplicate binding: %v", id)
				}

				body, err := funcs.CompileBody(args[3], block)

				if err != nil {
					return err
				}

				f.SetBody(body)
				return nil
			}))
	
	scope.BindSlot("if",
		types.Macro(),
		gofu.Macro("if", 3,
			func(pos gofu.TPos, args []gofu.Form, scope *gofu.TScope, block *gofu.TBlock) error {
				cond := args[0]

				if err := cond.Compile(scope, block); err != nil {
					return err
				}

				branch := block.Emit(ops.Nop())
				
				x, y := args[1], args[2]

				if err := x.Compile(scope, block); err != nil {
					return err
				}

				skip := block.Emit(ops.Nop())
				block.Set(branch, ops.Branch(block.Pc()))

				if err := y.Compile(scope, block); err != nil {
					return err
				}

				block.Set(skip, ops.Goto(block.Pc()))
				return nil
			}))

	scope.BindSlot("reset",
		types.Macro(),
		gofu.Macro("reset", 0,
			func(pos gofu.TPos, args []gofu.Form, scope *gofu.TScope, block *gofu.TBlock) error {
				block.Emit(ops.Reset())
				return nil
			}))
	
	scope.BindSlot("stack",
		types.Func(),
		gofu.Func("stack", nil, []gofu.Type{types.Stack(types.Any())},
			func(pos gofu.TPos, thread *gofu.TThread, _func *gofu.TFunc, pc *int) error {
				stack := thread.Stack()
				stack.Push(types.Stack(types.Any()), stack)
				return nil
			}))
}
