# gofu
## a scripting language toolkit in Go

### intro
[gofu](https://github.com/codr7/gofu) aims to provide a flexible toolkit for creating custom scripting languages in Go.

### functions
Functions have a name, an argument list, a result list and a body.

```go
p := gofu.Pos("Test", -1, -1)

add := gofu.Func("+", []gofu.Type{types.Int(), types.Int()}, []gofu.Type{types.Int()},
	func(pos gofu.TPos, thread *gofu.TThread, _func *gofu.TFunc, pc *int, check bool) error {
		stack := thread.Stack()
		stack.Push(types.Int(), stack.Pop().Value().(int) + stack.Pop().Value().(int))
		return nil
	})

scope := gofu.Scope()	
scope.BindSlot("+", types.Func(), add)

block := gofu.Block()
c := forms.Call(p, forms.Id(p, "+"), forms.Literal(p, types.Int(), 35), forms.Literal(p, types.Int(), 7))
c.Compile(scope, block)
block.Emit(ops.Stop())

thread := gofu.Thread(scope)
block.Run(thread, 0)
```

The same thing could be accomplished by manually emitting operations.

```go
block.Emit(ops.Push(types.Int(), 35))
block.Emit(ops.Push(types.Int(), 7))
block.Emit(ops.Call(p, add))
block.Emit(ops.Stop())
```

`fimps.Compile` may be used to compile function bodies.

```go
fimp, err := fimps.Compile(forms.Literal(p, []gofu.Type{types.Int()}, 42), block)
fortyTwo := gofu.Func("fortyTwo", nil, []gofu.Type{types.Int()}, fimp)
scope.BindSlot("fortyTwo", types.Func(), f)
```

#### multiple dispatch
The following example will dispatch to the right function based on the argument and push `"Bool!"` on the stack.

```go
f1 := gofu.Func("foo", []gofu.Type{types.Bool()}, []gofu.Type{types.Int()},
    func(pos gofu.TPos, thread *gofu.TThread, _func *gofu.TFunc, pc *int, check bool) error {
	    stack := thread.Stack()
	    stack.Pop()
	    stack.Push(types.String(), "Bool!")
	    return nil
    })

f2 := gofu.Func("foo", []gofu.Type{types.Int()}, []gofu.Type{types.Int()},
    func(pos gofu.TPos, thread *gofu.TThread, _func *gofu.TFunc, pc *int, check bool) error {
	    stack := thread.Stack()
	    stack.Pop()
	    stack.Push(types.String(), "Int!")
	    return nil
    })

m := gofu.Multi("foo", 1, f1, f2)
block.Emit(ops.Push(types.Bool(), true))
block.Emit(ops.Call(p, m))
block.Emit(ops.Stop())	
```

### macros
Macros are called at compile time and may emit different code depending on arguments.

```
scope.BindSlot("reset",
	types.Macro(),
	gofu.Macro("reset", 0,
		func(pos gofu.TPos, args []gofu.Form, scope *gofu.TScope, block *gofu.TBlock) error {
			block.Emit(ops.Reset())
			return nil
		}))
```

### types
The following list of types are provided but optional, anything implementing `gofu.Type` may be used as a type.

* Any: Any - Anything
* Bool: Any - t/f
* Char: Any - Characters
* Func: Any Target - Functions
* Int: Any Num - Integers
* Maybe[T]: Any - Contains T or Nil
* Meta: Any - The type of types
* Multi: Any Target - Multimethods
* Nil - Nothing, it's only value being `_`
* Num: Any - Parent of all numbers
* Seq[T]: Any - Parent of all sequences
* Stack[T]: Any Seq[T] - Stacks of values
* String: Any Seq[Char] - Strings
* Target: Any - Callable values

### repl
A primitive [REPL](https://github.com/codr7/gofu/blob/main/utils/repl.go) is provided, it reads one form at a time and prints the stack after each evaluation.

```
$ cd bin
$ ./mk
$ ./repl
gofu v1
  +(35 7)
[42]
```

Parens may be used to group forms.

```
  (1 2 3)
[1 2 3]
```

The stack may be directly modified using `d` and `reset`.

```
  (1 2 3)
[1 2 3]
  d
[1 2]
  reset
[]
```

Functions may be called by suffixing names with argument lists.

```
  +(35 7)
[42]
```

New functions may be defined using `func`.

```
  func(foo () (Int) 42)
[]
  foo
[42]
```

Values may be bound to identifiers using `bind`.

```
  bind(foo 42)
[]
  foo
[42]
```

The REPL is heavily parameterized and assumes very little about the actual language.

```
scope := gofu.Scope()
block := gofu.Block()
thread := gofu.Thread(scope)
utils.Repl(scope, parsers.Any(), block, thread)
```