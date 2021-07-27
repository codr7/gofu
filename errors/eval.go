package errors

import (
	"fmt"
	"github.com/codr7/gofu"
)

type TEval struct{
	pos gofu.TPos
	message string
}

var eval TEval

func Eval(pos gofu.TPos, spec string, args...interface{}) TEval {
	return TEval{pos: pos, message: fmt.Sprintf(spec, args...)}
}

func (self TEval) Error() string {
	return fmt.Sprintf("Eval error in %v: %v", self.pos, self.message)}
