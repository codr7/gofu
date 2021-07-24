package gofu

import (
	"fmt"
)

type TPos struct {
	source string
	line, column int
}

func Pos(src string, lin, col int) TPos {
	return TPos{source: src, line: lin, column: col}
}

func (self TPos) String() string {
	return fmt.Sprintf("'%v' at line @%v, column %v", self.source, self.line, self.column)
}
