# gofu
## a scripting language toolkit in Go

### intro
[gofu](https://github.com/codr7/gofu) aims to provide a flexible toolkit for creating custom scripting languages in Go.


### repl
A primitive [REPL](https://github.com/codr7/gofu/blob/main/utils/repl.go) is provided to allow playing around with the toolkit, it reads one form at a time and pushes prints the stack after each evaluation.

```
$ cd bin
$ ./mk
$ ./repl
gofu v1
  +(35 7)
[42]
```

Parens may be used to group multiple values into one form.

```
  (1 2 3)
[1 2 3]
```

Functions are called automatically when referenced, arguments are passed on the stack.

```
  (35 7 +)
[42]
```

The same result may be accomplished using call syntax, which is triggered by suffixing any identifier with parens.
All declared arguments must be included within the call form.

```
  +(35 7)
[42]
```

You may plug in a different parser to switch syntax, add custom types or operations; and still reuse the same REPL logic.

```
scope := gofu.Scope()
block := gofu.Block()
thread := gofu.Thread(scope)
utils.Repl(scope, parsers.Any(), block, thread)
```

### functions
Functions have a name, an argument list, a result list and a body.

```go
p := gofu.Pos("Test", -1, -1)

add := gofu.Func("+", []gofu.Type{types.Int(), types.Int()}, []gofu.Type{types.Int()},
	func(pos gofu.TPos, thread *gofu.TThread, pc *int) error {
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
    func(pos gofu.TPos, thread *gofu.TThread, pc *int) error {
	    stack := thread.Stack()
	    stack.Pop()
	    stack.Push(types.String(), "Bool!")
	    return nil
    })

f2 := gofu.Func("foo", []gofu.Type{types.Int()}, []gofu.Type{types.Int()},
    func(pos gofu.TPos, thread *gofu.TThread, pc *int) error {
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
* Bool: Any - true/false values
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