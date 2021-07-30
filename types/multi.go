package types

import (
	"github.com/codr7/gofu"
	"io"
)

type TMulti struct {
	gofu.BType
}

var multi *TMulti

func Multi() *TMulti {
	if multi == nil {
		multi = new(TMulti)
		multi.Init("Multi")
		multi.AddParent(Any(), false)
		multi.AddParent(Target(), true)
	}
	
	return multi
}

func (self TMulti) DumpValue(val interface{}, out io.Writer) {
	val.(*gofu.TMulti).Dump(out)
}

func (self TMulti) TargetArgCount(val interface{}) int {
	return val.(*gofu.TMulti).ArgCount()
}

func (self TMulti) TargetApplicable(val interface{}, stack *gofu.TStack) bool {
	return val.(*gofu.TMulti).Applicable(stack)	
}

func (self TMulti) CallTarget(val interface{}, pos gofu.TPos, thread *gofu.TThread, pc *int, check bool) error {
	return val.(*gofu.TMulti).Call(pos, thread, pc)		
}
