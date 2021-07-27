package errors

import (
	"github.com/codr7/gofu"
)

type TEval struct{
	gofu.TError
}

var eval TEval

func Eval(pos gofu.TPos, spec string, args...interface{}) TEval {
	var e TEval
	e.Init(pos, spec, args...)
	return e
}

func (self TEval) Error() string {
	return "Eval " + self.TError.Error()
}
