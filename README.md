# gofu
## a scripting language toolkit in Go

### intro
[gofu](https://github.com/codr7/gofu) is an attempt at providing a flexible toolkit for building custom scripting languages in Go.

### functions
Functions have a name, an argument list, a result list and a body.

```
p := gofu.Pos("Test", -1, -1)

add := gofu.Func("+", []gofu.Type{types.Int(), types.Int()}, []gofu.Type{types.Int()},
	func(pos gofu.TPos, thread *gofu.TThread, pc *int) error {
		stack := thread.Stack()
		stack.Push(types.Int(), stack.Pop().Value().(int) + stack.Pop().Value().(int))
		return nil
	})

var scope gofu.TScope	
scope.Init()
scope.BindSlot("+", types.Func(), add)

var block gofu.TBlock	
c := forms.Call(p, forms.Id(p, "+"), forms.Literal(p, types.Int(), 35), forms.Literal(p, types.Int(), 7))
c.Compile(&scope, &block)
block.Emit(ops.Stop())

var thread gofu.TThread
thread.Init(&scope)
block.Run(&thread, 0)
```

The same thing could be accomplished by manually emitting operations.

```
...
block.Emit(ops.Push(types.Int(), 35))
block.Emit(ops.Push(types.Int(), 7))
block.Emit(ops.Call(p, add))
block.Emit(ops.Stop())
...
```

`fimp.Compile` may be used to compile function bodies.

```
fimp, err := fimp.Compile(forms.Literal(p, []gofu.Type{types.Int()}, 42), &block)
fortyTwo := gofu.Func("fortyTwo", nil, []gofu.Type{types.Int()}, fimp)
scope.BindSlot("fortyTwo", types.Func(), f)
```

#### multiple dispatch

```
f1 := gofu.Func("foo", []gofu.Type{types.Bool()}, []gofu.Type{types.Int()},
    func(pos gofu.TPos, thread *gofu.TThread, pc *int) error {
	    stack := thread.Stack()
	    stack.Pop()
	    stack.Push(types.Int(), 7)
	    return nil
    })

f2 := gofu.Func("foo", []gofu.Type{types.Int()}, []gofu.Type{types.Int()},
    func(pos gofu.TPos, thread *gofu.TThread, pc *int) error {
	    stack := thread.Stack()
	    stack.Pop()
	    stack.Push(types.Int(), 14)
	    return nil
    })

m := gofu.Multi("foo", 1, f1, f2)
block.Emit(ops.Push(types.Bool(), true))
block.Emit(ops.Call(p, m))
block.Emit(ops.Stop())	
```
