package gofu

import (
	"bufio"
)

type Parser = func(pos *TPos, in *bufio.Reader) (Form, error)



