package ops

import (
	"github.com/codr7/gofu"
)

type TPush struct {
	t gofu.Type
	v interface{}
}

func Push(t gofu.Type, v interface{}) TPush {
	return TPush{t: t, v: v}
}

func (self TPush) Eval(thread *gofu.TThread, pc *int) error {
	thread.Stack().Push(self.t, self.v)
	*pc++
	return  nil
}
