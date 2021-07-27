package types

import (
	"github.com/codr7/gofu"
)

type TChar struct {
	gofu.BType
}

var _char *TChar

func Char() *TChar {
	if _char == nil {
		_char = new(TChar)
		_char.Init("Char")
	}
	
	return _char
}
