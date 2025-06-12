//go:build !solution

package varfmt

import (
	"fmt"
	"strings"
	"unicode"
	"unicode/utf8"
)

func Sprintf(format string, args ...interface{}) string {
	var b strings.Builder
	b.Grow(len(format))

	idx := -1
	open := false
	argsFound := 0

	for i := 0; i < len(format); {
		r, size := utf8.DecodeRuneInString(format[i:])

		if r == '{' {
			open = true
		} else if open && unicode.IsDigit(r) {
			if idx < 0 {
				idx = 0
			}

			idx = idx * 10 + int(r - '0')
		} else if open && r == '}' {
			if idx >= 0 {
				fmt.Fprintf(&b, "%v", args[idx])
			} else {
				fmt.Fprintf(&b, "%v", args[argsFound])
			}

			open = false
			idx = -1
			argsFound += 1
		} else {
			b.WriteRune(r)
		}

		i += size
	}

	return b.String()
}
