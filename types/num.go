package types

import (
	"github.com/codr7/gofu"
)

type TNum struct {
	gofu.BType
}

var num *TNum

func Num() *TNum {
	if num == nil {
		num = new(TNum)
		num.Init("Num")
	}
	
	return num
}
