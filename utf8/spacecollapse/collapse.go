//go:build !solution

package spacecollapse

import (
	"strings"
	"unicode"
	"unicode/utf8"
)

func CollapseSpaces(input string) string {
	var b strings.Builder
	b.Grow(len(input))

	prevSpace := false

	for i := 0; i < len(input); {
		r, size := utf8.DecodeRuneInString(input[i:])

		if isSpace := unicode.IsSpace(r); !isSpace || !prevSpace {
			if isSpace {
				b.WriteString(" ")
			} else {
				b.WriteRune(r)
			}

			prevSpace = isSpace
		}

		i += size
	}

	return b.String()
}
