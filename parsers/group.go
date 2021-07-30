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

func (self TGroup) Tail(fpos gofu.TPos, pos *gofu.TPos, in *bufio.Reader) (gofu.Form, error) {
	var c rune
	var f gofu.Form
	var out []gofu.Form
	var err error

	for {
		if c, _, err = in.ReadRune(); err == io.EOF {
			return nil, errors.Parse(*pos, "Missing group end")
		} else if err != nil {
			return nil, err
		}
		
		if c == ')' {
			break
		}

		if c == ';' {
			f, err = self.Tail(fpos, pos, in)

			if err != nil {
				if err == io.EOF {
					return nil, errors.Parse(*pos, "Missing group end")
				}

				return nil, err
			}

			out = append(out, f)
			break
		} else {
			in.UnreadRune()
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
	}

	if c != ';' && c != ')' {
		return nil, errors.Parse(*pos, "Missing group end: %v", c)
	}
	
	return forms.Group(fpos, out), nil
}

func (self TGroup) Form(pos *gofu.TPos, in *bufio.Reader) (gofu.Form, error) {
	var c rune
	var err error
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

	return self.Tail(fpos, pos, in)
}

