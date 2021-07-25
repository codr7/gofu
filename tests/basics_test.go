package tests

import (
	"testing"
	"github.com/codr7/gofu"
	"github.com/codr7/gofu/forms"
	"github.com/codr7/gofu/ops"
	"github.com/codr7/gofu/types"
)

func Expect7(t *testing.T, stack *gofu.Stack) {
	if s := stack.Pop(); s.Type() != types.Int() {
		t.Errorf("Expected Int: %v", s.Type())
	} else if s.Value() != 7 {
		t.Errorf("Expected 7: %v", s.Value())
	}
}

func TestLiteral(t *testing.T) {
	var block gofu.Block	
	var scope gofu.Scope
	
	scope.Init()
	p := gofu.Pos("TestLiteral", -1, -1)
	
	if err := forms.Literal(p, types.Int(), 7).Compile(&scope, &block); err != nil {
		t.Fatal(err)
	}
	
	block.Emit(ops.Stop())	
	var calls gofu.CallStack
	var stack gofu.Stack
	stack.Init(scope.StackDepth())
	
	if err := block.Run(0, &calls, &stack); err != nil {
		t.Fatal(err)
	}

	Expect7(t, &stack)

	if !stack.Empty() {
		t.Errorf("Expected empty stack: %v", stack)
	}
}

func TestBindSlot(t *testing.T) {
	var block gofu.Block	
	var scope gofu.Scope
	
	scope.Init()
	p := gofu.Pos("TestBindSlot", -1, -1)
	scope.BindSlot("foo", types.Int(), 7)
	
	if err := forms.Id(p, "foo").Compile(&scope, &block); err != nil {
		t.Fatal(err)
	}
	
	block.Emit(ops.Stop())	
	var calls gofu.CallStack
	var stack gofu.Stack
	stack.Init(scope.StackDepth())
	
	if err := block.Run(0, &calls, &stack); err != nil {
		t.Fatal(err)
	}

	Expect7(t, &stack)
	
	if !stack.Empty() {
		t.Errorf("Expected empty stack: %v", stack)
	}
}

func TestBindId(t *testing.T) {
	var block gofu.Block	
	var scope gofu.Scope
	
	scope.Init()
	p := gofu.Pos("TestBindId", -1, -1)
	block.Emit(ops.Push(types.Int(), 7))
	
	if err := forms.BindId(p, "foo").Compile(&scope, &block); err != nil {
		t.Fatal(err)
	}

	if err := forms.Id(p, "foo").Compile(&scope, &block); err != nil {
		t.Fatal(err)
	}
	
	block.Emit(ops.Stop())	
	var calls gofu.CallStack
	var stack gofu.Stack
	stack.Init(scope.StackDepth())
	
	if err := block.Run(0, &calls, &stack); err != nil {
		t.Fatal(err)
	}

	Expect7(t, &stack)
	Expect7(t, &stack)
}

func TestFunc(t *testing.T) {
	var block gofu.Block	
	var scope gofu.Scope
	
	scope.Init()
	p := gofu.Pos("TestFunc", -1, -1)

	f := gofu.Func("foo", []gofu.Type{types.Int()}, types.Int(), func(stack *gofu.Stack) error {
		stack.Push(types.Int(), stack.Pop().Value().(int) - 7)
		return nil
	})

	scope.BindSlot("foo", types.Func(), f)
	c := forms.Call(p, forms.Id(p, "foo"), forms.Literal(p, types.Int(), 14))

	if err := c.Compile(&scope, &block); err != nil {
		t.Fatal(err)
	}

	block.Emit(ops.Stop())	
	var calls gofu.CallStack
	var stack gofu.Stack
	stack.Init(scope.StackDepth())
	
	if err := block.Run(0, &calls, &stack); err != nil {
		t.Fatal(err)
	}

	Expect7(t, &stack)
}




