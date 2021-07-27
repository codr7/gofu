package errors

import (
	"github.com/codr7/gofu"
)

type TParse struct{
	gofu.TError
}

var parse TParse

func Parse(pos gofu.TPos, spec string, args...interface{}) TParse {
	var e TParse
	e.Init(pos, spec, args...)
	return e
}

func (self TParse) Error() string {
	return "Parse " + self.TError.Error()
}
