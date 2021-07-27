package errors

import (
	"fmt"
	"github.com/codr7/gofu"
)

type TCompile struct{
	pos gofu.TPos
	message string
}

var compile TCompile

func Compile(pos gofu.TPos, spec string, args...interface{}) TCompile {
	return TCompile{pos: pos, message: fmt.Sprintf(spec, args...)}
}

func (self TCompile) Error() string {
	return fmt.Sprintf("Compile error in %v: %v", self.pos, self.message)}
