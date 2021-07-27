package errors

import (
	"github.com/codr7/gofu"
)

type TCompile struct{
	gofu.TError
}

var compile TCompile

func Compile(pos gofu.TPos, spec string, args...interface{}) TCompile {
	var e TCompile
	e.Init(pos, spec, args...)
	return e
}

func (self TCompile) Error() string {
	return "Compile " + self.TError.Error()
}
