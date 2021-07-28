package parsers

import (
	"bufio"
	"github.com/codr7/gofu"
	"github.com/codr7/gofu/errors"
	"github.com/codr7/gofu/forms"
	"io"
)

type TGroup struct {
	memberParser gofu.Parser
}

func Group(memberParser gofu.Parser) *TGroup {
	return &TGroup{memberParser: memberParser}
}

func (self TGroup) Form(pos *gofu.TPos, in *bufio.Reader) (gofu.Form, error) {
	var c rune
	var f gofu.Form
	var err error
	var out []gofu.Form
	fpos := *pos
	
	if c, _, err = in.ReadRune(); err == io.EOF {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	if c != '(' {
		in.UnreadRune()
		return nil, nil
	}

	for {
		f, err = self.memberParser.Form(pos, in)

		if err != nil {
			if err == io.EOF {
				return nil, errors.Parse(*pos, "Missing group end")
			}
		}

		if f == nil {
			break
		}

		out = append(out, f)
	}

	if c, _, err = in.ReadRune(); err == io.EOF {
		return nil, errors.Parse(*pos, "Missing group end")
	} else if err != nil {
		return nil, err
	}
		
	if c != ')' {
		return nil, errors.Parse(*pos, "Missing group end: %v", c)
	}
	
	return forms.Group(fpos, out), nil
}

