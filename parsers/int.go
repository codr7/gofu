package parsers

import (
	"bufio"
	"github.com/codr7/gofu"
	"github.com/codr7/gofu/forms"
	"github.com/codr7/gofu/types"
	"io"
	"unicode"
)

type TInt struct {
	base int
}

func Int(base int) *TInt {
	return &TInt{base: base}
}

func (self TInt) Form(pos *gofu.TPos, in *bufio.Reader) (gofu.Form, error) {
	out, mult := 0, 1
	var c rune
	var err error
	fpos := *pos

	for {
		if c, _, err = in.ReadRune(); err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}

		if !unicode.IsDigit(c) {
			in.UnreadRune()
			break
		}

		pos.Next()
		out = out*mult+int(c)-'0'
		mult *= self.base
	}

	if mult == 1 {
		return nil, nil
	}

	return forms.Literal(fpos, types.Int(), out), nil
}
