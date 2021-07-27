package parsers

import (
	"bufio"
	"github.com/codr7/gofu"
	"io"
)

func Space(pos *gofu.TPos, in *bufio.Reader) (gofu.Form, error) {
	for {
		var c rune
		var err error
		
		if c, _, err = in.ReadRune(); err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}
		
		switch c {
		case ' ', '\t':
			pos.Next()
		case '\n':
			pos.NewLine()
		default:
			in.UnreadRune()
			c = 0
		}

		if c == 0 {
			break
		}
	}

	return nil, nil
}
