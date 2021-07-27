package types

import (
	"github.com/codr7/gofu"
)

type TSeq struct {
	gofu.BType
}

var seq *TSeq

func Seq() *TSeq {
	if seq == nil {
		seq = new(TSeq)
		seq.Init("Seq")
	}
	
	return seq
}
