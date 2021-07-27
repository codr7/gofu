package gofu

import (
	"bufio"
)

type ParserLink = func(pos *TPos, in *bufio.Reader) (Form, error)

type TParser struct {
	chain []ParserLink
}

func Parser(links...ParserLink) *TParser {
	p := new(TParser)
	p.Chain(links...)
	return p
}

func (self *TParser) Chain(links...ParserLink) {
	self.chain = append(self.chain, links...)
}

func (self *TParser) Form(pos *TPos, in *bufio.Reader) (Form, error) {
	for _, c := range self.chain {
		if f, err := c(pos, in); err != nil {
			return nil, err
		} else if f != nil {
			return f, nil
		}
	}

	return nil, nil
}



