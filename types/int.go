package types

import (
	"github.com/codr7/gofu"
)

type TInt struct {
	gofu.BasicType
}

var Int TInt

func init() {
	Int.Init("Int")
}
