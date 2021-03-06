package parsers

import (
	"bufio"
	"github.com/codr7/gofu"
	"github.com/codr7/gofu/forms"
	"io"
	"strings"
	"unicode"
)

type TId struct {}

func Id() *TId {
	return &TId{}
}

func (self TId) Form(pos *gofu.TPos, in *bufio.Reader) (gofu.Form, error) {
	var out strings.Builder
	var c rune
	var err error
	fpos := *pos

	for {
		if c, _, err = in.ReadRune(); err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}

		if unicode.IsSpace(c) || c == '(' || c == ')' {
			in.UnreadRune()
			break
		}

		out.WriteRune(c)
		pos.Next()
	}

	if out.Len() == 0 {
		return nil, err
	}

	f := forms.Id(fpos, out.String())

	if c == '(' {
		var af gofu.Form
		af, err = Group(Any()).Form(pos, in)
		as := af.(forms.TGroup).Members()
		return forms.Call(fpos, f, as...), nil
	}

	return f, nil
}
