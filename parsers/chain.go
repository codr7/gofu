package parsers

import (
	"bufio"
	"github.com/codr7/gofu"
)

func Chain(links...gofu.Parser) gofu.Parser {
	return func(pos *gofu.TPos, in *bufio.Reader) (gofu.Form, error) {
		for _, p := range links {
			if f, err := p(pos, in); err != nil {
				return nil, err
			} else if f != nil {
				return f, nil
			}
		}

		return nil, nil
	}
}
