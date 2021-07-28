package parsers

import (
	"bufio"
	"github.com/codr7/gofu"
)

type TChain struct {
	links []gofu.Parser
}

func Chain(links...gofu.Parser) *TChain {
	return &TChain{links: links}
}

func (self *TChain) Chain(links...gofu.Parser) {
	self.links = append(self.links, links...)
}

func (self TChain) Form(pos *gofu.TPos, in *bufio.Reader) (gofu.Form, error) {
	for _, p := range self.links {
		if f, err := p.Form(pos, in); err != nil {
			return nil, err
		} else if f != nil {
			return f, nil
			}
	}
	
	return nil, nil
}
