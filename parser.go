package gofu

import (
	"bufio"
)

type Parser interface {
	Form(pos *TPos, in *bufio.Reader) (Form, error)
}

