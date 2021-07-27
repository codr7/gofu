package errors

import (
	"fmt"
	"github.com/codr7/gofu"
)

type TParse struct{
	pos gofu.TPos
	message string
}

var parse TParse

func Parse(pos gofu.TPos, spec string, args...interface{}) TParse {
	return TParse{pos: pos, message: fmt.Sprintf(spec, args...)}
}

func (self TParse) Error() string {
	return fmt.Sprintf("Parse error in %v: %v", self.pos, self.message)}
