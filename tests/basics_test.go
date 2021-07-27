package tests

import (
	"testing"
	"github.com/codr7/gofu"
	"github.com/codr7/gofu/fimps"
	"github.com/codr7/gofu/forms"
	"github.com/codr7/gofu/inits"
	"github.com/codr7/gofu/ops"
	"github.com/codr7/gofu/types"
)

func Expect7(t *testing.T, stack *gofu.TStack) {
	if s := stack.Pop(); s.Type() != types.Int() {
		t.Errorf("Expected Int: %v", s.Type())
	} else if s.Value() != 7 {
		t.Errorf("Expected 7: %v", s.Value())
	}

	if !stack.Empty() {
		t.Errorf("Expected empty stack: %v", stack)
	}
}

func TestLiteral(t *testing.T) {
	block := gofu.Block()	
	scope := gofu.Scope()
	
	scope.Init()
	p := gofu.Pos("TestLiteral", -1, -1)
	
	if err := forms.Literal(p, types.Int(), 7).Compile(scope, block); err != nil {
		t.Fatal(err)
	}
	
	block.Emit(ops.Stop())	
	thread := gofu.Thread(scope)
	
	if err := block.Run(thread, 0); err != nil {
		t.Fatal(err)
	}

	Expect7(t, thread.Stack())
}

func TestBindSlot(t *testing.T) {
	block := gofu.Block()	
	scope := gofu.Scope()
	
	scope.Init()
	p := gofu.Pos("TestBindSlot", -1, -1)
	scope.BindSlot("foo", types.Int(), 7)
	
	if err := forms.Id(p, "foo").Compile(scope, block); err != nil {
		t.Fatal(err)
	}
	
	block.Emit(ops.Stop())	
	thread := gofu.Thread(scope)
	
	if err := block.Run(thread, 0); err != nil {
		t.Fatal(err)
	}

	Expect7(t, thread.Stack())
}

func TestBindId(t *testing.T) {
	block := gofu.Block()	
	scope := gofu.Scope()
	
	scope.Init()
	p := gofu.Pos("TestBindId", -1, -1)
	
	if err := forms.BindId(p, "foo", types.Int(), forms.Literal(p, types.Int(), 7)).Compile(scope, block); err != nil {
		t.Fatal(err)
	}

	if err := forms.Id(p, "foo").Compile(scope, block); err != nil {
		t.Fatal(err)
	}
	
	block.Emit(ops.Stop())	
	thread := gofu.Thread(scope)
	
	if err := block.Run(thread, 0); err != nil {
		t.Fatal(err)
	}

	Expect7(t, thread.Stack())
}

func TestFunc(t *testing.T) {
	block := gofu.Block()	
	scope := gofu.Scope()
	
	scope.Init()
	p := gofu.Pos("TestFunc", -1, -1)

	f := gofu.Func("foo", []gofu.Type{types.Int()}, []gofu.Type{types.Int()},
		func(pos gofu.TPos, thread *gofu.TThread, pc *int) error {
			stack := thread.Stack()
			stack.Push(types.Int(), stack.Pop().Value().(int) - 7)
			return nil
		})

	scope.BindSlot("foo", types.Func(), f)
	c := forms.Call(p, forms.Id(p, "foo"), forms.Literal(p, types.Int(), 14))

	if err := c.Compile(scope, block); err != nil {
		t.Fatal(err)
	}
	
	block.Emit(ops.Stop())	
	thread := gofu.Thread(scope)
	
	if err := block.Run(thread, 0); err != nil {
		t.Fatal(err)
	}

	Expect7(t, thread.Stack())
}

func TestFimp(t *testing.T) {
	block := gofu.Block()	
	scope := gofu.Scope()
	
	scope.Init()
	p := gofu.Pos("TestFimp", -1, -1)

	fimp, err := fimps.Compile(forms.Literal(p, types.Int(), 7), block)

	if err != nil {
		t.Fatal(err)
	}
	
	f := gofu.Func("foo", nil, []gofu.Type{types.Int()}, fimp)
	scope.BindSlot("foo", types.Func(), f)

	if err := forms.Call(p, forms.Id(p, "foo")).Compile(scope, block); err != nil {
		t.Fatal(err)
	}

	block.Emit(ops.Stop())	
	thread := gofu.Thread(scope)
	
	if err := block.Run(thread, 0); err != nil {
		t.Fatal(err)
	}

	Expect7(t, thread.Stack())
}

func TestMulti(t *testing.T) {
	block := gofu.Block()	
	scope := gofu.Scope()
	
	scope.Init()
	p := gofu.Pos("TestMulti", -1, -1)

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
	thread := gofu.Thread(scope)

	if err := block.Run(thread, 0); err != nil {
		t.Fatal(err)
	}

	Expect7(t, thread.Stack())
}

func TestDynamicCall(t *testing.T) {
	block := gofu.Block()	
	scope := gofu.Scope()
	inits.Scope(scope)
	
	p := gofu.Pos("TestDynamicCall", -1, -1)

	f := gofu.Func("foo", nil, []gofu.Type{types.Int()},
		func(pos gofu.TPos, thread *gofu.TThread, pc *int) error {
			thread.Stack().Push(types.Int(), 7)
			return nil
		})

	b := forms.BindId(p, "foo", types.Func(), forms.Literal(p, types.Func(), f))

	if err := b.Compile(scope, block); err != nil {
		t.Fatal(err)
	}
	
	c := forms.Call(p, forms.Id(p, "foo"))

	if err := c.Compile(scope, block); err != nil {
		t.Fatal(err)
	}
	
	block.Emit(ops.Stop())	
	thread := gofu.Thread(scope)
	
	if err := block.Run(thread, 0); err != nil {
		t.Fatal(err)
	}

	Expect7(t, thread.Stack())
}
