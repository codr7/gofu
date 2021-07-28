package parsers

import (
	"github.com/codr7/gofu"
)

func Any() gofu.Parser {
	p := Chain()
	p.Chain(Space(), Group(p), Id())
	return p
}
